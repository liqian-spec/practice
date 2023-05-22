package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/app/models/link"
	"github.com/practice/pkg/response"
)

type LinksController struct {
	BaseApiController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.All()
	response.Data(c, links)
}
