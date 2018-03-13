/**********************************************
** @Des: 用户
** @Author: 云飞扬
** @Date:
** @Last Modified by:   云飞扬
** @Last Modified time:
***********************************************/
package models

import "github.com/astaxie/beego/orm"

//用户表名
const TableNameAdmin string = "pp_uc_admin"

//用户
type Admin struct {
	//id
	Id int
	//登陆名
	LoginName string
	//真实名称
	RealName string
	//密码
	Password string
	//所属角色id串
	RoleIds string
	//电话
	Phone string
	//邮箱
	Email string
	//盐值
	Salt string
	//最后登陆时间
	LastLogin int64
	//最后登陆ip
	LastIp string
	//状态
	Status int
	//创建人id
	CreateId int
	//修改人id
	UpdateId int
	//创建时间
	CreateTime int64
	//更新时间
	UpdateTime int64
}

//获取表名，注册模型的时候会自动调用
func (this *Admin) TableName() string {
	return TableNameAdmin
}

//根据登陆名称获取用户信息
func GetAdminByLoginName(loginName string) (*Admin, error) {
	entity := new(Admin)
	orm := orm.NewOrm()
	err := orm.QueryTable(TableNameAdmin).Filter("login_name", loginName).One(entity)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

//更新
func (this *Admin) Update(fields ...string) error {
	_, err := orm.NewOrm().Update(this, fields...)
	if err != nil {
		return err
	}
	return nil
}
