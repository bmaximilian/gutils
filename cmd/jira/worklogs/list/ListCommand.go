package list

import (
	"github.com/bmaximilian/gutils/internal/jira/util"
	"github.com/bmaximilian/gutils/pkg/jira/projects"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	googleLogger "github.com/google/logger"
	"github.com/spf13/cobra"
	"log"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "List all worklogs of a given user",
	Run: func(_ *cobra.Command, _ []string) {
		l := logger.GetLogger()
		googleLogger.SetFlags(log.LUTC)

		jiraProjectsRequestService := projects.NewJiraProjectsRequestService(util.GetJiraRequestService())

		res, err := jiraProjectsRequestService.GetAllProjects()

		// GET /search?startAt=0 -> alle issues abholen
		// GET /issue/{issueIdOrKey}/worklog -> für jedes issue den Worklog abholen
		// Worklog-Response filtern nach name (als Flag konfigurierbar (required))
		// Gefilterte worklogs als Tabelle printen (Mit "Gesamt" Zeile) + CSV Export generieren

		if err != nil {
			l.Fatalln(err)
		}

		l.Infoln(res)
	},
}

// Set the default viper values
func SetDefaults() {}

// Initializes the command line tool
func InitCommand() {}
