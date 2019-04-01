package projects

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/levigross/grequests"
)

type JiraProjectsRequestService struct {
	requestService *connect.JiraRequestService
}

func NewJiraProjectsRequestService(jiraRequestService *connect.JiraRequestService) *JiraProjectsRequestService {
	return &JiraProjectsRequestService{requestService: jiraRequestService}
}

func (j *JiraProjectsRequestService) GetAllProjects() (*grequests.Response, error) {
	return j.requestService.Get("/project", &grequests.RequestOptions{})
}

func (j *JiraProjectsRequestService) GetUpdatedWorklogs() (*grequests.Response, error) {
	return j.requestService.Get("/worklog/updated", &grequests.RequestOptions{})
}
