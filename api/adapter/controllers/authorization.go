package controllers

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	firebase "firebase.google.com/go"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
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

func verifyIDToken(c echo.Context) (status int, err error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEYFILE_JSON"))
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
		status = 500
    return
  }

	client, _ := app.Auth(context.Background())
	auth := c.Request().Header.Get("Authorization")
	idToken := strings.Replace(auth, "Bearer ", "", 1)
	// NOTE: Farebaseへトークンの署名の検証を行っている
	// https://firebase.google.com/docs/admin/setup?hl=ja
	// ここではデコードしたトークンは使わないため、_としている
	_, err = client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		status = 400
		return
	}
	return
}
