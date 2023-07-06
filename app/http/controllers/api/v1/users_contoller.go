package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/app/models/user"
	"github.com/liqian-spec/practice/app/requests"
	"github.com/liqian-spec/practice/pkg/auth"
	"github.com/liqian-spec/practice/pkg/response"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context) {

	request := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateProfile); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Name = request.Name
	currentUser.City = request.City
	currentUser.Introduction = request.Introduction
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试～")
	}
}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {

	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试～")
	}
}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {

	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {

	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		response.Unauthorized(c, "原密码不正确")
	} else {
		currentUser.Password = request.NewPassword
		currentUser.Save()

		response.Success(c)
	}
}
