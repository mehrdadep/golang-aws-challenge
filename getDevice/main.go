package main

import (
	"encoding/json"
	"golang-aws-challenge/functions"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler for get device request
func Handler(request functions.Request) (functions.Response, error) {

	// Read path variable
	ID := request.PathParameters["id"]

	// Find device by id
	tempDevice, err := functions.FindDeviceByID(ID)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Something went wrong\",\"details\":\"check path\"}", 500)
	}
	if tempDevice == nil {
		return functions.ReturnResponse("{\"message\":\"Device id not found!\",\"details\":\"id is invalid\"}", 400)
	}
	// convert into JSON format
	body, err := json.Marshal(tempDevice)
	return functions.ReturnResponse(string(body), 200)
}

func main() {
	lambda.Start(Handler)
}
