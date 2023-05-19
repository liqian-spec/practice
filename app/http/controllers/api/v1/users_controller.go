package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/app/models/user"
	"github.com/practice/app/requests"
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
