package model

type Scenario struct {
	Id            int
	EndpointId    int
	Method        Method
	RequestFilter RequestFilter
	Response      Response
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
	GET Method = iota
	POST
	PUT
	DELETE
	PATCH
)
