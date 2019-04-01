package issues

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"github.com/levigross/grequests"
)

type JiraIssueRequestService struct {
	requestService *connect.JiraRequestService
}

func NewJiraIssueRequestService(jiraRequestService *connect.JiraRequestService) *JiraIssueRequestService {
	return &JiraIssueRequestService{requestService: jiraRequestService}
}

func (j *JiraIssueRequestService) LoadWorkLogsForIssue(issue *models.Issue) (*models.Issue, error) {
	parsedResponse := models.WorkLogsResponse{}
	response, err := j.requestService.Get("/issue/"+issue.Key+"/worklog", &grequests.RequestOptions{})
	if err != nil {
		return nil, err
	}

	jsonErr := response.JSON(&parsedResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}

	issue.WorkLogs = parsedResponse.WorkLogs

	return issue, nil
}

func (j *JiraIssueRequestService) LoadWorkLogsForIssues(issues *[]models.Issue) (*[]models.Issue, error) {
	for i := 0; i < len(*issues); i += 1 {
		_, err := j.LoadWorkLogsForIssue(&(*issues)[i])
		if err != nil {
			return nil, err
		}
	}

	return issues, nil
}

func (j *JiraIssueRequestService) GenerateWorkLogReportForIssues(issues *[]models.Issue, userNameFilter string) (*[]models.WorkLogReportItem, error) {
	_, fetchWorkLogsError := j.LoadWorkLogsForIssues(issues)
	if fetchWorkLogsError != nil {
		return nil, fetchWorkLogsError
	}

	workLogReportItems := make([]models.WorkLogReportItem, 0)
	for _, issue := range *issues {
		workLogs := make([]models.WorkLog, 0)
		if userNameFilter != "" {
			workLogs = *issue.WorkLogsWithAuthor(userNameFilter)
		} else {
			workLogs = issue.WorkLogs
		}

		for _, workLog := range workLogs {
			workLogReportItems = append(workLogReportItems, models.WorkLogReportItem{
				IssueKey:          issue.Key,
				Started:           workLog.Started,
				Updated:           workLog.Updated,
				TimeSpent:         workLog.TimeSpent,
				TimeSpentSeconds:  workLog.TimeSpentSeconds,
				WorkLogId:         workLog.Id,
				IssueId:           issue.Id,
				AuthorShortName:   workLog.Author.Name,
				AuthorDisplayName: workLog.Author.DisplayName,
				Comment:           workLog.Comment,
			})
		}
	}

	return &workLogReportItems, nil
}
