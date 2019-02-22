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

// Handler for add device request
func Handler(request functions.Request) (functions.Response, error) {

	d := functions.Device{}
	if err := json.Unmarshal([]byte(request.Body), &d.X); err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"JSON Parse error\",\"details\":\"%s\"}", err.Error()), 500)
	}
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

	// Create model from request body
	tempModel, err := functions.FindModelByID(d.DeviceModel)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Error in finding model device!\",\"details\":\"Database error\"}", 400)
	}
	if tempModel == nil {
		return functions.ReturnResponse("{\"message\":\"modelDevice id not found!\",\"details\":\"POST name to /devicemodels to get model id\"}", 400)
	}
	// Create device from request body
	tempDevice, err := functions.FindDeviceBySerial(d.Serial)
	if tempDevice == nil {
		deviceID, err := uuid.NewUUID()
		if err != nil {
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
		}

		device := functions.Device{
			ID:          deviceID.String(),
			Name:        d.Name,
			Note:        d.Note,
			Serial:      d.Serial,
			DeviceModel: d.DeviceModel,
		}

		input := &dynamodb.PutItemInput{
			TableName: aws.String("Device"),
			Item: map[string]*dynamodb.AttributeValue{
				"ID": {
					S: aws.String(device.ID),
				},
				"Name": {
					S: aws.String(device.Name),
				},
				"Note": {
					S: aws.String(device.Note),
				},
				"Serial": {
					S: aws.String(device.Serial),
				},
				"deviceModel": {
					S: aws.String(device.DeviceModel),
				},
			},
		}

		svc, err := functions.ConnectDB()
		if err != nil {
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "Database connection creation error:", err.Error()), 500)
		}

		_, err = svc.PutItem(input)

		if err != nil {
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "Got error inserting data", err.Error()), 500)
		}
		successMessage := fmt.Sprintf("{\"message\":\"New Device created successfully\", \"Deive ID\": \"%s\"}", device.ID)
		return functions.ReturnResponse(successMessage, 201)
	}
	duplicateMessage := fmt.Sprintf("{\"message\":\"Device with this serial already exists\", \"Serial\": \"%s\"}", d.Serial)
	return functions.ReturnResponse(duplicateMessage, 200)

}

func main() {
	lambda.Start(Handler)
}
