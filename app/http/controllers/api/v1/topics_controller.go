package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/app/models/topic"
	"github.com/practice/app/policies"
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

func (ctrl *TopicsController) Update(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Title
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.Data(c, topicModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试～")
	}
}

func (ctrl *TopicsController) Delete(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试～")
}

func (ctrl *TopicsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := topic.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
