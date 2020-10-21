package controllers

import (
	"encoding/json"
	"firstapi/dao"
	"firstapi/models"
	"github.com/astaxie/beego"
)

type ZabbixController struct {
	beego.Controller
}

type ZabbixGetRes struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    *models.ZbxServers `json:"data"`
}

type ZabbixPostRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (z *ZabbixController) Get() {
	orm := dao.GetConn()
	res := ZabbixGetRes{}
	servers := models.ZbxServers{}
	count, err := orm.Raw("select * from zbx_server").QueryRows(&servers)
	if err == nil {
		if count >= 0 {
			res.Success = true
			res.Data = &servers
		} else {
			res.Success = false
			res.Message = "暂无数据"
		}
	} else {
		res.Success = false
		res.Message = err.Error()
	}
	z.Data["json"] = res
	z.Ctx.Output.JSON(z.Data["json"], true, true)

}

func (z *ZabbixController) Post() {
	orm := dao.GetConn()
	newServerInfo := models.ZbxServer{}
	res := ZabbixPostRes{}
	if err := json.Unmarshal(z.Ctx.Input.RequestBody, &newServerInfo); err == nil {
		_, err := orm.Insert(&newServerInfo)
		if err == nil {
			res.Message = "success"
			res.Success = true
		} else {
			res.Message = err.Error()
			res.Success = false
		}
	} else {
		res.Message = err.Error()
		res.Success = false
	}
	z.Data["json"] = res
	z.Ctx.Output.JSON(z.Data["json"], true, true)
}
