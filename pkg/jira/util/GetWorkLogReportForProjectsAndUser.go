package util

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/jira/issues"
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"github.com/bmaximilian/gutils/pkg/jira/projects"
)

func GetWorkLogReportForProjectsAndUser(
	jiraRequestService *connect.JiraRequestService,
	projectKeys string,
	userName string,
) (
	*[]models.WorkLogReportItem,
	error,
) {
	jiraProjectsRequestService := projects.NewJiraProjectsRequestService(jiraRequestService)
	jiraSearchRequestService := issues.NewJiraSearchRequestService(jiraRequestService)
	jiraIssueRequestService := issues.NewJiraIssueRequestService(jiraRequestService)

	if projectKeys == "all" {
		projectKeys = ""
		// get all projects
		fetchedProjects, fetchProjectErr := jiraProjectsRequestService.GetAllProjects()
		if fetchProjectErr != nil {
			return nil, fetchProjectErr
		}

		// Accumulate project keys for issue search request
		for i, project := range *fetchedProjects {
			projectKeys += project.Key
			if i+1 < len(*fetchedProjects) {
				projectKeys += ", "
			}
		}
	}

	// Get all issues for the project keys
	issues, fetchIssuesErr := jiraSearchRequestService.GetAllIssuesForProject(projectKeys)
	if fetchIssuesErr != nil {
		return nil, fetchIssuesErr
	}

	return jiraIssueRequestService.GenerateWorkLogReportForIssues(issues, userName)
}
