package models

import "gorm.io/gorm"

type Latency struct {
	gorm.Model
	MonitorID uint `json:"monitorId"`
	Latency   uint `json:"latency"` // Latency in milliseconds
}