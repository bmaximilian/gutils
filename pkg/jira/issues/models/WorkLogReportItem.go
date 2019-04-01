package models

type WorkLogReportItem struct {
	IssueKey          string
	Started           string
	Updated           string
	TimeSpent         string
	TimeSpentSeconds  int64
	WorkLogId         string
	IssueId           string
	AuthorShortName   string
	AuthorDisplayName string
	Comment           string
}
