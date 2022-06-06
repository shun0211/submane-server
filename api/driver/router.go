package driver

import (
	"api/adapter/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	// Echo instance
	e := echo.New()
	e.Validator = NewValidator()

	userController := controllers.NewUserController(NewSqlHandler())
	subscriptionController := controllers.NewSubscriptionController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
}))

	// Login, Logout
	e.POST("/login", func(c echo.Context) error { return userController.Login(c) })
	e.DELETE("/logout", func(c echo.Context) error { return userController.Logout(c) })

	// User CRUD
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/user/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.Delete(c) })
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
