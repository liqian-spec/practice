package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/app/models/category"
	"github.com/liqian-spec/practice/app/requests"
	"github.com/liqian-spec/practice/pkg/response"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	fmt.Println(request.Name, 2222)
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
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
