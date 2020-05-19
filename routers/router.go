package routers

import (
	"file-upload/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/create_table", &controllers.MainController{}, "get:CreateTable")
	beego.Router("/login", &controllers.UserController{}, "get:Login")
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/register", &controllers.UserController{}, "post:Register")
	beego.Router("/upload", &controllers.FileController{}, "post:Upload")
	beego.Router("/download", &controllers.FileController{}, "get:Download")
	beego.Router("/files", &controllers.FileController{}, "post:Files")
	beego.Router("/file/delete", &controllers.FileController{}, "get:DeleteFile")
}
