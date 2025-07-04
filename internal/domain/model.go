package domain

import "time"

type ItemReport struct {
	Path   string   `json:"path"`
	SizeMB float64  `json:"size_mb"`
	Files  int      `json:"files"`
	Dirs   int      `json:"dirs"`
	Errors []string `json:"errors,omitempty"`
}

// SummaryMetrics provides general statistics.
type SummaryMetrics struct {
	TotalSizeMB      float64 `json:"total_size_mb"`
	TotalFiles       int     `json:"total_files"`
	TotalDirectories int     `json:"total_directories"`
	ItemsCleaned     int     `json:"items_cleaned"`
}

// SystemMetadata describes system context during cleanup.
type SystemMetadata struct {
	Hostname   string `json:"hostname"`
	OSVersion  string `json:"os_version"`
	User       string `json:"user"`
	AppVersion string `json:"app_version"`
}

// Report is the main structure summarizing a cleaning session.
type Report struct {
	StartedAt  time.Time      `json:"started_at"`
	FinishedAt time.Time      `json:"finished_at"`
	Duration   float64        `json:"duration"`
	Summary    SummaryMetrics `json:"summary"`
	Items      []ItemReport   `json:"items"`
	System     SystemMetadata `json:"system"`
}
