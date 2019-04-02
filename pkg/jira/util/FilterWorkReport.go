package util

import "github.com/bmaximilian/gutils/pkg/jira/issues/models"

func FilterWorklogs(s []models.WorkLogReportItem, p func(item models.WorkLogReportItem) bool) []models.WorkLogReportItem {
	b := s[:0]
	for _, x := range s {
		if p(x) {
			b = append(b, x)
		}
	}
	return b
}
