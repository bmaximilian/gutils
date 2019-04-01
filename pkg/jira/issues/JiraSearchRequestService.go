package issues

import (
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/bmaximilian/gutils/pkg/jira/issues/models"
	"github.com/levigross/grequests"
	"strconv"
)

type JiraSearchRequestService struct {
	requestService *connect.JiraRequestService
}

func NewJiraSearchRequestService(jiraRequestService *connect.JiraRequestService) *JiraSearchRequestService {
	return &JiraSearchRequestService{requestService: jiraRequestService}
}

func (j *JiraSearchRequestService) GetAllIssuesForProject(projectKey string) (*[]models.Issue, error) {
	parsedIssues := make([]models.Issue, 0)

	numberOfRequests := 1
	startAt := 0
	for numberOfRequests > 0 {
		parsedResponse := models.SearchResponse{}

		response, err := j.requestService.Get("/search", &grequests.RequestOptions{
			Params: map[string]string{
				"jql":     "project in (" + projectKey + ") ORDER BY key ASC",
				"startAt": strconv.Itoa(startAt),
			},
		})
		if err != nil {
			return nil, err
		}

		jsonErr := response.JSON(&parsedResponse)
		if jsonErr != nil {
			return nil, err
		}

		parsedIssues = append(parsedIssues, parsedResponse.Issues...)

		if startAt == 0 && parsedResponse.Total > 0 {
			numberOfRequests = int(parsedResponse.Total) / int(parsedResponse.MaxResults)
		} else {
			numberOfRequests -= 1
		}

		startAt += int(parsedResponse.MaxResults)
	}

	return &parsedIssues, nil
}
