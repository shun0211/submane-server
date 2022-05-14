package controllers

import (
	"api/adapter/database"
	"api/domain"
	"api/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
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
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte("secret"))
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	// NOTE: Cookieへ書き込み
	cookie := new(http.Cookie)
	cookie.Name = user.Name
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	c.SetCookie(cookie)

	c.JSON(200, user)
	return
}

func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := controller.Interactor.UserById(id)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, user)
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

func (controller *UserController) Create(c echo.Context) (err error) {
	// NOTE: ユーザの構造体を作成
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
