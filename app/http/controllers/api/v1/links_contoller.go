package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/app/models/link"
	"github.com/liqian-spec/practice/pkg/response"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCache())
}
