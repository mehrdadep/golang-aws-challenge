package main

import (
	"encoding/json"
	"fmt"
	"golang-aws-challenge/functions"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler for add model request
func Handler(request functions.Request) (functions.Response, error) {

	m := functions.Model{}
	if err := json.Unmarshal([]byte(request.Body), &m.X); err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"JSON Parse error\",\"details\":\"%s\"}", err.Error()), 500)
	}

	// Check for name field in request body
	if n, ok := m.X["name"].(string); ok {
		m.Name = string(n)
	} else {
		return functions.ReturnResponse("{\"message\":\"Name in request body is required\"}", 400)
	}

	// Check if model exists by name
	tempModel, err := functions.FindModelByName(m.Name)
	if err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
	}
	if tempModel != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"a Model with this name exists\", \"Model ID\": \"%s\"}", tempModel.ID), 200)
	}

	tempModel, err = functions.CreateModel(m.Name)

	if err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
	}
	return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"New Model created successfully\", \"Model ID\": \"%s\"}", tempModel.ID), 201)

}

func main() {
	lambda.Start(Handler)
}
