package connect

import (
	"crypto/tls"
	"github.com/levigross/grequests"
	"net/http"
	"reflect"
	"strconv"
)

type JiraRequestService struct {
	baseUrl               string
	defaultEndpointPrefix string
	authorizationToken    string
	requestOptions        *grequests.RequestOptions
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
		defaultEndpointPrefix: "/rest/api/" + strconv.Itoa(config.APIVersion),
		requestOptions: &grequests.RequestOptions{
			HTTPClient: &http.Client{
				Transport: &http.Transport{TLSClientConfig: &tlsConfig},
			},
			Headers: map[string]string{
				"Authorization": "Basic " + config.Token,
				"Content-Type":  "application/json",
			},
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
	return j.baseUrl + j.defaultEndpointPrefix + endpoint
}

func (j *JiraRequestService) Get(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	return grequests.Get(j.getUrl(endpoint), j.assignOptions(options))
}

func (j *JiraRequestService) Post(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	return grequests.Post(j.getUrl(endpoint), j.assignOptions(options))
}

func (j *JiraRequestService) Put(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	return grequests.Put(j.getUrl(endpoint), j.assignOptions(options))
}

func (j *JiraRequestService) Patch(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	return grequests.Patch(j.getUrl(endpoint), j.assignOptions(options))
}

func (j *JiraRequestService) Delete(endpoint string, options *grequests.RequestOptions) (*grequests.Response, error) {
	return grequests.Delete(j.getUrl(endpoint), j.assignOptions(options))
}
