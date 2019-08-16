package models

import (
    "crypto/md5"
    "fmt"
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func init() {
    dbhost := beego.AppConfig.String("dbhost")
    dbport := beego.AppConfig.String("dbport")
    dbuser := beego.AppConfig.String("dbuser")
    dbpassword := beego.AppConfig.String("dbpassword")
    dbname := beego.AppConfig.String("dbname")
    if dbport == "" {
        dbport = "3306"
    }
    //&loc=Asia%2FShanghai（已单独做了时区自动处理）
    dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
    orm.RegisterDataBase("default", "mysql", dburl)
    orm.RegisterModel(
        new(User))
    if beego.AppConfig.String("runmode") == "dev" {
        orm.Debug = true
    }
    //orm.RunSyncdb("default", false, true)
}

func Md5(buf []byte) string {
    hash := md5.New()
    hash.Write(buf)
    return fmt.Sprintf("%x",hash.Sum(nil))
}

//返回带前缀的表名
func TableName(str string) string{
    return fmt.Sprintf("%s%s",beego.AppConfig.String("dbprefix"),str)
}