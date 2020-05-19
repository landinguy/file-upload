package controllers

import (
	"encoding/json"
	"file-upload/models"
	"file-upload/util"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Printf("用户登录,username#%s,password#%s\n", username, password)
	result := util.GetInstance()
	user := models.User{}
	err := util.O.QueryTable("user").Filter("Username", username).Filter("Password", password).One(&user)
	if err == nil {
		fmt.Println(user)
		var data = map[string]interface{}{"uid": user.Id, "username": user.Username, "role": user.Role, "phone": user.PhoneNumber, "email": user.Email}
		result.Data = data
		this.SetSession("uid", user.Id)
	} else {
		fmt.Println(err.Error())
		result.Code = -1
		result.Msg = "账号或密码错误"
	}
	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *UserController) Logout() {
	this.DelSession("uid")
	result := util.GetInstance()
	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *UserController) Register() {
	result := util.GetInstance()
	var user models.User
	params := this.Ctx.Input.RequestBody
	fmt.Printf("用户注册,params#%s\n", string(params))
	err := json.Unmarshal(params, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	u := models.User{}
	err = util.O.QueryTable("user").Filter("Username", user.Username).One(&u)
	if err == nil {
		result.Code = -1
		result.Msg = "账号已存在！"
	} else {
		user.Role = "NORMAL"
		_, err := util.O.Insert(&user)
		if err != nil {
			result.Code = -1
			result.Msg = "用户注册失败"
		}
	}
	this.Data["json"] = &result
	this.ServeJSON()
}
