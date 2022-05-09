package infrastructure

import (
	"api/interfaces/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	// Echo instance
	e := echo.New()

	userController := controllers.NewUserController(NewSqlHandler())
	subscriptionController := controllers.NewSubscriptionController(NewSqlHandler())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// User CRUD
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/user/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/create", func(c echo.Context) error { return userController.Create(c) })
	e.PUT("/user/:id", func(c echo.Context) error { return userController.Save(c) })
	e.DELETE("/user/:id", func(c echo.Context) error { return userController.Delete(c) })

	// Subscription CRUD
	e.GET("/subscription", func(c echo.Context) error { return subscriptionController.Index(c) })

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
