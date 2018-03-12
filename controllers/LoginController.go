/**********************************************
** @Des: 登陆控制器
** @Author: 云飞扬
** @Date:
** @Last Modified by:   云飞扬
** @Last Modified time:
***********************************************/
package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

//登陆控制器
type LoginController struct {
	BaseController
}

//登陆页
func (this *LoginController) Index() {
	if this.userId > 0 {
		this.redirect(beego.URLFor("HomeController.Index"))
		return
	}
	this.TplName = "login/index.html"
}

//登陆
func (this *LoginController) Login() {
	userName := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))
	if len(userName) > 0 && len(password) > 0 {
		
	}
}
