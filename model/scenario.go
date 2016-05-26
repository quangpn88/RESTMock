package model

type Scenario struct {
	Id            string
	EndpointId    string
	Method        Method
	RequestFilter RequestFilter
	Response      Response
	Name          string
	Description   string
}

type ScenarioBuilder struct {
	id               string
	endpointId       string
	method           Method
	name             string
	description      string
	headerFilters    []KeyValue
	queryParameters  []KeyValue
	bodyFilter       string
	responseHttpCode int
	responseHeaders  []KeyValue
	responseBody     string
}

func (scenarioBuilder *ScenarioBuilder) Id(id string) *ScenarioBuilder {
	scenarioBuilder.id = id
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) EndpointId(endpointId string) *ScenarioBuilder {
	scenarioBuilder.endpointId = endpointId
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) Method(method Method) *ScenarioBuilder {
	scenarioBuilder.method = method
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) HeaderFilters(headerFilters []KeyValue) *ScenarioBuilder {
	scenarioBuilder.headerFilters = headerFilters
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) QueryParameters(queryParameters []KeyValue) *ScenarioBuilder {
	scenarioBuilder.queryParameters = queryParameters
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) BodyFilter(bodyFilter string) *ScenarioBuilder {
	scenarioBuilder.bodyFilter = bodyFilter
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) ResponseHttpCode(responseHttpCode int) *ScenarioBuilder {
	scenarioBuilder.responseHttpCode = responseHttpCode
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) ResponseHeaders(responseHeaders []KeyValue) *ScenarioBuilder {
	scenarioBuilder.responseHeaders = responseHeaders
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) ResponseBody(responseBody string) *ScenarioBuilder {
	scenarioBuilder.responseBody = responseBody
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) Name(name string) *ScenarioBuilder {
	scenarioBuilder.name = name
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) Description(description string) *ScenarioBuilder {
	scenarioBuilder.description = description
	return scenarioBuilder
}
func (scenarioBuilder *ScenarioBuilder) BuildScenario() (scenario Scenario, error error) {
	scenario = Scenario{Id: scenarioBuilder.id,
		EndpointId:  scenarioBuilder.endpointId,
		Method:      scenarioBuilder.method,
		Name:        scenarioBuilder.name,
		Description: scenarioBuilder.description,
		RequestFilter: RequestFilter{HeaderFilters: scenarioBuilder.headerFilters,
			QueryParameters: scenarioBuilder.queryParameters,
			Body:            scenarioBuilder.bodyFilter},
		Response: Response{HttpCode: scenarioBuilder.responseHttpCode,
			Headers: scenarioBuilder.responseHeaders,
			Body:    scenarioBuilder.responseBody}}
	return
}

type RequestFilter struct {
	HeaderFilters   []KeyValue
	QueryParameters []KeyValue
	Body            string
}

type KeyValue struct {
	Key   string
	Value interface{}
}

type Response struct {
	HttpCode int
	Headers  []KeyValue
	Body     string
}

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	PATCH  Method = "PATCH"
)
