package controllers

import (
	"encoding/json"
	"encoding/xml"
	"github.com/astaxie/beego"
	"strconv"
)

type PersonController struct {
	beego.Controller
}

type Address struct {
	Details string `json:"details" xml:"details"`
	Code    int    `json:"code" xml:"code"`
}

type Person struct {
	Name string  `json:"name" xml:"name"`
	Age  int     `json:"age" xml:"age"`
	Addr Address `json:"addr" xml:"addr"`
}

// 接受指针变量 Override beego.Controller的Get方法
func (p *PersonController) Get() {
	// 接收参数 :不可缺省
	id := p.Ctx.Input.Param(":id")
	// 直接赋值作为返回参数
	p.Ctx.WriteString("default get method in PersonController, param id " + id)
}

// 自定义方法
func (p *PersonController) Calc() {
	num1, err := p.GetInt("num1")
	num2, err1 := p.GetInt("num2")
	if err != nil || err1 != nil {
		p.Ctx.WriteString("params wrong")
		return
	}
	res := strconv.Itoa(num1 + num2)
	p.Ctx.WriteString("calc method in PersonController, result is " + res)
}

// Override beego.Controller的Post方法
//request Body
//{
//    "name":"abc",
//    "age":20,
//    "addr":{
//        "details":"XXXXXXXX",
//        "code":12345
//    }
//}
func (p *PersonController) Post() {
	var err error
	newPerson := Person{}
	if err = json.Unmarshal(p.Ctx.Input.RequestBody, &newPerson); err == nil {
		p.Ctx.WriteString("post method (json type) in PersonController, name is " + newPerson.Name +
			" age is " + strconv.Itoa(newPerson.Age) +
			" address is " + newPerson.Addr.Details +
			" address code is " + strconv.Itoa(newPerson.Addr.Code))
	} else {
		p.Ctx.WriteString("request body (json type) type wrong details[ " + err.Error() + " ]")
	}
}

// Override beego.Controller的Put方法
func (p *PersonController) Put() {
	var err error
	newPerson := Person{}
	if err = xml.Unmarshal(p.Ctx.Input.RequestBody, &newPerson); err == nil {
		p.Ctx.WriteString("put method (xml type) in PersonController, name is " + newPerson.Name +
			" age is " + strconv.Itoa(newPerson.Age) +
			" address is " + newPerson.Addr.Details +
			" address code is " + strconv.Itoa(newPerson.Addr.Code))
	} else {
		p.Ctx.WriteString("request body (xml type) type wrong details[ " + err.Error() + " ]")
	}
}

// 伪静态路由跳转
