package controllers

import (
	"net/http"
	"strconv"
	"time"

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

// NOTE: Cookieの処理移す
func setCookie(c echo.Context, jwt string) {
	cookie := new(http.Cookie)
	cookie.Name = "userId"
	cookie.Value = jwt
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// NOTE: https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Set-Cookie
	// JavaScript が Document.cookie プロパティなどを介してこのクッキーにアクセスすることを禁止します。
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteNoneMode
	// クッキーは、リクエストが SSL と HTTPS プロトコルを使用して行われた場合にのみサーバーに送信されます。
	cookie.Secure = true

	c.SetCookie(cookie)
}

func generateJWT(value string) (jwtString string, err error) {
	payload := jwt.StandardClaims{
		Subject: value,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtString, err = jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	return
}
