/**********************************************
** @Des: main
** @Author: 云飞扬
** @Date:
** @Last Modified by:   云飞扬
** @Last Modified time:
***********************************************/
package main

import (
	"github.com/astaxie/beego"
	"github.com/yunfeiyang1916/cloud-go/models"
	_ "github.com/yunfeiyang1916/cloud-go/routers"
)

func main() {
	models.Init()
	beego.Run()
}
