// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"firstapi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/person",
			beego.NSInclude(
				&controllers.PersonController{},
			),
		),
	)
	beego.AddNamespace(ns)
	// 默认get方法 :id 接收参数相当于PathVariable
	beego.Router("/person/:id", &controllers.PersonController{})
	// 自定义get方法
	beego.Router("/person/calc", &controllers.PersonController{}, "get:Calc")

	beego.Router("/zbxinfo", &controllers.ZabbixController{})
}
