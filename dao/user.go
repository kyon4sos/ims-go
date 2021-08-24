package dao

import (
	. "img-server/model"
)

func GetUserByUsername(username string) *User {
	var user User
	find := GetDb().Limit(1).Where(&User{
		Username: username,
	}).Find(&user)
	if find.RowsAffected>0 {
		return &user
	}
	return nil
}
