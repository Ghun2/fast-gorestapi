package response

import (
	"github.com/Ghun2/fast-gorestapi/model"
)

type userResponse struct {
	User struct{
		UserID      uint   `json:"user_id"`
		UserName 	string `json:"user_name"`
		Email    	string `json:"email"`
		Birth    	string `json:"birth"`
		Sex    	 	string `json:"sex"`
		Phone    	string `json:"phone"`
	} `json:"user"`
}

func NewUserResponse(u *model.User) *userResponse {
	r := new(userResponse)
	r.User.UserID = u.UserID
	r.User.UserName = u.UserName
	r.User.Email = u.Email
	r.User.Birth = u.Birth
	r.User.Sex = u.Sex
	r.User.Phone = u.Phone
	return r
}
