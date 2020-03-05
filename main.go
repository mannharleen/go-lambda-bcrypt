package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"golang.org/x/crypto/bcrypt"
)

type RequestBody struct {
	Password string `json:"password"`
	Hash     string `json:"hash"`
}

type ResponseBody struct {
	Message string `json:"message"`
}

func HandleLambdaEvent(event events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var rqb RequestBody
	var rsb ResponseBody
	var toReturnBody string

	err := json.Unmarshal([]byte(event.Body), &rqb)
	if err != nil {
		return ReturnError(err, 400)
	}
	if rqb.Password == "" {
		return ReturnError(errors.New("Request body must contain a valid json with a nonempty 'password' attribute"), 400)
	}

	if event.Path == "/bcrypt/hash" {
		r, err := bcrypt.GenerateFromPassword([]byte(rqb.Password), 10)
		if err != nil {
			return ReturnError(err, 500)
		}
		rsb.Message = string(r)
		body, err := json.Marshal(rsb)
		if err != nil {
			return ReturnError(err, 500)
		}
		toReturnBody = string(body)
	} else if event.Path == "/bcrypt/verify" {
		if rqb.Hash == "" {
			return ReturnError(errors.New("Request body must contain a valid json with a nonempty 'hash' attribute"), 400)
		}
		err := bcrypt.CompareHashAndPassword([]byte(rqb.Hash), []byte(rqb.Password))
		if err != nil {
			toReturnBody = fmt.Sprintf(`{"message" : %q}`, "invalid")
		} else {
			toReturnBody = fmt.Sprintf(`{"message" : %q}`, "valid")
		}
	} else {
		return ReturnError(errors.New("Path does not exist. Use one of: /bcrypt/hash or /bcrypt/verify"), 400)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode:      200,
		Headers:         map[string]string{"content-type": "application/json"},
		Body:            toReturnBody,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}

func ReturnError(err error, code int) (*events.APIGatewayProxyResponse, error) {
	fmt.Println("error", err)
	return &events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    map[string]string{"content-type": "application/json"},
		Body:       fmt.Sprintf(`{"message" : %q}`, err.Error()),
	}, nil
}
