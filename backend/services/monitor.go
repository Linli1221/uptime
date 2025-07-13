package services

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"
	"time"
	"uptime/backend/database"
	"uptime/backend/models"
)

var DeletedMonitorIDs = make(chan uint, 100)

func StartMonitoring() {
	log.Printf("[INFO] Monitoring service started at %v", time.Now())
	ticker := time.NewTicker(10 * time.Second) // Check for new monitors every 10 seconds
	monitorTimers := make(map[uint]*time.Timer)

	for {
		select {
		case <-ticker.C:
			var monitors []models.Monitor
			database.DB.Find(&monitors)

			for i := range monitors {
				monitor := &monitors[i]

				// Skip if timer already exists for this monitor
				if _, exists := monitorTimers[monitor.ID]; exists {
					continue
				}

				// Create individual timer for each monitor based on its interval
				interval := time.Duration(monitor.Interval) * time.Second
				log.Printf("[INFO] Creating timer for monitor ID=%d Name=%s Interval=%v", monitor.ID, monitor.Name, interval)
				monitorTimers[monitor.ID] = time.AfterFunc(0, func() {
					scheduleMonitorCheck(monitor, monitorTimers, interval)
				})
			}
		case id := <-DeletedMonitorIDs:
			if timer, exists := monitorTimers[id]; exists {
				timer.Stop()
				delete(monitorTimers, id)
				log.Printf("[INFO] Stopped and removed timer for monitor ID=%d", id)
			}
		}
	}
}

func scheduleMonitorCheck(monitor *models.Monitor, timers map[uint]*time.Timer, interval time.Duration) {
	// Add a delay before the check if specified
	if monitor.Delay > 0 {
		delay := time.Duration(monitor.Delay) * time.Second
		log.Printf("[INFO] Monitor ID=%d Name=%s: Delaying check by %v", monitor.ID, monitor.Name, delay)
		time.Sleep(delay)
	}

	checkMonitorWithRetries(monitor)

	// Schedule next check
	timers[monitor.ID] = time.AfterFunc(interval, func() {
		scheduleMonitorCheck(monitor, timers, interval)
	})
}

func checkMonitorWithRetries(monitor *models.Monitor) {
	maxRetries := int(monitor.Retries)
	retryInterval := time.Duration(monitor.RetryInterval) * time.Second
	
	for attempt := 0; attempt <= maxRetries; attempt++ {
		success := checkMonitor(monitor)
		if success || attempt == maxRetries {
			break
		}
		
		if attempt < maxRetries {
			log.Printf("[WARN] Monitor ID=%d Name=%s failed (attempt %d/%d), retrying in %v",
				monitor.ID, monitor.Name, attempt+1, maxRetries+1, retryInterval)
			time.Sleep(retryInterval)
		}
	}
}

func checkMonitor(monitor *models.Monitor) bool {
	client := &http.Client{
		Timeout: time.Duration(monitor.Timeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: monitor.IgnoreTLS},
		},
	}

	req, err := http.NewRequest(monitor.HttpMethod, monitor.URL, strings.NewReader(monitor.HttpBody))
	if err != nil {
		log.Printf("[ERROR] Monitor ID=%d Name=%s: Error creating request: %s", monitor.ID, monitor.Name, err.Error())
		updateMonitorStatus(monitor, "down")
		return false
	}

	// Add headers
	// For simplicity, assuming headers are in "Key:Value\nKey2:Value2" format
	if monitor.HttpHeaders != "" {
		headers := strings.Split(monitor.HttpHeaders, "\n")
		for _, header := range headers {
			parts := strings.SplitN(header, ":", 2)
			if len(parts) == 2 {
				req.Header.Set(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			}
		}
	}

	// Add basic auth
	if monitor.AuthMethod == "basic" && monitor.AuthUsername != "" {
		req.SetBasicAuth(monitor.AuthUsername, monitor.AuthPassword)
	}

	start := time.Now()
	resp, err := client.Do(req)
	latency := time.Since(start).Milliseconds()

	if err != nil {
		log.Printf("[ERROR] Monitor ID=%d Name=%s: HTTP request failed: %s", monitor.ID, monitor.Name, err.Error())
		updateMonitorStatus(monitor, "down")
		return false
	}
	defer resp.Body.Close()

	// Store latency
	latencyRecord := models.Latency{MonitorID: monitor.ID, Latency: uint(latency)}
	database.DB.Create(&latencyRecord)

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("[INFO] Monitor ID=%d Name=%s: Status UP (HTTP %d), Latency: %dms", monitor.ID, monitor.Name, resp.StatusCode, latency)
		updateMonitorStatus(monitor, "up")
		return true
	} else {
		log.Printf("[WARN] Monitor ID=%d Name=%s: Status DOWN (HTTP %d), Latency: %dms", monitor.ID, monitor.Name, resp.StatusCode, latency)
		updateMonitorStatus(monitor, "down")
		return false
	}
}

func updateMonitorStatus(monitor *models.Monitor, status string) {
	oldStatus := monitor.Status
	monitor.Status = status
	database.DB.Save(monitor)
	log.Printf("[INFO] Monitor ID=%d Name=%s: Status changed from %s to %s", monitor.ID, monitor.Name, oldStatus, status)
}