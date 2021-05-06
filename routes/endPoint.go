package routes

import (
	"be-sharing-vision/service"

	"github.com/labstack/echo"
)

func EndPoint() {
	e := echo.New()

	e.POST("user", service.CreateNewUser)
	e.GET("user/:limit/:offset", service.ReadAllUsers)
	e.GET("user/:id", service.ReadUserById)
	e.PUT("user/:id", service.UpdateUser)
	e.DELETE("user/:id", service.DeleteUser)

	e.Logger.Fatal((e.Start(":3000")))
}
