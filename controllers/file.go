package controllers

import (
	"encoding/json"
	"file-upload/models"
	"file-upload/util"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type FileController struct {
	beego.Controller
}

var filePath = beego.AppConfig.String("filePath")
var url = beego.AppConfig.String("url")

func (this *FileController) Upload() {
	result := util.GetInstance()
	file, head, err := this.GetFile("file")
	if err == nil {
		filename := head.Filename
		size := int(head.Size)
		fmt.Printf("上传文件,filename#%s,size#%d\n", filename, size)
		err = this.SaveToFile("file", filePath+filename)
		if err != nil {
			fmt.Println(err.Error())
			result.Code = -1
			result.Msg = "文件上传失败"
		} else {
			createTs := time.Now().Format("2006-01-02 15:04:05")
			f := models.File{Name: filename, Size: int(size), CreateTs: createTs}
			id, err := util.O.Insert(&f)
			if err == nil {
				f.Url = url + strconv.FormatInt(id, 10)
				util.O.Update(&f)

				uid, _ := this.GetInt("uid")
				fmt.Println("----uid", uid)
				fidUid := models.FidUid{Fid: int(id), Uid: uid}
				util.O.Insert(&fidUid)
			}
		}
	}
	defer file.Close()
	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *FileController) Download() {
	fid, _ := this.GetInt("fid")
	fmt.Printf("下载文件,fid#%d\n", fid)
	if fid != 0 {
		file := models.File{}
		util.O.QueryTable("file").Filter("Id", fid).One(&file)
		filename := file.Name
		//第一个参数是文件的地址，第二个参数是下载显示的文件的名称
		this.Ctx.Output.Download(filePath+filename, filename)
	}
}

func (this *FileController) Files() {
	result := util.GetInstance()
	params := this.Ctx.Input.RequestBody
	fmt.Printf("查询文件,params#%s\n", string(params))
	m := map[string]interface{}{}
	err := json.Unmarshal(params, &m)
	if err == nil {
		query := util.O.QueryTable("fid_uid")
		uid := m["uid"].(string)
		user := models.User{}
		util.O.QueryTable("user").Filter("Id", uid).One(&user)
		if user.Role != "ADMIN" {
			query = query.Filter("Uid", uid)
		}

		var fidUidList []*models.FidUid
		query.All(&fidUidList)
		fidMap := map[int]int{}
		for _, it := range fidUidList {
			fidMap[it.Fid] = it.Fid
		}
		var fidList []int
		for it := range fidMap {
			fidList = append(fidList, it)
		}
		fmt.Println("fidList#", fidList)
		if len(fidList) == 0 {
			fidList = append(fidList, 0)
		}

		fileQuery := util.O.QueryTable("file").Filter("id__in", fidList)
		total, _ := fileQuery.Count()

		pageNo := int(m["pageNo"].(float64))
		pageSize := int(m["pageSize"].(float64))
		var list []*models.File
		fileQuery.Offset((pageNo - 1) * pageSize).Limit(pageSize).OrderBy("-id").All(&list)
		result.Data = map[string]interface{}{"total": total, "list": list}
	}
	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *FileController) DeleteFile() {
	fid, _ := this.GetInt("id")
	fmt.Printf("删除文件,fid#%d\n", fid)
	if fid != 0 {
		util.O.Delete(&models.File{Id: fid})
		var list []*models.FidUid
		util.O.QueryTable("fid_uid").Filter("Fid", fid).All(&list)
		for _, item := range list {
			util.O.Delete(&models.FidUid{Id: item.Id})
		}
	}
	result := util.GetInstance()
	this.Data["json"] = &result
	this.ServeJSON()
}
