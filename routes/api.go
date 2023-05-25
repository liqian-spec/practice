package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/app/http/controllers/api/v1/auth"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}
}
