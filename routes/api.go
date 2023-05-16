package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/app/http/controllers/api/v1/auth"
	"github.com/practice/app/http/middlewares"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//测试一个v1的路由组，我们所有的v1版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		//// 注册一个路由
		//v1.GET("/", func(c *gin.Context) {
		//	// 以JSON格式响应
		//	c.JSON(http.StatusOK, gin.H{
		//		"Hello": "World",
		//	})
		//})

		v1.Use(middlewares.LimitIP("200-H"))
		authGroup := v1.Group("/auth")
		{
			// 注册
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), suc.IsEmailExist)
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("50-H"), vcc.SendUsingEmail)

			// 登录
			lgc := new(auth.LoginController)
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)
		}
	}
}
