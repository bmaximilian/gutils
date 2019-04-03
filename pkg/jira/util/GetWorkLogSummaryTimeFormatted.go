package util

import (
	"fmt"
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"math"
)

func GetWorkLogSummaryTimeFormatted(workLogReportItems []models.WorkLogReportItem) string {
	timeSummaryInSeconds := int64(0)

	for _, workLogReportItem := range workLogReportItems {
		timeSummaryInSeconds += workLogReportItem.TimeSpentSeconds
	}

	timeSummaryInHours := (float64(timeSummaryInSeconds) / 60.00) / 60.00
	remainingMinutes := timeSummaryInHours - math.Floor(timeSummaryInHours)
	return fmt.Sprintf("%vh %vm", math.Floor(timeSummaryInHours), remainingMinutes*60)
}
