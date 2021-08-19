package dao

import (
	"img-server/model"
	"log"
)

func GetAllMenus() []*model.Menu {
	var menus []*model.Menu
	res := Db().Find(&menus)
	if res.Error != nil {
		log.Println("menus err",res.Error.Error())
		return nil
	}
	return menus
}
