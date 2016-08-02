package screwdriver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// API is a Screwdriver API endpoint
type API interface {
	BuildFromID(buildID string) (Build, error)
	JobFromID(jobID string) (Job, error)
}

type api struct {
	url    string
	token  string
	client *http.Client
}

// New returns a new API object
func New(url, token string) (API, error) {
	api := api{
		url,
		token,
		&http.Client{},
	}
	return API(api), nil
}

// Job is a Screwdriver Job
type Job struct {
	ID         string `json:"id"`
	PipelineID string `json:"pipelineId"`
}

// Build is a Screwdriver Build
type Build struct {
	ID    string `json:"id"`
	JobID string `json:"jobId"`
}

func (a api) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Generating request to Screwdriver: %v", err)
	}
	token := fmt.Sprintf("Bearer %s", a.token)
	req.Header.Set("Authorization", token)

	response, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Reading response from Screwdriver: %v", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Reading response Body from Screwdriver: %v", err)
	}
	return body, nil
}

// BuildFromID fetches and returns a Build object from its ID
func (a api) BuildFromID(buildID string) (build Build, err error) {
	url := fmt.Sprintf("%s/builds/%s", a.url, buildID)
	body, err := a.get(url)
	if err != nil {
		return build, fmt.Errorf("Reading response Body from Screwdriver: %v", err)
	}

	err = json.Unmarshal(body, &build)
	if err != nil {
		return build, fmt.Errorf("Parsing JSON response %q: %v", body, err)
	}
	return build, nil
}

// BuildFromID fetches and returns a Build object from its ID
func (a api) JobFromID(jobID string) (job Job, err error) {
	url := fmt.Sprintf("%s/jobs/%s", a.url, jobID)
	body, err := a.get(url)
	if err != nil {
		return job, fmt.Errorf("Reading response Body from Screwdriver: %v", err)
	}

	err = json.Unmarshal(body, &job)
	if err != nil {
		return job, fmt.Errorf("Parsing JSON response %q: %v", body, err)
	}
	return job, nil
}
