package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"models/vo"
	"service"
	"strconv"
)

type UserAccountController struct {
	beego.Controller
}

func (c *UserAccountController) CreateUserAccount() {
	body := c.Ctx.Input.RequestBody
	fmt.Println("body: ",string(body))
	var user vo.User
	if err:=json.Unmarshal(body,&user);err!=nil{
		fmt.Println("err: ",err)
		c.Data["json"]= `"err":""`+err.Error()+`,"status":"failed"}`
		c.ServeJSON()
		return
	}
	var userAccSrv *service.UserAccountService = &service.UserAccountService{}
	userAccSrv.UserName = user.Name
	userAccSrv.Pwd = user.Pwd
	id,err:=userAccSrv.CreateUserAccount()
	if err != nil {
		c.Data["error:"]=err
		return
	}
	c.Data["status"] = "success"
	c.Data["id"] = id
	//  `"id":`+ strconv.Itoa(id) +`,"status":"success"}`
	c.Data["json"]= `"id":`+ strconv.Itoa(id) +`,"status":"success"}`
	c.ServeJSON()
}