package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}

type File struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Size     int    `json:"size"`
	Url      string `json:"url"`
	CreateTs string `json:"create_ts"`
}

type FidUid struct {
	Id  int
	Fid int
	Uid int
}

func init() {
	orm.RegisterModel(new(User), new(File), new(FidUid))
}
