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

	var body map[string]*json.RawMessage
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "JSON Parse error", err.Error()), 500)
	}

	var name string
	err = json.Unmarshal(*body["name"], &name)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Name in request body is required\"}", 400)
	}

	var note string
	err = json.Unmarshal(*body["note"], &note)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Note in request body is required\"}", 400)
	}

	var serial string
	err = json.Unmarshal(*body["serial"], &serial)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"Serial in request body is required\"}", 400)
	}

	var deviceModel string
	err = json.Unmarshal(*body["deviceModel"], &deviceModel)
	if err != nil {
		return functions.ReturnResponse("{\"message\":\"deviceModel in request is required\"}", 400)
	}

	// Create model from request body
	tempModel, err := functions.FindModelByID(deviceModel)
	if tempModel == nil {
		return functions.ReturnResponse("{\"message\":\"modelDevice id not found!\",\"details\":\"POST name to /devicemodels to get model id\"}", 400)
	}
	// Create device from request body
	tempDevice, err := functions.FindDeviceBySerial(serial)
	if tempDevice == nil {
		deviceID, err := uuid.NewUUID()
		if err != nil {
			return functions.ReturnResponse(fmt.Sprintf("{\"message\":\"%s\"}", err.Error()), 500)
		}

		device := functions.Device{
			ID:          deviceID.String(),
			Name:        name,
			Note:        note,
			Serial:      serial,
			DeviceModel: deviceModel,
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
	duplicateMessage := fmt.Sprintf("{\"message\":\"Device with this serial already exists\", \"Serial\": \"%s\"}", serial)
	return functions.ReturnResponse(duplicateMessage, 200)

}

func main() {
	lambda.Start(Handler)
}
