package functions

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// A Model contains model name
type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// A Device contains name, serial, note and deviceModel
type Device struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Serial      string `json:"serial"`
	Note        string `json:"note"`
	DeviceModel string `json:"deviceModel"`
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

	if err != nil {
		return nil, nil
	}
	if len(result.Items) == 0 {
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
	if err != nil {
		return nil, nil
	}
	if len(result.Items) == 0 {
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

	if err != nil {
		return nil, nil
	}
	if len(result.Items) == 0 {
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
func FindDeviceByID(serial string) (*Device, error) {
	svc, err := ConnectDB()
	if err != nil {
		return nil, err
	}
	filt := expression.Name("ID").Equal(expression.Value(serial))
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

	if err != nil {
		return nil, nil
	}
	if len(result.Items) == 0 {
		return nil, nil
	}
	item := Device{}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
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
