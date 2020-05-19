package util

import (
	"file-upload/models"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

type Result struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func GetInstance() Result {
	return Result{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}
}

func CreateTable() {
	//注册数据驱动
	orm.RegisterDriver("mysql", orm.DRMySQL) // mysql / sqlite3 / postgres 这三种是beego默认已经注册过的，所以可以无需设置
	//注册数据库 ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/file_upload?charset=utf8")
	//注册模型
	orm.RegisterModel(new(models.User), new(models.File), new(models.FidUid))
	//自动创建表 参数二为是否开启创建表   参数三是否更新表
	orm.RunSyncdb("default", false, true)
}
