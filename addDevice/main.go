package main

import (
	"encoding/json"
	"fmt"
	"golang-aws-challenge/functions"

	"github.com/aws/aws-lambda-go/lambda"
)

// Handler for add device request
func Handler(request functions.Request) (functions.Response, error) {

	d := functions.Device{}
	if err := json.Unmarshal([]byte(request.Body), &d.X); err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"JSON Parse error\",\"details\":\"%s\"}", err.Error()), 500)
	}

	//Check for fields one by one

	if n, ok := d.X["name"].(string); ok {
		d.Name = string(n)
	} else {
		return functions.ReturnResponse("{\"message\":\"Name in request body is required\"}", 400)
	}

	if n, ok := d.X["note"].(string); ok {
		d.Note = string(n)
	} else {
		return functions.ReturnResponse("{\"message\":\"Note in request body is required\"}", 400)
	}

	if n, ok := d.X["serial"].(string); ok {
		d.Serial = string(n)
	} else {
		return functions.ReturnResponse("{\"message\":\"Serial in request body is required\"}", 400)
	}

	if n, ok := d.X["deviceModel"].(string); ok {
		d.DeviceModel = string(n)
	} else {
		return functions.ReturnResponse("{\"message\":\"deviceModel in request body is required\"}", 400)
	}

	// Check model id from request body
	tempModel, err := functions.FindModelByID(d.DeviceModel)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Error in finding model device!\",\"details\":\"Database error\"}", 400)
	}
	if tempModel == nil {
		return functions.ReturnResponse("{\"message\":\"modelDevice id not found!\",\"details\":\"POST name to /devicemodels to get model id\"}", 400)
	}

	// Check for device serial
	tempDevice, err := functions.FindDeviceBySerial(d.Serial)
	if tempDevice != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"Device with this serial already exists\", \"Serial\": \"%s\"}", d.Serial), 200)
	}
	tempDevice, err = functions.CreateDevice(d.Name, d.Note, d.Serial, d.DeviceModel)

	if err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
	}
	return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"New Device created successfully\", \"Deive ID\": \"%s\"}", tempDevice.ID), 201)

}

func main() {
	lambda.Start(Handler)
}
