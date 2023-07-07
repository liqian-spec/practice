package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liqian-spec/practice/pkg/response"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {

	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认格式是否正确，上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		fmt.Println(err.Error())
		return false
	}

	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}

func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).Validate()
}
