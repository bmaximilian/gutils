package list

import (
	"fmt"
	"github.com/araddon/dateparse"
	"github.com/bmaximilian/gutils/internal/jira/util"
	"github.com/bmaximilian/gutils/pkg/jira/issues"
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"github.com/bmaximilian/gutils/pkg/jira/projects"
	jiraUtil "github.com/bmaximilian/gutils/pkg/jira/util"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	googleLogger "github.com/google/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"math"
	"os"
	"sort"
)

var projectsFilter string
var user string
var startDateFlag string
var endDateFlag string

var Command = &cobra.Command{
	Use:   "list",
	Short: "List all worklogs of a given user",
	Run: func(cmd *cobra.Command, _ []string) {
		l := logger.GetLogger()
		googleLogger.SetFlags(log.LUTC)
		jiraRequestService := util.GetJiraRequestService()

		jiraProjectsRequestService := projects.NewJiraProjectsRequestService(jiraRequestService)
		jiraSearchRequestService := issues.NewJiraSearchRequestService(jiraRequestService)
		jiraIssueRequestService := issues.NewJiraIssueRequestService(jiraRequestService)

		projectKeys := cmd.Flag("projects").Value.String()
		userNameFilter := cmd.Flag("user").Value.String()
		fromDateFilter := cmd.Flag("from").Value.String()
		toDateFilter := cmd.Flag("to").Value.String()

		if projectKeys == "all" {
			projectKeys = ""
			// get all projects
			projects, fetchProjectErr := jiraProjectsRequestService.GetAllProjects()
			if fetchProjectErr != nil {
				l.Fatalln(fetchProjectErr)
			}

			// Accumulate project keys for issue search request
			for i, project := range *projects {
				projectKeys += project.Key
				if i+1 < len(*projects) {
					projectKeys += ", "
				}
			}
		}

		l.Infoln("Search in Projects: " + projectKeys)
		// Get all issues for the project keys
		issues, fetchIssuesErr := jiraSearchRequestService.GetAllIssuesForProject(projectKeys)
		if fetchIssuesErr != nil {
			l.Fatalln(fetchIssuesErr)
		}

		rawWorkLogReportItems, getReportErr := jiraIssueRequestService.GenerateWorkLogReportForIssues(issues, userNameFilter)
		if getReportErr != nil {
			l.Fatalln(getReportErr)
		}

		table := tablewriter.NewWriter(os.Stdout)
		timeSummaryInSeconds := int64(0)

		workLogReportItems := *rawWorkLogReportItems
		sort.Slice(workLogReportItems, func(i, j int) bool {
			return workLogReportItems[i].StartedDate.Before(workLogReportItems[j].StartedDate)
		})

		if fromDateFilter != "" {
			parsedFromDate, fromParseErr := dateparse.ParseAny(fromDateFilter + " 00:00:01+0100")
			if fromParseErr != nil {
				l.Fatalln(fromParseErr)
			}
			workLogReportItems = jiraUtil.FilterWorklogs(workLogReportItems, func(item models.WorkLogReportItem) bool {
				return item.StartedDate.After(parsedFromDate) || item.StartedDate.Equal(parsedFromDate)
			})
		}

		if toDateFilter != "" {
			parsedToDate, toParseErr := dateparse.ParseAny(toDateFilter + " 23:59:59+0100")
			if toParseErr != nil {
				l.Fatalln(toParseErr)
			}
			workLogReportItems = jiraUtil.FilterWorklogs(workLogReportItems, func(item models.WorkLogReportItem) bool {
				return item.StartedDate.Before(parsedToDate) || item.StartedDate.Equal(parsedToDate)
			})
		}

		if userNameFilter == "" {
			table.SetHeader([]string{"Date", "Ticket", "Time Spent", "Author"})

			for _, workLogReportItem := range workLogReportItems {
				table.Append([]string{
					workLogReportItem.StartedDate.Format("2006-01-02") + " " + workLogReportItem.StartedDate.Format("15:04"),
					workLogReportItem.IssueKey,
					workLogReportItem.TimeSpent,
					workLogReportItem.AuthorDisplayName,
				})
				timeSummaryInSeconds += workLogReportItem.TimeSpentSeconds
			}

			timeSummaryInHours := (float64(timeSummaryInSeconds) / 60.00) / 60.00
			remainingMinutes := timeSummaryInHours - math.Ceil(timeSummaryInHours)
			table.SetFooter([]string{"", "", "", fmt.Sprintf("%vh %vm", math.Ceil(timeSummaryInHours), remainingMinutes*60)})
		} else {
			table.SetHeader([]string{"Date", "Ticket", "Time Spent"})

			for _, workLogReportItem := range workLogReportItems {
				table.Append([]string{
					workLogReportItem.StartedDate.Format("2006-01-02") + " " + workLogReportItem.StartedDate.Format("15:04"),
					workLogReportItem.IssueKey,
					workLogReportItem.TimeSpent,
				})
				timeSummaryInSeconds += workLogReportItem.TimeSpentSeconds
			}

			timeSummaryInHours := (float64(timeSummaryInSeconds) / 60.00) / 60.00
			remainingMinutes := timeSummaryInHours - math.Floor(timeSummaryInHours)
			table.SetFooter([]string{"", "", fmt.Sprintf("%vh %vm", math.Floor(timeSummaryInHours), remainingMinutes*60)})
		}

		table.Render()
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
}
