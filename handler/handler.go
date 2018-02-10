package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router interface {
	Route(g *echo.Group)
}

type OtherRouter struct {
}

type RestrictedHogeRouter struct {
}

type RestrictedHugaRouter struct {
}
