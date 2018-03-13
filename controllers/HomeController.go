/**********************************************
** @Des: 主页控制器
** @Author: 云飞扬
** @Date:
** @Last Modified by:   云飞扬
** @Last Modified time:
***********************************************/
package controllers

//主页控制器
type HomeController struct {
	BaseController
}

//首页
func (this *HomeController) Index() {
	this.Data["pageTitle"] = "系统首页"
	this.TplName = "home/main.html"
}
