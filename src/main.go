package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "routers"
)

func main() {

	//orm.RegisterModel(new(models.UserInfo),new(models.UserAccount),new(models.Bill))
	orm.RunSyncdb("default",false,true)
	beego.Run()
}

