package policies

import (
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/app/models/topic"
	"github.com/liqian-spec/practice/pkg/auth"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
