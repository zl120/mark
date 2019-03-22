package service

import (
	"github.com/astaxie/beego/orm"
	"models"
	//"models/vo"
)

var Orm orm.Ormer = orm.NewOrm()

type UserAccountService struct {
	Id          int
	UserName    string
	Pwd         string
	UserInfoId  int
	Status      int
	Description string
	Role        int
}

func (s *UserAccountService) CreateUserAccount() (int,error) {
	var userAccMod models.UserAccount
	userAccMod.UserName = s.UserName
	userAccMod.Pwd = s.Pwd
	return userAccMod.CreateUser()
}