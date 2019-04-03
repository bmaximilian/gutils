package util

import (
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"sort"
)

func SortWorklogsByStartDate(workLogReportItems []models.WorkLogReportItem) {
	sort.Slice(workLogReportItems, func(i, j int) bool {
		return workLogReportItems[i].StartedDate.Before(workLogReportItems[j].StartedDate)
	})
}
