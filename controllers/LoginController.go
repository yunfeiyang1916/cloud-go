/**********************************************
** @Des: 登陆控制器
** @Author: 云飞扬
** @Date:
** @Last Modified by:   云飞扬
** @Last Modified time:
***********************************************/
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/yunfeiyang1916/cloud-go/models"
	"github.com/yunfeiyang1916/cloud-go/utils"
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
	//读取flash数据
	beego.ReadFromRequest(&this.Controller)
	this.TplName = "login/index.html"
}

//登陆
func (this *LoginController) Login() {
	userName := strings.TrimSpace(this.GetString("username"))
	password := strings.TrimSpace(this.GetString("password"))
	if len(userName) > 0 && len(password) > 0 {
		user, err := models.GetAdminByLoginName(userName)
		fmt.Println(user)
		flash := beego.NewFlash()
		errMsg := ""
		if err != nil {
			errMsg = "账户错误"
		} else if user.Password != utils.Md5([]byte(password+user.Salt)) {
			errMsg = "密码错误"
		} else {
			user.LastIp = this.getClientIP()
			user.LastLogin = time.Now().Unix()
			user.Update()
			authKey := utils.Md5([]byte(user.LastIp + "|" + user.Password + user.Salt))
			this.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authKey, 7*86400)
			this.redirect(beego.URLFor("HomeController.Index"))
		}
		//走到这说明有错误
		flash.Error(errMsg)
		flash.Store(&this.Controller)
		this.redirect(beego.URLFor("LoginController.Index"))
	}
}
