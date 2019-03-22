package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func init() {
	err:=orm.RegisterDriver("mysql", orm.DRMySQL)    //可以不加
	if err != nil {
		fmt.Println("err: ",err)
		return
	}
	orm.RegisterModel(new(UserInfo),new(UserAccount),new(Bill))
	//ORM 必须注册一个别名为default的数据库，作为默认使用。
	if err:=orm.RegisterDataBase("default", "mysql", "mysql:Huawei@123@tcp(116.62.60.254:3306)/zl_test?charset=utf8");err!=nil{
		fmt.Println("err: ",err)
		return
	}
}
