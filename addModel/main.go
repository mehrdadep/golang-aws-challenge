package main

import (
	"encoding/json"
	"fmt"
	"gochallenge/code/functions"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

// Response Type
type Response events.APIGatewayProxyResponse

// Handler for add model request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var body map[string]*json.RawMessage
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "JSON Parse error", err.Error()), StatusCode: 400}, nil
	}

	var name string
	err = json.Unmarshal(*body["name"], &name)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Name in request body is required\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}
	tempModel, err := functions.FindModelByName(name)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("{\"message\":\"%s\"}", err.Error()),
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
	}
	if tempModel == nil {
		modelID, err := uuid.NewUUID()
		if err != nil {
			return events.APIGatewayProxyResponse{Body: fmt.Sprintf("{\"message\":\"%s\"}", err.Error()),
				Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
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
			message := fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "Database connection creation error:", err.Error())
			return events.APIGatewayProxyResponse{Body: message,
				Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
		}

		_, err = svc.PutItem(input)
		if err != nil {
			message := fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "Got error calling PutItem:", err.Error())
			return events.APIGatewayProxyResponse{Body: message,
				Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
		}

		successMessage := fmt.Sprintf("{\"message\":\"New Model created successfully\", \"Model ID\": \"%s\"}", modelID.String())
		return events.APIGatewayProxyResponse{Body: successMessage,
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 201}, nil
	}
	successMessage := fmt.Sprintf("{\"message\":\"a Model with this name exists\", \"Model ID\": \"%s\"}", tempModel.ID)
	return events.APIGatewayProxyResponse{Body: successMessage,
		Headers: map[string]string{"content-type": "application/json"}, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
