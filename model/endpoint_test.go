package model

import "testing"

func TestCreateScenarioOfGetMethod(t *testing.T) {
	endpoint := Endpoint{Id: "aaa", ProjectId: "bbb", Path: "test"}
	var headerFilters []KeyValue
	headerFilters = append(headerFilters, KeyValue{"string", "test"})
	headerFilters = append(headerFilters, KeyValue{"number", 1})
	var queryParameters []KeyValue
	queryParameters = append(queryParameters, KeyValue{"string", "test"})
	queryParameters = append(queryParameters, KeyValue{"number", 1})
	httpCode := 200
	var responseHeaders []KeyValue
	responseHeaders = append(responseHeaders, KeyValue{"test", "test"})
	responseBody := "hello"
	name := "name"
	description := "description"
	scenario := endpoint.createScenarioGET(headerFilters, queryParameters, httpCode, responseHeaders, responseBody, name, description)
	if scenario.Method != GET {
		t.Error("method must be GET")
	}
	if scenario.EndpointId != "aaa" {
		t.Error("endpointId must be 1")
	}
	if scenario.Id == "" {
		t.Error("id must be set")
	}
	if scenario.Name != "name" {
		t.Error("Name must be set")
	}
	if scenario.Description != "description" {
		t.Error("Description must be set")
	}
}

func TestCreateScenarioOfPostMethod(t *testing.T)  {
	
}
