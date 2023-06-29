package server

import (
	"lambda-test/internal/api/controller"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"
	"github.com/labstack/echo/v4"
)

func NewServer(hc controller.IHelloController) *echoadapter.EchoLambdaV2 {
	
	e := echo.New()
	e.GET("/", hc.GetHelloWorld)
	
	echoLambda := echoadapter.NewV2(e)

	return echoLambda
}