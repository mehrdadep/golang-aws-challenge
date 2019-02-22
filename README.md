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
- install nodejs
- install go
- install aws-cli
- install serverless framework using `npm install -g serverless`

###### Dependencies
Place this folder into the `src` folder of `${GOPATH}`, This code needs the following packages in go vendors to work correctly:
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
<br>
2. `POST` methods are protected by api-key and you should generate an api key and use it as the value of the `x-api-key` key in request header. To do so, edit `serverless.yml` file and set `stage` and `apiKeys` for the first time. An api key will be generated after the first deploy and `stage` will be in the api's final url as follows:
<pre>
https://e7rjun495i.execute-api.eu-west-3.amazonaws.com/<b>stage</b>/api/devices
</pre>
<br>
3. All neccessary commands are called using `make` command (Linux64-only). Alternatively you can run the following commands one by one:<br>
```
build clean deploy
dep ensure -v
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addModel addModel/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/addDevice addDevice/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getDevice getDevice/main.go
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/getModel getModel/main.go
rm -rf ./bin ./vendor Gopkg.lock
clean build
sls deploy --verbose
```
<br>
You can see the output of a successful deploy on amazon servers.
![Terminal Output](img/terminal.png?raw=true "Terminal Output")

## Routes
There are four final routes in this api as follows:
1. `POST` on `/api/devices` will add a new device to database. Payload should be plain text (`JSON`). The api also check for duplicate serial numbers and prevent you from storing a device twice. `deviceModel` should be a valid `ID` of a model. you can get a valid `ID` of a model from part 2 of this section. The result of this call is an new id for new device. Remember this is a protected uri, so you have to set api key in `x-api-key` key of the header. `JSON` payload should be formated as follows:<br>
```
{
	"name": "Test name",
	"serial": "A205ad0500",
	"deviceModel":"ee230a7b-3615-11e9-88d9-2288fa44503c9",
	"note": "Test note"
}
```
<br>
2. `POST` on `/api/devicemodels` will add a new model to database. Payload should be plain text (`JSON`). The api also check for duplicate names and prevent you from storing a model twice. In any successful case a call will return the model `ID`. `JSON` payload should be formated as follows:<br>
```
{
	"name": "Test name"
}
```
<br>
3. `GET` on `/api/devices/{id}` will return the device information in `JSON` format. The `{id}` path variable should be replaced with a valid id. This call will return the device name and other details if exists and not found error, otherwise.
<br>
4. `GET` on `/api/devicemodels/{id}` will return the model information in `JSON` format. The `{id}` path variable should be replaced with a valid id. This call will return the model name and id if exists and not found error, otherwise.
## Test

## Maintainer
Mehrdad Esmaeilpour:
- [LinkedIn](https://www.linkedin.com/in/mehrdadep/)
- [Twitter](https://twitter.com/mehrdadep)
- [Stackoverflow](https://stackoverflow.com/users/8844510/mehrdadep)