package models

import "time"

type WorkLogReportItem struct {
	IssueKey          string
	Started           string
	StartedDate       time.Time
	Updated           string
	UpdatedDate       time.Time
	TimeSpent         string
	TimeSpentSeconds  int64
	WorkLogId         string
	IssueId           string
	AuthorShortName   string
	AuthorDisplayName string
	Comment           string
}
