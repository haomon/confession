package router

import (
	"testproject/confession/controller"

	"github.com/labstack/echo"
)

func Router(e *echo.Echo) {
	f := e.Group("/confession")
	f.GET("/9527/version", controller.Version)
}
