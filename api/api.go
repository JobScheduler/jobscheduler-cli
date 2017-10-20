package api

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/parnurzeal/gorequest"
)

const (
	BaseURL        = "https://jobscheduler.xyz/api/v1"
	RouteJobsGET   = "/jobs"
	RouteJobsIDGET = "/jobs/%v"
)

type Job struct {
	ID                    string            `json:"id"`
	Name                  string            `json:"name"`
	Tag                   string            `json:"tag,omitempty"`
	CronRule              string            `json:"cron_rule"`
	IsActive              bool              `json:"is_active"`
	ActionType            string            `json:"action_type"`
	ActionMethod          string            `json:"action_method"`
	ActionURL             string            `json:"action_url"`
	ActionQueryParameters interface{}       `json:"action_query_parameters"`
	ActionHeaders         map[string]string `json:"action_headers"`
	ActionBody            interface{}       `json:"action_body"`
	Timezone              string            `json:"timezone,omitempty"`
	StartTime             string            `json:"start_time,omitempty"`
}

type APIClient struct {
	APIKey  string `json:"api_key"`
	request *gorequest.SuperAgent
}

func New(apiKey string) *APIClient {
	request := gorequest.New()
	c := APIClient{apiKey, request}
	return &c
}

func (c *APIClient) requestGET(route string) (body string, err error) {
	_, body, errs := c.request.Get(BaseURL+route).Set("Authorization", c.APIKey).End()
	if errs != nil {
		return "", errors.New("lol")
	}
	return
}

func (c *APIClient) requestPUT(route string, data Job) (body string, err error) {
	_, body, errs := c.request.Put(BaseURL+route).
		Set("Authorization", c.APIKey).
		Send(data).
		End()
	if errs != nil {
		return "", errors.New("lol")
	}
	return
}

func (c *APIClient) GetJobs() []Job {
	res, _ := c.requestGET(RouteJobsGET)

	var jobs []Job
	err := json.Unmarshal([]byte(res), &jobs)
	if err != nil {
		panic(err)
	}

	return jobs
}

func (c *APIClient) GetJobByID(id string) Job {
	res, _ := c.requestGET(fmt.Sprintf(RouteJobsIDGET, id))

	var job Job
	err := json.Unmarshal([]byte(res), &job)
	if err != nil {
		panic(err)
	}

	return job
}

func (c *APIClient) UpdateJob(newJob Job) Job {
	res, _ := c.requestPUT(fmt.Sprintf(RouteJobsIDGET, newJob.ID), newJob)

	var job Job
	err := json.Unmarshal([]byte(res), &job)
	if err != nil {
		panic(err)
	}

	return job
}
