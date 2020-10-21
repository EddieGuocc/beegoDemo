package models

type ZbxServer struct {
	Zid       int64  `orm:"pk"`
	ZbxServer string `orm:"column(zbx_server)"`
	ZbxUser   string `orm:"column(zbx_user)"`
	ZbxPwd    string `orm:"column(zbx_pswd)"`
}

type ZbxServers []ZbxServer
