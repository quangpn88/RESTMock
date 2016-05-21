package model

type Endpoint struct {
	Id        int
	ProjectId int
	Path      string
}

func (endpoint *Endpoint) createScenarioGET(headerFilters []KeyValue, queryParameters []KeyValue, httpCode int, responseHeaders []KeyValue, responseBody string) Scenario {
	scenario := Scenario{}
	scenario.EndpointId = endpoint.Id
	scenario.Method = GET
	requestFilter := RequestFilter{}
	requestFilter.HeaderFilters = headerFilters
	requestFilter.QueryParameters = queryParameters
	scenario.RequestFilter = requestFilter
	response := Response{}
	response.Body = responseBody
	response.Header = responseHeaders
	response.HttpCode = httpCode
	return scenario
}
