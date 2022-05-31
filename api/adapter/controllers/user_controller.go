package controllers

import (
	"api/adapter/database"
	"api/domain"
	"api/usecase"
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

type UserController struct {
	Interactor usecase.UserInteractor
}

// NOTE: usecaseには&がなくdatabaseには&があるのはなぜ??
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Login(c echo.Context) (err error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEYFILE_JSON"))
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
		c.JSON(500, NewError(err))
    return
  }

	client, err := app.Auth(context.Background())
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	auth := c.Request().Header.Get("Authorization")
	idToken := strings.Replace(auth, "Bearer ", "", 1)
	_, err = client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(400, NewError(err))
		return
	}

	// NOTE: JSONデータの場合、下のようにしないとBodyのデータを取得できない FormValueは使えない
	userParam := new(domain.User)
	c.Bind(userParam)

	user, err := controller.Interactor.UserByEmail(userParam.Email)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	uid := c.FormValue("uid")
	if err = user.CompareUid(uid); err != nil {
		c.JSON(401, err.Error())
		return
	}

	payload := jwt.StandardClaims{
		Subject: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	_, err = jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	// NOTE: Cookieへ書き込み
	cookie := new(http.Cookie)
	cookie.Name = "uid"
	cookie.Value = string(user.Uid)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	c.JSON(200, user)
	return
}

func (controller *UserController) Logout(c echo.Context) (err error) {
	cookie := new(http.Cookie)
	cookie.Name = "uid"
	cookie.Value = ""
	cookie.Expires = time.Now().Add(-time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	c.JSON(200, "logout")
	return
}

func (controller *UserController) Index(c echo.Context) (err error) {
	users, err := controller.Interactor.Users()
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, users)
	return
}

func (controller *UserController) Show(c echo.Context) (err error) {
	cookie, err := c.Cookie("uid")
	if err != nil {
		c.JSON(401, NewError((err)))
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	if err = user.CompareUid(cookie.Value); err != nil {
		c.JSON(401, err.Error())
		return
	}

	c.JSON(200, user)
	return
}

func (controller *UserController) Create(c echo.Context) (err error) {
	// NOTE: ユーザの構造体を作成
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_KEYFILE_JSON"))
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
		c.JSON(500, NewError(err))
    return
  }

	client, err := app.Auth(context.Background())
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	auth := c.Request().Header.Get("Authorization")
	idToken := strings.Replace(auth, "Bearer ", "", 1)
	// NOTE: Farebaseへトークンの署名の検証を行っている
	// https://firebase.google.com/docs/admin/setup?hl=ja
	// ここではデコードしたトークンは使わないため、_としている
	_, err = client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		c.JSON(400, NewError(err))
		return
	}

	u := domain.User{}
	c.Bind(&u)
	u.SetUid(c.FormValue("uid"))

	if err = c.Validate(&u); err != nil {
		c.JSON(400, err.Error())
		return
	}

	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	cookie := new(http.Cookie)
	cookie.Name = "uid"
	cookie.Value = string(user.Uid)
	cookie.Expires = time.Now().Add(24 * time.Hour)
	// NOTE: https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Set-Cookie
	// JavaScript が Document.cookie プロパティなどを介してこのクッキーにアクセスすることを禁止します。
	cookie.HttpOnly = true
	// クッキーは、リクエストが SSL と HTTPS プロトコルを使用して行われた場合にのみサーバーに送信されます。
	// cookie.Secure = true
	cookie.Path = "/"

	c.SetCookie(cookie)
	c.JSON(201, user)
	return
}

func (controller *UserController) Save(c echo.Context) (err error) {
	u := domain.User{}
	c.Bind(&u)
	if err = c.Validate(u); err != nil {
		c.JSON(400, err.Error())
		return
	}
	user, err := controller.Interactor.Update(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(201, user)
	return
}

func (controller *UserController) Delete(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	err = controller.Interactor.DeleteById(user)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
