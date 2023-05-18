package factories

import (
	"github.com/bxcodec/faker/v3"
	"github.com/practice/app/models/user"
	"github.com/practice/pkg/helpers"
)

func Makeusers(times int) []user.User {

	var objs []user.User

	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Name:     faker.Username(),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}
		objs = append(objs, model)
	}

	return objs
}
