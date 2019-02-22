# Golang AWS challenge

## Challenge
*Implement a simple Restful API on AWS using the following tech stack:*
- Serverless Framework (https://serverless.com/) 
- Go language (https://golang.org/) 
- AWS API Gateway
- AWS Lambda
- AWS DynamoDB

#### API
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
#### Pre-requisites
- Install [NodeJS](https://nodejs.org/en/download/)
- Install [Go languge](https://golang.org/doc/install)
- Install [aws-cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)
- Install serverless framework using `npm install -g serverless`

#### Dependencies
Place this repo  into the `src` folder of `${GOPATH}` or add it to your `${PATH}`. Note that the folder name must be `golang-aws-challenge`, otherwise you have to deal with package imports manually. This repo needs the following packages in golang vendors to work correctly:
- `github.com/aws/aws-lambda-go/*`
- `github.com/aws/aws-sdk-go/*`
- `github.com/google/uuid`
#### Configure AWS & Deploy
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

3. All neccessary commands are called using `make` command (Linux64-only). Alternatively (on Windows or mac) you can run the following commands one by one:
<pre>
dep ensure -v
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addModel addModel/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addDevice addDevice/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getDevice getDevice/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getModel getModel/main.go
sls deploy --verbose
</pre>

You can see the output of a successful deploy on amazon servers.

![Terminal Output](img/terminal.png?raw=true "Terminal Output")

## Routes
There are four final routes in this api as follows:

#### `POST` on `/api/devices`
This route will add a new device to database. Payload should be plain text (`JSON`). The api also checks for duplicate serial numbers and prevents you from storing a device twice. `deviceModel` should be a valid `ID` of a model. you can get a valid `ID` of a model from part 2 of this section. The result of this call is a new id for a new device. Remember this is a protected url, so you have to set api key in `x-api-key` key of the header. `JSON` payload should be formated as follows:

```
{
	"name": "Test name",
	"serial": "A205ad0500",
	"deviceModel":"ee230a7b-3615-11e9-88d9-2288fa44503c9",
	"note": "Test note"
}
```

The successful (Code: 201) result:
```
{
    "message": "New Device created successfully",
    "Deive ID": "e1d99eca-3684-11e9-be8b-e6039fb0c953"
}
```

If deviceModel id not exists (Code: 400):
```
{
    "message": "modelDevice id not found!",
    "details": "POST name to /devicemodels to get model id"
}
``` 

If serial number is duplicate (Code: 200):
```
{
    "message": "Device with this serial already exists",
    "Serial": "A205ad056500"
}
```

If a field is not provided (Code: 400):
```
{
    "message": "Serial in request body is required"
}
```

If `JSON` is malformed (Code: 500):
```
{
    "message": "JSON Parse error",
    "details": "invalid character '\n' in string literal"
}
```

If api key is not provided correctly (Code: 403):
```
{
    "message": "Forbidden"
}
```
####  `POST` on `/api/devicemodels` 
This route will add a new model to database. Payload should be plain text (`JSON`). The api also checks for duplicate names and prevents you from storing a model twice. In any successful case a call will return the model `ID`. Remember this is also a protected url, so you have to set api key in `x-api-key` key of the header. `JSON` payload should be formated as follows:
```
{
	"name": "Test name"
}
```

The successful result (Code: 201):
```
{
    "message": "New Model created successfully",
    "Model ID": "b6390810-3686-11e9-98d7-ca8618db46e6"
}
```

If a model with duplicate name exists (Code: 200):
```
{
    "message": "a Model with this name exists",
    "Model ID": "4d4a1973-3682-11e9-8c33-9ecfa7849824"
}
```

If `JSON` is malformed (Code: 500):
```
{
    "message": "JSON Parse error",
    "details": "invalid character '}' looking for beginning of object key string"
}
```

If name field is not provided (Code: 400):
```
{
    "message": "Name in request body is required"
}
```

If api key is not provided correctly (Code: 403):
```
{
    "message": "Forbidden"
}
```
####  `GET` on `/api/devices/{id}`
This route will return the device information in `JSON` format. The `{id}` path variable should be replaced with a valid id. If the device exsits, the call will return the device name and other details, otherwise the result will be a `404 Error`

The successful result (Code: 200):
```
{
    "id": "249cae6a-3619-11e9-872e-3ec11d02de65",
    "name": "Test device name",
    "serial": "A205ad050",
    "note": "This is a simple note",
    "deviceModel": "ee230a7b-3615-11e9-88d9-2288fa4453c9"
}
```

If device id not exists (Code: 400):
```
{
    "message": "Device id not found!",
    "details": "id is invalid"
}
```
#### `GET` on `/api/devicemodels/{id}`
This route will return the model information in `JSON` format. The `{id}` path variable should be replaced with a valid id. If the model exsits, the call will return the model name and id, otherwise the result will be a `404 Error`

The successful result (Code: 200):
```
{
    "id": "ee230a7b-3615-11e9-88d9-2288fa4453c9",
    "name": "Model test"
}
```

If model id not exists (Code: 400):
```
{
    "message": "Model id not found!",
    "details": "id is invalid"
}
```
## Test
#### cURL on existings (uploaded) api
This test use `cURL` on an existing api (test acount). Note that this url may be unreachable after a while. 

1. `POST` on `/api/devices`:
```
curl -X POST \
  https://6e1qp5h76k.execute-api.eu-west-3.amazonaws.com/v2/api/devices \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: b91961fc-1508-4889-94a5-e3236200c6a2' \
  -H 'cache-control: no-cache' \
  -H 'x-api-key: 3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr' \
  -d '{
	"name": "Test name2",
	"serial": "A205ad056500",
	"deviceModel":"ee230a7b-3615-11e9-88d9-2288fa4453c9",
	"note": "Test note"
}'
```

2. `POST` on `/api/devicemodels`:
```
curl -X POST \
  https://6e1qp5h76k.execute-api.eu-west-3.amazonaws.com/v2/api/devicemodels \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 819a9d4b-564c-474a-99af-f39fcecad90c' \
  -H 'cache-control: no-cache' \
  -H 'x-api-key: 3d83tuCd9f4X4yzTeOGMD8TNU6AM3xMH9vWVTcSr' \
  -d '{
	"name": "Model Thre2e05"
}'
```

3. `GET` on `/api/devicemodels/{id}`:
```
curl -X GET \
  https://6e1qp5h76k.execute-api.eu-west-3.amazonaws.com/v2/api/devices/249cae6a-3619-11e9-872e-30ec11d02de65 \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: c9dba994-c2fd-47a0-8634-1e7dd5be4e9d' \
  -H 'cache-control: no-cache'
```

4. `GET` on `/api/devices/{id}`:
```
curl -X GET \
  https://6e1qp5h76k.execute-api.eu-west-3.amazonaws.com/v2/api/devicemodels/ee230a7b-3615-11e9-88d9-2288fa4453c9 \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 282e1e43-eeaa-4dd9-be6f-2e148357585e' \
  -H 'cache-control: no-cache'
```

#### Unit Tests
In `terminal` (or `cmd`) change directory to `tests` by using `cd tests`. Now execute unit tests using `go test` command. There are two files in the `tests` directory containing unit tests on get and add device and models. You can change mocks to examine different behaviours on api. This the results of a simple run:
```
PASS
ok  	golang-aws-challenge/tests	0.046s
```
## TODO
This challenge could be expanded to ask and do more with devices and models

- Save extra fileds for `Device` and `Model` from `JSON` file into `dynamodb` as well
- Add search by other fields
- Create new model along with a new device

## Author & Maintainer
Mehrdad Esmaeilpour:
- [LinkedIn](https://www.linkedin.com/in/mehrdadep/)
- [Twitter](https://twitter.com/mehrdadep)
- [Stackoverflow](https://stackoverflow.com/users/8844510/mehrdadep)