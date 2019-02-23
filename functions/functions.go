package functions

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/google/uuid"
)

// A Model contains model name
type Model struct {
	ID   string                 `json:"id"`
	Name string                 `json:"name"`
	X    map[string]interface{} `json:"-"`
}

// A Device contains name, serial, note and deviceModel
type Device struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Serial      string                 `json:"serial"`
	Note        string                 `json:"note"`
	DeviceModel string                 `json:"deviceModel"`
	X           map[string]interface{} `json:"-"`
}

// Response Type
type Response events.APIGatewayProxyResponse

// Request Type
type Request events.APIGatewayProxyRequest

// FindModelByID find a model based on it's ID
// return Model item if exists
// return nil if not
func FindModelByID(id string) (*Model, error) {
	svc, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	filt := expression.Name("ID").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Model"),
	}
	result, err := svc.Scan(params)

	if err != nil || len(result.Items) == 0 {
		return nil, nil
	}
	item := Model{}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

// FindModelByName Find a model based on it's name
// return Model item if exists
// return nil if not
func FindModelByName(name string) (*Model, error) {
	svc, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	filt := expression.Name("Name").Equal(expression.Value(name))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Model"),
	}
	result, err := svc.Scan(params)
	if err != nil || len(result.Items) == 0 {
		return nil, nil
	}
	item := Model{}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)

	if err != nil {
		return nil, err
	}

	return &item, nil
}

// FindDeviceBySerial find a device based on it's serial
// return Device item if exists
// return nil if not
func FindDeviceBySerial(serial string) (*Device, error) {
	svc, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	filt := expression.Name("Serial").Equal(expression.Value(serial))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"),
		expression.Name("Serial"), expression.Name("Note"), expression.Name("deviceModel"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Device"),
	}
	result, err := svc.Scan(params)

	if err != nil || len(result.Items) == 0 {
		return nil, nil
	}
	item := Device{}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// FindDeviceByID find a device based on it's id
// return Device item if exists
// return nil if not
func FindDeviceByID(id string) (*Device, error) {
	svc, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	filt := expression.Name("ID").Equal(expression.Value(id))
	proj := expression.NamesList(expression.Name("ID"), expression.Name("Name"),
		expression.Name("Serial"), expression.Name("Note"), expression.Name("deviceModel"))
	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Device"),
	}
	result, err := svc.Scan(params)

	if err != nil || len(result.Items) == 0 {
		return nil, nil
	}
	item := Device{}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

//CreateDevice create a new device in the database
func CreateDevice(name string, note string, serial string, deviceModel string) (*Device, error) {
	deviceID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	// Create device from request body
	device := Device{
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

	svc, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &device, nil

}

//CreateModel create a new model based on a name
func CreateModel(name string) (*Model, error) {
	modelID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
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
	svc, err := ConnectDB()
	if err != nil {
		return nil, err
	}

	_, err = svc.PutItem(input)
	if err != nil {
		return nil, err
	}
	newModel := Model{
		Name: name,
		ID:   modelID.String(),
	}

	return &newModel, nil

}

// ConnectDB connect to database and return db
func ConnectDB() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{
		// This region is set to France (Paris)
		Region: aws.String("eu-west-3")},
	)
	if err != nil {
		return nil, err
	}
	// Create DynamoDB client
	svc := dynamodb.New(sess)
	return svc, nil
}

// ReturnResponse resturns json body with selected status code
func ReturnResponse(body string, code int) (Response, error) {
	return Response{Body: body, Headers: map[string]string{"content-type": "application/json"}, StatusCode: code}, nil
}
