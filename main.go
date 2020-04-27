package main

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/YouDad/localhost/modules"
	_ "github.com/YouDad/localhost/routers"
)

func init() {
	dbAlias := beego.AppConfig.String("db_alias")
	dbName := beego.AppConfig.String("db_name")
	dbUser := beego.AppConfig.String("db_user")
	dbPwd := beego.AppConfig.String("db_pwd")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbCharset := beego.AppConfig.String("db_charset")

	orm.RegisterDataBase(dbAlias, "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset="+dbCharset, 30)

	//如果是开发模式，则显示命令信息
	isDev := (beego.AppConfig.String("runmode") == "dev")
	//自动建表
	// orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ltime)
	log.SetPrefix("[INFO]: ")

	beego.Run()
}
