package util

import (
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"github.com/olekukonko/tablewriter"
)

type TableOptions struct {
	WithAuthor bool
}

func GenerateWorklogTable(
	table *tablewriter.Table,
	workLogReportItems []models.WorkLogReportItem,
	options *TableOptions,
) {
	headerSlice := []string{"Date", "Ticket", "Time Spent"}
	if options.WithAuthor {
		headerSlice = append(headerSlice, "Author")
	}
	table.SetHeader(headerSlice)

	generateTableBody(table, workLogReportItems, func(item interface{}) []string {
		workLogReportItem := item.(models.WorkLogReportItem)
		bodySlice := []string{
			workLogReportItem.StartedDate.Format("Mon") + " " +
				workLogReportItem.StartedDate.Format("2006-01-02") + " " +
				workLogReportItem.StartedDate.Format("15:04"),
			workLogReportItem.IssueKey,
			workLogReportItem.TimeSpent,
		}

		if options.WithAuthor {
			bodySlice = append(bodySlice, workLogReportItem.AuthorDisplayName)
		}

		return bodySlice
	})

	footerSlice := []string{"", "", GetWorkLogSummaryTimeFormatted(workLogReportItems)}
	if options.WithAuthor {
		footerSlice = append(footerSlice, "")
	}
	table.SetFooter(footerSlice)
}

func addRowToTableForMap(table *tablewriter.Table, headerMap []string) {
	table.Append(headerMap)
}

func generateTableBody(table *tablewriter.Table, workLogReportItems []models.WorkLogReportItem, getBodyRowValues func(item interface{}) []string) {
	for _, item := range workLogReportItems {
		addRowToTableForMap(table, getBodyRowValues(item))
	}
}
