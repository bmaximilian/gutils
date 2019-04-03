package list

import (
	"github.com/bmaximilian/gutils/internal/jira/util"
	"github.com/bmaximilian/gutils/pkg/jira/projects"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	googleLogger "github.com/google/logger"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var Command = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Run: func(cmd *cobra.Command, args []string) {
		l := logger.GetLogger()
		googleLogger.SetFlags(log.LUTC)

		server := cmd.Flag("server").Value.String()

		jiraProjectsRequestService := projects.NewJiraProjectsRequestService(util.GetJiraRequestService(server))

		// get all projects
		fetchedProjects, fetchProjectErr := jiraProjectsRequestService.GetAllProjects()
		if fetchProjectErr != nil {
			l.Fatalln(fetchProjectErr)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Key", "Name"})

		for _, project := range *fetchedProjects {
			table.Append([]string{project.Key, project.Name})
		}

		table.Render()
	},
}

// Set the default viper values
func SetDefaults() {}

// Initializes the command line tool
func InitCommand() {}
