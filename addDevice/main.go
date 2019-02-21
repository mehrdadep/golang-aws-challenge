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

// Handler for add device request
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var body map[string]*json.RawMessage
	err := json.Unmarshal([]byte(request.Body), &body)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "JSON Parse error", err.Error()), StatusCode: 500}, nil
	}

	var name string
	err = json.Unmarshal(*body["name"], &name)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Name in request body is required\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}

	var note string
	err = json.Unmarshal(*body["note"], &note)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Note in request body is required\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}

	var serial string
	err = json.Unmarshal(*body["serial"], &serial)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"Serial in request body is required\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}

	var deviceModel string
	err = json.Unmarshal(*body["deviceModel"], &deviceModel)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"deviceModel in request is required\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}

	// Create model from request body
	tempModel, err := functions.FindModelByID(deviceModel)
	if tempModel == nil {
		return events.APIGatewayProxyResponse{Body: "{\"message\":\"modelDevice id not found!\",\"details\":\"post to /devicemodels to get model id\"}",
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 400}, nil
	}
	// Create device from request body
	tempDevice, err := functions.FindDeviceBySerial(serial)
	if tempDevice == nil {
		deviceID, err := uuid.NewUUID()
		if err != nil {
			return events.APIGatewayProxyResponse{Body: fmt.Sprintf("{\"message\":\"%s\"}", err.Error()),
				Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
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
			message := fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "Database connection creation error:", err.Error())
			return events.APIGatewayProxyResponse{Body: message,
				Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
		}

		_, err = svc.PutItem(input)

		if err != nil {
			message := fmt.Sprintf("{\"message\":\"%s\",\"details\":\"%s\"}", "Got error inserting data", err.Error())
			return events.APIGatewayProxyResponse{Body: message,
				Headers: map[string]string{"content-type": "application/json"}, StatusCode: 500}, nil
		}
		successMessage := fmt.Sprintf("{\"message\":\"New Device created successfully\", \"Deive ID\": \"%s\"}", device.ID)
		return events.APIGatewayProxyResponse{Body: successMessage,
			Headers: map[string]string{"content-type": "application/json"}, StatusCode: 201}, nil
	}
	duplicateMessage := fmt.Sprintf("{\"message\":\"Device with this serial already exists\", \"Serial\": \"%s\"}", serial)
	return events.APIGatewayProxyResponse{Body: duplicateMessage,
		Headers: map[string]string{"content-type": "application/json"}, StatusCode: 200}, nil

}

func main() {
	lambda.Start(Handler)
}
