package auth

import (
	v1 "github.com/practice/app/http/controllers/api/v1"
	"github.com/practice/app/requests"
	"github.com/practice/pkg/auth"
	"github.com/practice/pkg/jwt"
	"github.com/practice/pkg/response"

	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseApiController
}

// LoginByPhone 手机登录
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
