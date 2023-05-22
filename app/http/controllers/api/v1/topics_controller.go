package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/app/models/topic"
	"github.com/practice/app/requests"
	"github.com/practice/pkg/auth"
	"github.com/practice/pkg/response"
)

type TopicsController struct {
	BaseApiController
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试～")
	}
}
