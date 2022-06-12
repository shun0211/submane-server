package controllers

import (
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func verifyCookie(c echo.Context) (userId int, err error) {
	cookie, err := c.Cookie("userId")
	if err != nil {
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil || !token.Valid {
		return
	}

	payload := token.Claims.(*jwt.StandardClaims)
	userId, _ = strconv.Atoi(payload.Subject)
	return
}
