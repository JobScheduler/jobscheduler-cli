package api

import (
	"encoding/json"
	"errors"

	"github.com/parnurzeal/gorequest"
)

const (
	BaseURL = "https://jobscheduler.xyz/api/v1"
)

type Job struct {
	ID                    string      `json:"id"`
	Name                  string      `json:"name"`
	Tag                   string      `json:"tag"`
	CronRule              string      `json:"cron_rule"`
	IsActive              bool        `json:"is_active"`
	ActionType            string      `json:"action_type"`
	ActionMethod          string      `json:"action_method"`
	ActionURL             string      `json:"action_url"`
	ActionQueryParameters interface{} `json:"action_query_parameters"`
	ActionHeaders         interface{} `json:"action_headers"`
	ActionBody            interface{} `json:"action_body"`
	Timezone              string      `json:"timezone"`
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

func (c *APIClient) GetJobs() []Job {
	res, _ := c.requestGET("/jobs")

	var jobs []Job
	err := json.Unmarshal([]byte(res), &jobs)
	if err != nil {
		panic(err)
	}

	return jobs
}
