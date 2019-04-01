package models

type Issue struct {
	Expand   string `json:"expand"`
	Id       string `json:"id"`
	Self     string `json:"self"`
	Key      string `json:"key"`
	WorkLogs []WorkLog
}

func (i *Issue) WorkLogsWithAuthor(name string) *[]WorkLog {
	filteredWorkLogs := make([]WorkLog, 0)

	for _, workLog := range i.WorkLogs {
		if workLog.Author.Name == name {
			filteredWorkLogs = append(filteredWorkLogs, workLog)
		}
	}

	return &filteredWorkLogs
}
