package controller

import (
	"github.com/Ghun2/fast-gorestapi/router/middleware"
	"github.com/Ghun2/fast-gorestapi/utils"
	"github.com/labstack/echo/v4"
)

func (ctrl *Controller) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("", ctrl.SignUp)
	guestUsers.POST("/login", ctrl.Login)

	user := v1.Group("/user", jwtMiddleware)
	user.GET("", ctrl.CurrentUser)
	user.PUT("", ctrl.UpdateUser)
	user.DELETE("", ctrl.DeleteUser)

}
