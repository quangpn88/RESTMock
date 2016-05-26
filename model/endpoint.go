package model

import "github.com/satori/go.uuid"

type Endpoint struct {
	Id        string
	ProjectId string
	Path      string
}

func (endpoint *Endpoint) createScenarioGET(headerFilters []KeyValue, queryParameters []KeyValue, httpCode int, responseHeaders []KeyValue, responseBody string, name string, description string) (scenario Scenario) {
	scenarioBuilder := ScenarioBuilder{}
	scenario, _ = scenarioBuilder.
		Id(uuid.NewV4().String()).
		EndpointId(endpoint.Id).
		Method(GET).Name(name).
		Description(description).
		HeaderFilters(headerFilters).
		QueryParameters(queryParameters).
		ResponseBody(responseBody).
		ResponseHeaders(responseHeaders).
		ResponseHttpCode(httpCode).
		BuildScenario()
	return
}
