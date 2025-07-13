package controllers

import (
	"net/http"
	"uptime/backend/database"
	"uptime/backend/models"
	"uptime/backend/services"

	"github.com/gin-gonic/gin"
)

// CreateMonitor creates a new monitor
func CreateMonitor(c *gin.Context) {
	var monitor models.Monitor
	if err := c.ShouldBindJSON(&monitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&monitor)
	c.JSON(http.StatusOK, monitor)
}

// GetMonitors retrieves all monitors
func GetMonitors(c *gin.Context) {
	var monitors []models.Monitor
	database.DB.Find(&monitors)

	type MonitorWithLatency struct {
		models.Monitor
		Latency uint `json:"latency"`
	}

	var monitorsWithLatency []MonitorWithLatency

	for _, m := range monitors {
		var latency models.Latency
		database.DB.Where("monitor_id = ?", m.ID).Order("created_at desc").First(&latency)
		monitorsWithLatency = append(monitorsWithLatency, MonitorWithLatency{
			Monitor: m,
			Latency: latency.Latency,
		})
	}

	c.JSON(http.StatusOK, monitorsWithLatency)
}

// GetMonitor retrieves a single monitor by ID
func GetMonitor(c *gin.Context) {
	id := c.Params.ByName("id")
	var monitor models.Monitor
	if err := database.DB.First(&monitor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found"})
		return
	}

	var latencies []models.Latency
	database.DB.Where("monitor_id = ?", monitor.ID).Order("created_at desc").Limit(100).Find(&latencies)

	c.JSON(http.StatusOK, gin.H{
		"monitor":   monitor,
		"latencies": latencies,
	})
}

// UpdateMonitor updates an existing monitor
func UpdateMonitor(c *gin.Context) {
	id := c.Params.ByName("id")
	var monitor models.Monitor
	if err := database.DB.First(&monitor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found"})
		return
	}

	if err := c.ShouldBindJSON(&monitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Save(&monitor)
	c.JSON(http.StatusOK, monitor)
}

// DeleteMonitor deletes a monitor
func DeleteMonitor(c *gin.Context) {
	id := c.Params.ByName("id")
	var monitor models.Monitor
	if err := database.DB.First(&monitor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found"})
		return
	}
	database.DB.Unscoped().Delete(&monitor)
	services.DeletedMonitorIDs <- monitor.ID
	c.JSON(http.StatusOK, gin.H{"message": "Monitor deleted successfully"})
}

// ManualCheck triggers a manual check for a specific monitor
func ManualCheck(c *gin.Context) {
	id := c.Params.ByName("id")
	var monitor models.Monitor
	if err := database.DB.First(&monitor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Monitor not found"})
		return
	}

	// TODO: Trigger manual check - for now just update status to pending
	monitor.Status = "pending"
	database.DB.Save(&monitor)
	
	c.JSON(http.StatusOK, gin.H{"message": "Manual check triggered", "monitor": monitor})
}