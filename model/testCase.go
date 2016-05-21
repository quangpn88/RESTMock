package model

type TestCase struct {
	Id            int
	EndpointId    int
	RequestFilter RequestFilter
}

type RequestFilter struct {
	HeaderFilters   []KeyValueFilter
	QueryParameters []KeyValueFilter
	Body            string
}

type KeyValueFilter struct {
	Key   string
	Value interface{}
}

type Response struct {
	Body string
}
