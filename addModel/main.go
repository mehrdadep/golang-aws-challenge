package main

import (
	"encoding/json"
	"fmt"
	"golang-aws-challenge/functions"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
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
	if tempModel == nil {
		modelID, err := uuid.NewUUID()
		if err != nil {
			// Error creating uuid
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String("Model"),
			Item: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(modelID.String()),
				},
				"Name": {
					S: aws.String(m.Name),
				},
			},
		}
		svc, err := functions.ConnectDB()
		if err != nil {
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}",
				"Database connection creation error:", err.Error()), 500)

		}

		_, err = svc.PutItem(input)
		if err != nil {
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "Error inserting data", err.Error()), 500)
		}
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"New Model created successfully\", \"Model ID\": \"%s\"}", modelID.String()), 201)
	}
	return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"a Model with this name exists\", \"Model ID\": \"%s\"}", tempModel.ID), 200)

}

func main() {
	lambda.Start(Handler)
}
