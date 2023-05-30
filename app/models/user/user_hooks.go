package user

import (
	"github.com/liqian-spec/practice/pkg/hash"
	"gorm.io/gorm"
)

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}

	return
}
