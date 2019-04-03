package tempo

import (
	"fmt"
	"github.com/bmaximilian/gutils/pkg/jira/connect"
	"github.com/levigross/grequests"
	"strconv"
	"time"
)

type TempoWorkLogRequestService struct {
	tempoApiVersion int
	requestService  *connect.JiraRequestService
}

func NewTempoWorkLogRequestService(jiraRequestService *connect.JiraRequestService) *TempoWorkLogRequestService {
	tempoRequestService := *jiraRequestService
	tempoRequestService.DefaultEndpointPrefix = "/rest/tempo-timesheets/" + strconv.Itoa(jiraRequestService.JiraOptions.Tempo.ApiVersion)

	return &TempoWorkLogRequestService{requestService: &tempoRequestService}
}

func (t *TempoWorkLogRequestService) GetWorkLogs(
	projectKey string,
	username string,
	dateFrom time.Time,
	dateTo time.Time,
	accountKey string,
	teamId string,
) (*grequests.Response, error) {
	res, err := t.requestService.Get("/worklogs", &grequests.RequestOptions{
		Params: map[string]string{
			"projectKey": projectKey,
			"username":   username,
			"dateFrom":   dateFrom.Format("2006-01-02"),
			"dateTo":     dateTo.Format("2006-01-02"),
			"accountKey": accountKey,
			"teamId":     teamId,
		},
	})

	fmt.Println(res)
	if err != nil {
		return nil, err
	}

	//jsonErr := response.JSON(&parsedResponse)
	//if jsonErr != nil {
	//	return nil, jsonErr
	//}

	return res, err
}
