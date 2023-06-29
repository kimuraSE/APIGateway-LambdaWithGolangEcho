package main

import (
	"context"
	"lambda-test/internal/api/controller"
	"lambda-test/pkg/server"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echoadapter.EchoLambdaV2

func init() {
	helloController := controller.NewHelloController()
	echoLambda = server.NewServer(helloController)
}

func Handler(ctx context.Context ,req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}