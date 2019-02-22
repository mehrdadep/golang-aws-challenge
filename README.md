# Golang AWS challenge

## Challenge
*Implement a simple Restful API on AWS using the following tech stack:*
- Serverless Framework (https://serverless.com/) 
- Go language (https://golang.org/) 
- AWS API Gateway
- AWS Lambda
- AWS DynamoDB

###### Challenge apis
The API should accept the following JSON requests and produce the corresponding HTTP responses:
```
Request 1:
HTTP POST
URL: https://<api-gateway-url>/api/devices
Body (application/json):
{
  "id": "/devices/id1",
  "deviceModel": "/devicemodels/id1",
  "name": "Sensor",
  "note": "Testing a sensor.",
  "serial": "A020000102"
}
```

```
Request 2:
HTTP GET
URL: https://<api-gateway-url>/api/devices/{id}
Example: GET https://api123.amazonaws.com/api/devices/id1
```

## Setup
###### Pre-requisites
- Install [nodejs](https://nodejs.org/en/download/)
- Install [Go languge](https://golang.org/doc/install)
- Install [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)
- Install serverless framework using `npm install -g serverless`

###### Dependencies
Place this repo  into the `src` folder of `${GOPATH}` or add it to your `${PATH}`. Note that the folder name must be `golang-aws-challenge`, otherwise you have to deal with package imports manually. This repo needs the following packages in golang vendors to work correctly:
- `github.com/aws/aws-lambda-go/*`
- `github.com/aws/aws-sdk-go/*`
- `github.com/google/uuid`
###### Configure AWS & Deploy
1. Configure aws cli (use your access key and secret key):
`aws configure`
Enter values as follows:
```
AWS Access Key ID [None]: your-access-key
AWS Secret Access Key [None]: your-secret-key
Default region name [None]: eu-west-3
Default output format [None]: json
```
2. Create two tables in the database named `Model` and `Device`. These tables store information about models and devices. You can create these tables using `aws-cli` with the following commands:
```
aws dynamodb create-table --table-name Device --attribute-definitions \
AttributeName=ID,AttributeType=S \
--key-schema AttributeName=ID,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
```
```
aws dynamodb create-table --table-name Model --attribute-definitions \
AttributeName=ID,AttributeType=S --key-schema AttributeName=ID,KeyType=HASH \
--provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5
```
2. `POST` methods are protected by api-key and you should generate an api key and use it as the value of the `x-api-key` key in request header. To do so, edit `serverless.yml` file and set `stage` and `apiKeys` for the first time. An api key will be generated after the first deploy and `stage` will be in the api's final url as follows:
<pre>
https://e7rjun495i.execute-api.eu-west-3.amazonaws.com/<b>stage</b>/api/devices
</pre>

3. All neccessary commands are called using `make` command (Linux64-only). Alternatively you can run the following commands one by one:
<pre>
build clean deploy
dep ensure -v
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addModel addModel/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addDevice addDevice/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getDevice getDevice/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getModel getModel/main.go
rm -rf ./bin ./vendor Gopkg.lock
clean build
sls deploy --verbose
</pre>

You can see the output of a successful deploy on amazon servers.
![Terminal Output](img/terminal.png?raw=true "Terminal Output")

## Routes
There are four final routes in this api as follows:
1. `POST` on `/api/devices` will add a new device to database. Payload should be plain text (`JSON`). The api also checks for duplicate serial numbers and prevents you from storing a device twice. `deviceModel` should be a valid `ID` of a model. you can get a valid `ID` of a model from part 2 of this section. The result of this call is a new id for a new device. Remember this is a protected url, so you have to set api key in `x-api-key` key of the header. `JSON` payload should be formated as follows:
```
{
	"name": "Test name",
	"serial": "A205ad0500",
	"deviceModel":"ee230a7b-3615-11e9-88d9-2288fa44503c9",
	"note": "Test note"
}
```

2. `POST` on `/api/devicemodels` will add a new model to database. Payload should be plain text (`JSON`). The api also checks for duplicate names and prevents you from storing a model twice. In any successful case a call will return the model `ID`. Remember this is also a protected url, so you have to set api key in `x-api-key` key of the header. `JSON` payload should be formated as follows:
```
{
	"name": "Test name"
}
```

3. `GET` on `/api/devices/{id}` will return the device information in `JSON` format. The `{id}` path variable should be replaced with a valid id. If the device exsits, the call will return the device name and other details, otherwise the result will be a `404 Error`

4. `GET` on `/api/devicemodels/{id}` will return the model information in `JSON` format. The `{id}` path variable should be replaced with a valid id. If the model exsits, the call will return the model name and id, otherwise the result will be a `404 Error`
## Test

## Maintainer
Mehrdad Esmaeilpour:
- [LinkedIn](https://www.linkedin.com/in/mehrdadep/)
- [Twitter](https://twitter.com/mehrdadep)
- [Stackoverflow](https://stackoverflow.com/users/8844510/mehrdadep)