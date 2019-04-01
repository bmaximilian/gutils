package models

type WorkLogsResponse struct {
	StartAt    int64     `json:"startAt"`
	MaxResults int64     `json:"maxResults"`
	Total      int64     `json:"total"`
	WorkLogs   []WorkLog `json:"worklogs"`
}
