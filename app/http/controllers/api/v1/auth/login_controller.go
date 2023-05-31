package auth

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/liqian-spec/practice/app/http/controllers/api/v1"
	"github.com/liqian-spec/practice/app/requests"
	"github.com/liqian-spec/practice/pkg/auth"
	"github.com/liqian-spec/practice/pkg/jwt"
	"github.com/liqian-spec/practice/pkg/response"
)

type LoginController struct {
	v1.BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {

	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "账号不存在")
	} else {

		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
