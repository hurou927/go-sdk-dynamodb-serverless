package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
// type Response events.APIGatewayProxyResponse
// type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
var globalVar int = 0;

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("%+v\n", os.Getenv("LAMBDA_TEST"));
	var buf bytes.Buffer

	globalVar = globalVar + 1

	body, err := json.Marshal(map[string]interface{}{
		"value": globalVar,
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return events.APIGatewayProxyResponse {StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := events.APIGatewayProxyResponse {
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	globalVar = 100
	lambda.Start(Handler)
}
