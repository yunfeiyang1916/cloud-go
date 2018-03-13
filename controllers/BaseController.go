/**********************************************
** @Des: 控制器基类
** @Author: 云飞扬
** @Date:
** @Last Modified by:   云飞扬
** @Last Modified time:
***********************************************/
package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/yunfeiyang1916/cloud-go/models"
)

//控制器基类
type BaseController struct {
	beego.Controller
	//控制器名称
	controllerName string
	//操作名称
	actionName string
	//用户
	user *models.Admin
	//用户id
	userId int
	//用户名
	userName string
	//登陆名
	loginName string
	//每页记录条数
	pageSize int
	allowUrl string
}

//前期准备
func (this *BaseController) Prepare() {
	this.pageSize = 20
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = actionName
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("siteName")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	fmt.Println(this.controllerName)

	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
}

//重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//获取客户端ip
func (this *BaseController) getClientIP() string {
	ip := strings.Split(this.Ctx.Request.RemoteAddr, ":")
	return ip[0]
}
