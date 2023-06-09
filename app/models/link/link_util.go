package link

import (
	"github.com/gin-gonic/gin"
	"github.com/practice/pkg/app"
	"github.com/practice/pkg/cache"
	"github.com/practice/pkg/database"
	"github.com/practice/pkg/helpers"
	"github.com/practice/pkg/paginator"
	"time"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCache() (links []Link) {

	cacheKey := "links:all"
	expireTime := 120 * time.Minute
	cache.GetObject(cacheKey, &links)

	if helpers.Empty(links) {
		links = All()
		if helpers.Empty(links) {
			return links
		}
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
