package driver

import (
	"api/adapter/controllers"

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

	// Login
	e.POST("/login", func(c echo.Context) error { return userController.Login(c) })

	// User CRUD
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/user/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.Delete(c) })

	// Subscription CRUD
	e.GET("/subscriptions", func(c echo.Context) error { return subscriptionController.Index(c) })
	e.POST("/subscriptions", func(c echo.Context) error { return subscriptionController.Create(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
