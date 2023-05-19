package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/app/models/user"
	"github.com/practice/pkg/auth"
	"github.com/practice/pkg/response"
)

type UsersController struct {
	BaseApiController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context) {
	data := user.All()
	response.Data(c, data)
}
