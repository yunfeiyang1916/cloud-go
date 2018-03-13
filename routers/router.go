/**********************************************
** @Des: 路由
** @Author: 云飞扬
** @Date:
** @Last Modified by:   云飞扬
** @Last Modified time:
***********************************************/
package routers

import (
	"github.com/astaxie/beego"
	"github.com/yunfeiyang1916/cloud-go/controllers"
)

func init() {
	beego.Router("/login", &controllers.LoginController{}, "*:Index")
	beego.Router("/login/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/home", &controllers.HomeController{}, "*:Index")
}
