package main

import (
	"encoding/json"
	"golang-aws-challenge/functions"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler for get model request
func Handler(request functions.Request) (functions.Response, error) {

	ID := request.PathParameters["id"]
	tempModel, err := functions.FindModelByID(ID)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Something went wrong\",\"details\":\"check path\"}", 500)
	}
	if tempModel == nil {
		return functions.ReturnResponse("{\"message\":\"Model id not found!\",\"details\":\"id is invalid\"}", 400)
	}
	body, err := json.Marshal(tempModel)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Something went wrong\",\"details\":\"dabase error\"}", 500)
	}
	return functions.ReturnResponse(string(body), 200)
}

func main() {
	lambda.Start(Handler)
}
