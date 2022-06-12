package controllers

import (
	"api/adapter/database"
	"api/domain"
	"api/usecase"
	"api/utils"
	"context"
	"os"
	"strconv"
	"strings"

	firebase "firebase.google.com/go"
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
		c.JSON(404, err.Error())
		return
	}

	uid := c.FormValue("uid")
	if err = user.CompareUid(uid); err != nil {
		c.JSON(401, err.Error())
		return
	}

	jwt, err := generateJWT(strconv.Itoa(int(user.ID)))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	setCookie(c, jwt)

	c.JSON(200, user)
	return
}

func (controller *UserController) Logout(c echo.Context) (err error) {
	setCookie(c, "")
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
	cookie, err := c.Cookie("userId")
	if err != nil {
		c.JSON(401, NewError(err))
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if cookie.Value == c.Param("id") {
		c.JSON(401, err.Error())
		return
	}

	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(404, NewError(err))
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
		c.JSON(400, "不正なIDトークンです")
		return
	}

	u := domain.User{}
	c.Bind(&u)
	u.SetUid(c.FormValue("uid"))

	if err = c.Validate(&u); err != nil {
		messages := utils.GetErrorMessages(err)
		c.JSON(400, messages)
		return
	}

	user, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}

	jwt, err := generateJWT(strconv.Itoa(int(user.ID)))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}
	setCookie(c, jwt)

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
	c.JSON(200, user)
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


func (controller *UserController) ShowCurrentUser(c echo.Context) (err error) {
	userId, err := verifyCookie(c)
	if err != nil {
		c.JSON(401, err)
		return
	}

	user, err := controller.Interactor.UserById(userId)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
	return
}
