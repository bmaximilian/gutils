package util

import (
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"github.com/tealeg/xlsx"
)

type XLSXExportOptions struct {
	WithAuthor bool
}

func GenerateWorklogXLSXExportSheet(file *xlsx.File, workLogReportItems []models.WorkLogReportItem, options *XLSXExportOptions) error {
	// file := xlsx.NewFile()
	sheet, createSheetErr := file.AddSheet("Worklog")
	if createSheetErr != nil {
		return createSheetErr
	}

	headerSlice := []string{"Date", "Ticket", "Time Spent"}
	if options.WithAuthor {
		headerSlice = append(headerSlice, "Author")
	}
	headerSlice = append(headerSlice, []string{
		"Comment",
	}...)
	addRowToSheetForMap(sheet, headerSlice)

	generateBody(sheet, workLogReportItems, func(item interface{}) []string {
		workLogReportItem := item.(models.WorkLogReportItem)
		bodySlice := []string{
			workLogReportItem.StartedDate.Format("Mon") + " " + workLogReportItem.StartedDate.Format("2006-01-02"),
			workLogReportItem.IssueKey,
			workLogReportItem.TimeSpent,
		}

		if options.WithAuthor {
			bodySlice = append(bodySlice, workLogReportItem.AuthorDisplayName)
		}

		bodySlice = append(bodySlice, []string{
			workLogReportItem.Comment,
		}...)
		return bodySlice
	})

	footerSlice := []string{"", "", GetWorkLogSummaryTimeFormatted(workLogReportItems)}
	if options.WithAuthor {
		footerSlice = append(footerSlice, "")
	}
	footerSlice = append(footerSlice, []string{
		"",
	}...)
	addRowToSheetForMap(sheet, []string{"", "", GetWorkLogSummaryTimeFormatted(workLogReportItems), "", ""})

	return nil
}

func addRowToSheetForMap(sheet *xlsx.Sheet, headerMap []string) {
	row := sheet.AddRow()

	for _, headerValue := range headerMap {
		cell := row.AddCell()
		cell.Value = headerValue
	}
}

func generateBody(sheet *xlsx.Sheet, workLogReportItems []models.WorkLogReportItem, getBodyRowValues func(item interface{}) []string) {
	for _, item := range workLogReportItems {
		addRowToSheetForMap(sheet, getBodyRowValues(item))
	}
}
