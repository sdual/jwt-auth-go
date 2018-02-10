package main

import (
	"./handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	runRoute(e)
}

func runRoute(e *echo.Echo) {

}
