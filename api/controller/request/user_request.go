package request

import (
	"github.com/Ghun2/fast-gorestapi/model"
	"github.com/labstack/echo/v4"
)

type UserRegisterRequest struct {
	UserName 	string `json:"user_name"`
	AuthID	 	string `json:"auth_id"`
	Email    	string `json:"email"`
	Birth    	string `json:"birth"`
	Sex    	 	string `json:"sex"`
	Phone    	string `json:"phone"`
}

func (r *UserRegisterRequest) Bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.UserName = r.UserName
	u.AuthID = r.AuthID
	u.Email = r.Email
	u.Birth = r.Birth
	u.Sex = r.Sex
	u.Phone = r.Phone
	return nil
}

type UserLoginRequest struct {
	AuthID	string `json:"auth_id"`
}

func (r *UserLoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type UserUpdateRequest struct {
	UserName 	string `json:"user_name"`
	Email    	string `json:"email"`
	Birth    	string `json:"birth"`
	Sex    	 	string `json:"sex"`
	Phone    	string `json:"phone"`
}

func NewUserUpdateRequest() *UserUpdateRequest {
	return new(UserUpdateRequest)
}

func (r *UserUpdateRequest) Populate(u *model.User) {
	r.UserName = u.UserName
	r.Email = u.Email
	r.Birth = u.Birth
	r.Sex = u.Sex
	r.Phone = u.Phone
}

func (r *UserUpdateRequest) Bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.UserName = r.UserName
	u.Email = r.Email
	u.Birth = r.Birth
	u.Sex = r.Sex
	u.Phone = r.Phone
	return nil
}
