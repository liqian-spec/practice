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

func (ctrl *CategoriesController) Update(c *gin.Context) {

	categoryModel := category.Get(c.Param("id"))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel.Name = request.Name
	categoryModel.Description = request.Description
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Data(c, categoryModel)
	} else {
		response.Abort500(c)
	}
}
