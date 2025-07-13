package models

import "gorm.io/gorm"

type Monitor struct {
	gorm.Model
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string         `json:"name"`
	URL            string `json:"url"`
	Interval       uint   `json:"interval"` // Interval in seconds
	Status         string `json:"status"`   // "up", "down", "pending"
	Timeout        uint   `json:"timeout"`
	Retries        uint   `json:"retries"`
	RetryInterval  uint   `json:"retryInterval"`
	IgnoreTLS      bool   `json:"ignoreTls"`
	HttpMethod     string `json:"httpMethod"`
	HttpBody       string `json:"httpBody"`
	HttpHeaders    string `json:"httpHeaders"`
	AuthMethod     string `json:"authMethod"`
	AuthUsername   string `json:"authUsername"`
	AuthPassword   string `json:"authPassword"`
	Delay          uint   `json:"delay"` // Delay in seconds before check
}