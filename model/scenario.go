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
	Header   []KeyValue
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
