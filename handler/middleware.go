package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	RoleId int64 `json:"role_id"`
	jwt.StandardClaims
}

func restrictedMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		roleId := claims.RoleId
		if roleId != 1 {
			return c.String(http.StatusUnauthorized, "Unauthorized.")
		}
		return next(c)
	}
}
