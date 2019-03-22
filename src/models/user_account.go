package models

import "github.com/astaxie/beego/orm"



type UserAccount struct {
	Id          int    `orm:(id)`
	UserName    string `orm:(username)`
	Pwd         string `orm:(pwd)`
	UserInfoId  int    `orm:(user_info_id)`
	Status      int    `orm:(status)`
	Description string `orm:(string)`
	Role        int    `orm:(role)`
}

func (m *UserAccount) CreateUser() (int,error) {
	var Orm orm.Ormer = orm.NewOrm()
	id,err:=Orm.Insert(m)
	return int(id), err
}
