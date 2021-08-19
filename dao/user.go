package dao

import (
	"img-server/model"
	. "img-server/model"
)

func GetUserByUsername(username string) *model.User {
	var user *User

	Db().Where(&User{
		Username: username,
	}).First(user)
	// if res.Error != nil {
	// 	return nil
	// }
	return user
}
