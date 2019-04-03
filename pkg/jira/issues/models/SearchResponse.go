package models

type SearchResponse struct {
	Expand     string  `json:"expand"`
	StartAt    int64   `json:"startAt"`
	MaxResults int64   `json:"maxResults"`
	Total      int64   `json:"total"`
	Issues     []Issue `json:"issues"`
}
