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
	var body map[string]*json.RawMessage
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"JSON Parse error\",\"details\":\"%s\"}", err.Error()), 500)
	}

	var name string
	err = json.Unmarshal(*body["name"], &name)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Name in request body is required\"}", 400)
	}
	tempModel, err := functions.FindModelByName(name)
	if err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
	}
	if tempModel == nil {
		modelID, err := uuid.NewUUID()
		if err != nil {
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String("Model"),
			Item: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(modelID.String()),
				},
				"Name": {
					S: aws.String(name),
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
		successMessage := fmt.Sprintf("{\"message\":\"New Model created successfully\", \"Model ID\": \"%s\"}", modelID.String())
		return functions.ReturnResponse(successMessage, 201)
	}
	duplicateMessage := fmt.Sprintf("{\"message\":\"a Model with this name exists\", \"Model ID\": \"%s\"}", tempModel.ID)
	return functions.ReturnResponse(duplicateMessage, 200)

}

func main() {
	lambda.Start(Handler)
}
