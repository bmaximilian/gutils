package connect

import (
	"crypto/tls"
	"errors"
	"github.com/bmaximilian/gutils/pkg/util/logger"
	"github.com/levigross/grequests"
	"net/http"
	"reflect"
	"strconv"
)

type JiraRequestServiceOptions struct {
	Tempo *TempoOptions
}

type JiraRequestService struct {
	baseUrl               string
	DefaultEndpointPrefix string
	authorizationToken    string
	requestOptions        *grequests.RequestOptions
	JiraOptions           *JiraRequestServiceOptions
}

func NewJiraRequestService(config *JiraServerConfig) (*JiraRequestService, error) {
	tlsConfig := tls.Config{}

	if config.TlsConfig.CertPath != "" {
		cert, err := tls.LoadX509KeyPair(config.TlsConfig.CertPath, config.TlsConfig.KeyPath)
		if err != nil {
			return nil, err
		}

		tlsConfig.Certificates = []tls.Certificate{cert}

		tlsConfig.BuildNameToCertificate()
	}

	return &JiraRequestService{
		baseUrl:               config.Url,
		DefaultEndpointPrefix: "/rest/api/" + strconv.Itoa(config.APIVersion),
		requestOptions: &grequests.RequestOptions{
			HTTPClient: &http.Client{
				Transport: &http.Transport{TLSClientConfig: &tlsConfig},
			},
			Headers: map[string]string{
				"Authorization": "Basic " + config.Token,
				"Content-Type":  "application/json",
			},
		},
		JiraOptions: &JiraRequestServiceOptions{
			Tempo: config.Tempo,
		},
	}, nil
}

func (j *JiraRequestService) assignOptions(options *grequests.RequestOptions) *grequests.RequestOptions {
	assignedOptions := *j.requestOptions

	if !reflect.DeepEqual(options.Params, assignedOptions.Params) {
		assignedOptions.Params = options.Params
	}

	if !reflect.DeepEqual(options.RequestBody, assignedOptions.RequestBody) {
		assignedOptions.RequestBody = options.RequestBody
	}

	if !reflect.DeepEqual(options.Data, assignedOptions.Data) {
		assignedOptions.Data = options.Data
	}

	if val, ok := options.Headers["Authorization"]; ok && val != assignedOptions.Headers["Authorization"] {
		assignedOptions.Headers["Authorization"] = val
	}

	return &assignedOptions
}

func (j *JiraRequestService) getUrl(endpoint string) string {
	return j.baseUrl + j.DefaultEndpointPrefix + endpoint
}

func (j *JiraRequestService) handleResponseReceived(url string, options *grequests.RequestOptions, response *grequests.Response, responseErr error) (*grequests.Response, error) {
	if responseErr != nil {
		return response, responseErr
	}

	if response != nil && response.StatusCode > 399 {
		logger.GetLogger().Warningln(response.String())
		return response, errors.New(url + " responded with status code " + strconv.Itoa(response.StatusCode))
	}

	return response, nil
}

func (j *JiraRequestService) Get(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	url := j.getUrl(endpoint)
	assignedOptions := j.assignOptions(options)
	response, err := grequests.Get(url, assignedOptions)

	return j.handleResponseReceived(url, assignedOptions, response, err)
}

func (j *JiraRequestService) Post(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	url := j.getUrl(endpoint)
	assignedOptions := j.assignOptions(options)
	response, err := grequests.Post(url, assignedOptions)

	return j.handleResponseReceived(url, assignedOptions, response, err)
}

func (j *JiraRequestService) Put(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	url := j.getUrl(endpoint)
	assignedOptions := j.assignOptions(options)
	response, err := grequests.Put(url, assignedOptions)

	return j.handleResponseReceived(url, assignedOptions, response, err)
}

func (j *JiraRequestService) Patch(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	url := j.getUrl(endpoint)
	assignedOptions := j.assignOptions(options)
	response, err := grequests.Patch(url, assignedOptions)

	return j.handleResponseReceived(url, assignedOptions, response, err)
}

func (j *JiraRequestService) Delete(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	url := j.getUrl(endpoint)
	assignedOptions := j.assignOptions(options)
	response, err := grequests.Delete(url, assignedOptions)

	return j.handleResponseReceived(url, assignedOptions, response, err)
}
