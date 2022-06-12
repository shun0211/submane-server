package driver

import (
	"api/adapter/controllers"
	"api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		print("before next")
		err := next(c)
		print("after next")
		if err != nil {
			c.Error(err)
		}
		return nil
	}
}

func Init() {
	// Echo instance
	e := echo.New()
	e.Validator = utils.NewValidator()

	userController := controllers.NewUserController(NewSqlHandler())
	subscriptionController := controllers.NewSubscriptionController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// NOTE: クッキーのような資格情報を含んでいてAccess-Control-Allow-Origin: *のようにワイルドカードが指定された場合、ブラウザは資格情報を破棄する
		// NOTE: 実際のオリジンを指定する必要がある
		AllowOrigins:     []string{"http://localhost:3001"},
		// NOTE: プリフライトのレスポンスで用いられ、リソースへのアクセス時に許可するメソッドを指定する
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		// NOTE: プリフライトでの場合、実際のリクエストを資格情報付きで行うことができることを示す
		// NOTE: GETリクエストの場合、リソースとともにこのヘッダーを返さないと、ブラウザがレスポンスを破棄する
		AllowCredentials: true,
	}))
	e.Use(Process)

	// Login, Logout
	e.POST("/login", func(c echo.Context) error { return userController.Login(c) })
	e.DELETE("/logout", func(c echo.Context) error { return userController.Logout(c) })

	// User CRUD
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/users/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Delete(c) })
	e.GET("/current-user", func(c echo.Context) error { return userController.ShowCurrentUser(c) })

	// Subscription CRUD
	e.GET("/subscriptions", func(c echo.Context) error { return subscriptionController.Index(c) })
	e.GET("/subscriptions/:id", func(c echo.Context) error { return subscriptionController.Show(c) })
	e.POST("/subscriptions", func(c echo.Context) error { return subscriptionController.Create(c) })
	e.PUT("/subscriptions/:id", func(c echo.Context) error { return subscriptionController.Save(c) })
	e.DELETE("/subscriptions/:id", func(c echo.Context) error { return subscriptionController.Delete(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
