package controller

import (
	"github.com/Ghun2/fast-gorestapi/api/controller/request"
	"github.com/Ghun2/fast-gorestapi/api/controller/response"
	"github.com/Ghun2/fast-gorestapi/model"
	"github.com/Ghun2/fast-gorestapi/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *Controller) SignUp(c echo.Context) error {
	var u model.User
	req := &request.UserRegisterRequest{}
	if err := req.Bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := ctrl.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	c.Response().Header().Set("Authorization", utils.GenerateJWT(u.UserID))
	return c.JSON(http.StatusCreated, response.NewUserResponse(&u))
}

func (ctrl *Controller) Login(c echo.Context) error {
	req := &request.UserLoginRequest{}
	if err := req.Bind(c); err != nil {
		return err
	}
	u, err := ctrl.userStore.GetByAuthID(req.AuthID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	c.Response().Header().Set("Authorization", utils.GenerateJWT(u.UserID))
	return c.JSON(http.StatusOK, response.NewUserResponse(u))
}

func (ctrl *Controller) CurrentUser(c echo.Context) error {
	u, err := ctrl.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, response.NewUserResponse(u))
}

func (ctrl *Controller) UpdateUser(c echo.Context) error {
	u, err := ctrl.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := request.NewUserUpdateRequest()
	req.Populate(u)
	if err := req.Bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := ctrl.userStore.Update(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, response.NewUserResponse(u))
}

func (ctrl *Controller) DeleteUser(c echo.Context) error {
	u, err := ctrl.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	if err := ctrl.userStore.Delete(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}