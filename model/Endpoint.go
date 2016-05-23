package model

import "github.com/satori/go.uuid"

type Endpoint struct {
	Id        string
	ProjectId string
	Path      string
}

func (endpoint *Endpoint) createScenarioGET(headerFilters []KeyValue, queryParameters []KeyValue, httpCode int, responseHeaders []KeyValue, responseBody string, name string, description string) Scenario {
	scenario := Scenario{}
	scenario.Id = uuid.NewV4().String()
	scenario.EndpointId = endpoint.Id
	scenario.Method = GET
	scenario.Name = name
	scenario.Description = description

	requestFilter := RequestFilter{}
	requestFilter.HeaderFilters = headerFilters
	requestFilter.QueryParameters = queryParameters
	scenario.RequestFilter = requestFilter

	response := Response{}
	response.Body = responseBody
	response.Header = responseHeaders
	response.HttpCode = httpCode
	scenario.Response = response
	return scenario
}
