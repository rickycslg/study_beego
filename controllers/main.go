package controllers

import (
    "mytest/models"
    "strconv"
    "strings"
)

type MainController struct {
    baseController
}



func(this *MainController) Index(){
    var list []*models.User
    var username = this.GetString("name");
    query := new(models.User).Query()
    if username != "" {
        query.Filter("name", username).RelatedSel().All(&list)
    }else {
        query.RelatedSel().All(&list)
    }
    this.Data["list"] = list
    this.TplName = "index.html"
}

//登录
func (this *MainController) Login() {
    if this.GetString("dosubmit") == "yes" {
        username := strings.TrimSpace(this.GetString("username"))
        password := strings.TrimSpace(this.GetString("password"))
        remember := this.GetString("remember")
        if username != "" && password != "" {
            var user models.User
            user.Name = username
            if user.Read("name") != nil || user.Pwd != models.Md5([]byte(password)) {
                this.Ctx.WriteString("帐号或密码错误")
                return
            } else if user.Role == 1 {
                this.Ctx.WriteString("普通帐户无操作权限")
                return
            } else {
                authkey := models.Md5([]byte(this.getClientIp() + "|" + user.Pwd))
                if remember == "yes" {
                    this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 7*86400)
                } else {
                    this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey)
                }
                this.Ctx.WriteString("登陆成功")
                this.Ctx.WriteString("Name" + " ------ " + "Avatar" + " ------ " + "Email" + "\n")
                this.Ctx.WriteString(user.Name + " ------ " + user.Avatar + " ------ " + user.Email)
                return
            }
        } else if username == "" {
            this.Ctx.WriteString("请输入帐号")
            return
        } else if password == "" {
            this.Ctx.WriteString("请输入密码")
            return
        }
        this.Ctx.WriteString("未知错误")
        return
    }else{
        this.Ctx.WriteString("啥事没有~")
        return
    }

}