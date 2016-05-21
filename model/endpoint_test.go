package model

import "testing"

func TestCreateScenarioOfGetType(t *testing.T) {
	endpoint := Endpoint{Id: 1, ProjectId: 1, Path: "test"}
	var headerFilters []KeyValue
	headerFilters[0] = KeyValue{"string", "test"}
	headerFilters[0] = KeyValue{"number", 1}
	var queryParameters []KeyValue
	queryParameters[0] = KeyValue{"string", "test"}
	queryParameters[1] = KeyValue{"number", 1}
	httpCode := 200
	var responseHeader []KeyValue
	responseHeader[0] = KeyValue{"test", "test"}
	responseBody := "hello"
	scenario := endpoint.createScenarioGET(headerFilters, queryParameters, httpCode, responseHeader, responseBody)
	if scenario.Method != GET {
		t.Error("method must be GET")
	}
}
