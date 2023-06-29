package controller

import (
	"encoding/json"
	"lambda-test/internal/api/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IHelloController interface {
	GetHelloWorld(c echo.Context) error
}

type helloController struct {
}

func NewHelloController() IHelloController {
	return &helloController{}
}

func (h *helloController) GetHelloWorld(c echo.Context)error{

	msg := model.Message{
		Message: "hello world",
	}
	res,err :=json.Marshal(msg)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, string(res))
	
}



