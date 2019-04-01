package list

import (
	"fmt"
	"github.com/bmaximilian/gutils/internal/jira/util"
	"github.com/bmaximilian/gutils/pkg/jira/issues"
	"github.com/bmaximilian/gutils/pkg/jira/projects"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	googleLogger "github.com/google/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"math"
	"os"
	"sort"
	"time"
)

var projectsFilter string
var user string

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
			parsedITime, timeParseErrI := time.Parse("2006-01-02T15:04:05.999+0100", workLogReportItems[i].Started)
			parsedJTime, timeParseErrJ := time.Parse("2006-01-02T15:04:05.999+0100", workLogReportItems[j].Started)
			if timeParseErrI != nil || timeParseErrJ != nil {
				l.Fatalln(timeParseErrJ, timeParseErrI)
			}
			return parsedITime.Before(parsedJTime)
		})

		if userNameFilter == "" {
			table.SetHeader([]string{"Date", "Ticket", "Time Spent", "Author"})

			for _, workLogReportItem := range workLogReportItems {
				parsedStartTime, _ := time.Parse("2006-01-02T15:04:05.999+0100", workLogReportItem.Started)

				table.Append([]string{
					parsedStartTime.Format("2006-01-02") + " " + parsedStartTime.Format("15:04"),
					workLogReportItem.IssueKey,
					workLogReportItem.TimeSpent,
					workLogReportItem.AuthorDisplayName,
				})
				timeSummaryInSeconds += workLogReportItem.TimeSpentSeconds
			}

			timeSummaryInHours := float64((timeSummaryInSeconds / 60.00) / 60.00)
			remainingMinutes := timeSummaryInHours - math.Ceil(timeSummaryInHours)
			table.SetFooter([]string{"", "", "", fmt.Sprintf("%vh %vm", math.Ceil(timeSummaryInHours), remainingMinutes*60)})
		} else {
			table.SetHeader([]string{"Date", "Ticket", "Time Spent"})

			for _, workLogReportItem := range workLogReportItems {
				parsedStartTime, _ := time.Parse("2006-01-02T15:04:05.999+0100", workLogReportItem.Started)

				table.Append([]string{
					parsedStartTime.Format("2006-01-02") + " " + parsedStartTime.Format("15:04"),
					workLogReportItem.IssueKey,
					workLogReportItem.TimeSpent,
				})
				timeSummaryInSeconds += workLogReportItem.TimeSpentSeconds
			}

			timeSummaryInHours := float64((timeSummaryInSeconds / 60.00) / 60.00)
			remainingMinutes := timeSummaryInHours - math.Ceil(timeSummaryInHours)
			table.SetFooter([]string{"", "", fmt.Sprintf("%vh %vm", math.Ceil(timeSummaryInHours), remainingMinutes*60)})
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
}
