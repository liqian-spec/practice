package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/liqian-spec/practice/app/http/controllers/api/v1"
	"github.com/liqian-spec/practice/app/models/user"
	"github.com/liqian-spec/practice/app/requests"
	"net/http"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	request := requests.SignupEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}
