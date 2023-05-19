package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/app/models/category"
	"github.com/practice/app/requests"
	"github.com/practice/pkg/response"
)

type CategoriesController struct {
	BaseApiController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name:        request.Name,
		Description: request.Description,
	}
	categoryModel.Create()
	if categoryModel.ID > 0 {
		response.Created(c, categoryModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试～")
	}
}
