package models

import "github.com/astaxie/beego/orm"

type User struct {
    Id          int64       `json:"id"`
    Name        string      `json:"name" orm:"unique;size(20);index"`
    Pwd         string      `orm:"size(32)"`
    Avatar      string      `json:"avatar" orm:"size(255);default(/static/upload/default/user-default-60x60.png)"`
    Email       string      `json:"email" orm:size(50);index`
    Role        int         `json:"role" orm:size(2)`
}

func (m *User) TableName() string {
    return TableName("user")
}

//插入
func (m *User) Insert() error {
    if _, err := orm.NewOrm().Insert(m);err != nil {
        return err
    }
    return nil
}

//读数据
func (m *User) Read(fields ...string) error {
    if err := orm.NewOrm().Read(m,fields...);err != nil {
        return err
    }
    return nil
}

//更新数据      /*...string 是什么意思？ 需要去了解一下*/
func (m *User) Update(fields ...string) error {
    if _,err := orm.NewOrm().Update(m, fields...);err != nil{
        return err
    }
    return nil
}

//删除数据
func (m *User) Delete() error {
    if _, err := orm.NewOrm().Delete(m); err != nil {
        return err
    }
    return nil
}

//获取数据？
func (m *User) Query() orm.QuerySeter {
    return orm.NewOrm().QueryTable(m)
}