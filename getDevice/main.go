package main

import (
	"encoding/json"
	"gochallenge/code/functions"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response Type
type Response events.APIGatewayProxyResponse

// Handler for get device request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ID := request.PathParameters["id"]
	tempDevice, err := functions.FindDeviceByID(ID)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Something went wrong\",\"details\":\"check path\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
	}
	if tempDevice == nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Device id not found!\",\"details\":\"id is invalid\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}
	body, err := json.Marshal(tempDevice)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Something went wrong\",\"details\":\"dabase error\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
	}
	return events.APIGatewayProxyResponse{Body: string(body),
		Headers: map[string]string{"content-type": "application/json"}, StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
