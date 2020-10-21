package dao

import (
	"firstapi/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	port, _ := beego.AppConfig.Int("mysql.Port")
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		beego.AppConfig.String("mysql.Username"), beego.AppConfig.String("mysql.Password"),
		beego.AppConfig.String("mysql.Host"), port,
		beego.AppConfig.String("mysql.Schema"))
	fmt.Println(url)
	MaxIdleConn, _ := beego.AppConfig.Int("mysql.MaxIdleConnCnt")
	MaxOpenConn, _ := beego.AppConfig.Int("mysql.MaxOpenConnCnt")
	orm.RegisterDataBase("default", "mysql", url, MaxIdleConn, MaxOpenConn)
	orm.RegisterModel(new(models.ZbxServer))
}

func GetConn() orm.Ormer {
	return orm.NewOrm()
}
