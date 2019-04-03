package list

import (
	"github.com/bmaximilian/gutils/internal/jira/util"
	jiraUtil "github.com/bmaximilian/gutils/pkg/jira/util"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	googleLogger "github.com/google/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/tealeg/xlsx"
	"log"
	"os"
	"time"
)

var projectsFilter string
var user string
var generateReportFlag bool
var startDateFlag string
var endDateFlag string

var Command = &cobra.Command{
	Use:   "list",
	Short: "List all worklogs of a given user",
	Run: func(cmd *cobra.Command, _ []string) {
		l := logger.GetLogger()
		googleLogger.SetFlags(log.LUTC)

		projectKeys := cmd.Flag("projects").Value.String()
		userNameFilter := cmd.Flag("user").Value.String()
		fromDateFilter := cmd.Flag("from").Value.String()
		toDateFilter := cmd.Flag("to").Value.String()

		table := tablewriter.NewWriter(os.Stdout)

		workLogReportItemsRaw, fetchWorkLogItemsErr := jiraUtil.GetWorkLogReportForProjectsAndUser(util.GetJiraRequestService(), projectKeys, userNameFilter)
		if fetchWorkLogItemsErr != nil {
			l.Fatalln(fetchWorkLogItemsErr)
		}

		workLogReportItems := *workLogReportItemsRaw

		jiraUtil.SortWorklogsByStartDate(workLogReportItems)

		var filterErr error = nil
		if fromDateFilter != "" {
			workLogReportItems, filterErr = jiraUtil.OnlyWorkLogsStartDateAfterDate(workLogReportItems, fromDateFilter+" 00:00:01+0100")
		}

		if toDateFilter != "" {
			workLogReportItems, filterErr = jiraUtil.OnlyWorkLogsStartDateBeforeDate(workLogReportItems, toDateFilter+" 23:59:59+0100")
		}

		if filterErr != nil {
			l.Fatalln(filterErr)
		}

		if userNameFilter == "" {
			table.SetHeader([]string{"Date", "Ticket", "Time Spent", "Author"})

			for _, workLogReportItem := range workLogReportItems {
				table.Append([]string{
					workLogReportItem.StartedDate.Format("Mon") + " " + workLogReportItem.StartedDate.Format("2006-01-02") + " " + workLogReportItem.StartedDate.Format("15:04"),
					workLogReportItem.IssueKey,
					workLogReportItem.TimeSpent,
					workLogReportItem.AuthorDisplayName,
				})
			}

			table.SetFooter([]string{"", "", "", jiraUtil.GetWorkLogSummaryTimeFormatted(workLogReportItems)})
		} else {
			table.SetHeader([]string{"Date", "Ticket", "Time Spent"})

			for _, workLogReportItem := range workLogReportItems {
				table.Append([]string{
					workLogReportItem.StartedDate.Format("Mon") + " " + workLogReportItem.StartedDate.Format("2006-01-02") + " " + workLogReportItem.StartedDate.Format("15:04"),
					workLogReportItem.IssueKey,
					workLogReportItem.TimeSpent,
				})
			}

			table.SetFooter([]string{"", "", jiraUtil.GetWorkLogSummaryTimeFormatted(workLogReportItems)})
		}

		table.Render()
		if cmd.Flag("report").Value.String() == "true" {
			file := xlsx.NewFile()
			xlsxCreateErr := jiraUtil.GenerateWorklogXLSXExportSheet(file, workLogReportItems, &jiraUtil.XLSXExportOptions{})
			if xlsxCreateErr != nil {
				l.Fatalln(xlsxCreateErr.Error())
			}

			xlsxWriteErr := file.Save("report_" + userNameFilter + "_" + time.Now().Format(time.RFC3339) + ".xlsx")
			if xlsxWriteErr != nil {
				l.Fatalln(xlsxWriteErr.Error())
			}
		}
	},
}

// Set the default viper values
func SetDefaults() {}

// Initializes the command line tool
func InitCommand() {
	Command.Flags().StringVarP(
		&user,
		"user",
		"u",
		"",
		"The username to get the worklog for",
	)

	Command.Flags().StringVarP(
		&projectsFilter,
		"projects",
		"p",
		"all",
		"The projects that should be looked up",
	)

	Command.Flags().StringVarP(
		&startDateFlag,
		"from",
		"s",
		"",
		"The start date",
	)

	Command.Flags().StringVarP(
		&endDateFlag,
		"to",
		"t",
		"",
		"The end date",
	)
	Command.Flags().BoolVarP(
		&generateReportFlag,
		"report",
		"r",
		false,
		"Pass that flag to generate an xslx report",
	)
}
