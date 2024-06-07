package request

import (
	"io"
	"net/http"
	"resq/pkg/auth"
	"resq/pkg/utils"
	"strings"

	"gopkg.in/yaml.v2"
)

type RequestSchema struct {
	Version string            `yaml:"version"`
	Name    string            `yaml:"name"`
	URL     string            `yaml:"url"`
	Method  string            `yaml:"method"`
	Path    map[string]string `yaml:"path"`
	Headers map[string]string `yaml:"headers"`
	Query   map[string]string `yaml:"query"`
	Body    *string           `yaml:"body"`
	Auth    *AuthSchema       `yaml:"auth"`
}

type AuthSchema struct {
	Type       string `yaml:"type"` // basic, api_key
	UserName   string `yaml:"username"`
	Password   string `yaml:"password"`
	Key        string `yaml:"key"`
	HeaderName string `yaml:"header_name"`
}

func ParseRequestSchema(data []byte) (RequestSchema, error) {
	var request RequestSchema
	err := yaml.Unmarshal(data, &request)

	if err != nil {
		return RequestSchema{}, err
	}

	return request, nil
}

func CreateRequest(request RequestSchema) (*http.Request, error) {
	// Initialize request body if it exists
	var reqBody io.Reader
	if request.Body != nil {
		reqBody = strings.NewReader(*request.Body)
	}

	// Create the HTTP request
	req, err := http.NewRequest(request.Method, request.URL, reqBody)
	if err != nil {
		return nil, err
	}

	// Set headers, query parameters, and path parameters
	setHeaders(req, request.Headers)
	setQueryParams(req, request.Query)
	setPathParams(req, request.Path)

	// Set authentication if needed
	if err := setAuth(req, request); err != nil {
		return nil, err
	}

	return req, nil
}

func ExecuteRequest(req *http.Request) (*http.Response, error) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func setHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}

func setQueryParams(req *http.Request, queryParams map[string]string) {
	query := req.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()
}
func setPathParams(req *http.Request, pathParams map[string]string) {
	for key, value := range pathParams {
		req.URL.Path = strings.Replace(req.URL.Path, "{"+key+"}", value, -1)
	}
}

func setAuth(req *http.Request, schema RequestSchema) error {
	if schema.Auth == nil {
		return nil
	}

	var authenticator auth.Auth
	switch schema.Auth.Type {
	case "basic":
		authenticator = auth.BasicAuth{
			Username: schema.Auth.UserName,
			Password: schema.Auth.Password,
		}
	case "api_key":
		authenticator = auth.APIKeyAuth{
			Key:        schema.Auth.Key,
			HeaderName: schema.Auth.HeaderName,
		}

	default:
		return utils.NewError("unsupported auth type: " + schema.Auth.Type)
	}

	for key, value := range authenticator.Header() {
		req.Header.Add(key, value)
	}

	return nil
}
