package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"models"
	_ "routers"
)

func main() {

	//orm.RegisterModel(new(models.UserInfo),new(models.UserAccount),new(models.Bill))
	orm.RunSyncdb("default",false,true)
	o := orm.NewOrm()
	o.Using("default")
	user := new(models.UserInfo)
	user.Age = 30
	user.Name = "slene"

	fmt.Println(user)
	beego.Run()
}

