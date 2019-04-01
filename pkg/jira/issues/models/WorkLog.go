package models

type WorkLog struct {
	Self             string     `json:"self"`
	Author           Author     `json:"author"`
	UpdateAuthor     Author     `json:"updateAuthor"`
	Comment          string     `json:"comment"`
	Updated          string     `json:"updated"`
	Started          string     `json:"started"`
	TimeSpent        string     `json:"timeSpent"`
	Visibility       Visibility `json:"visibility"`
	TimeSpentSeconds int64      `json:"timeSpentSeconds"`
	Id               string     `json:"id"`
	IssueId          string     `json:"issueId"`
}
