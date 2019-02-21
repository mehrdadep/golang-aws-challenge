package main

import (
	"fmt"
	"gochallenge/code/functions"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response Type
type Response events.APIGatewayProxyResponse

// Handler for get model request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ID := request.PathParameters["id"]
	tempModel, err := functions.FindModelByID(ID)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Something went wrong\",\"details\":\"check path\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
	}
	if tempModel == nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Model id not found!\",\"details\":\"id is invalid\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}
	successMessage := fmt.Sprintf("{\"id\":\"%s\", \"name\": \"%s\"}", tempModel.ID, tempModel.Name)
	return events.APIGatewayProxyResponse{Body: successMessage,
		Headers: map[string]string{"content-type": "application/json"}, StatusCode: 200}, nil

}

func main() {
	lambda.Start(Handler)
}
