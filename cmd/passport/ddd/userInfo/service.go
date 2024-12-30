package userInfo

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport-go/internal/db/model"
)

func GetUserInfoByUserName(userName string) (info model.UserInfo, err error) {
	cond := orm.NewCondition()
	cond = cond.Or("UserName", userName)
	cond = cond.Or("Phone", userName)
	cond = cond.Or("Email", userName)
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).SetCond(cond).One(&info)
	return
}
func GetUserInfoByCode(code string) (info model.UserInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).Filter("UserCode", code).One(&info)
	return
}
func UpdateUserInfo(user model.UserInfo, cols ...string) (info model.UserInfo, err error) {
	_, err = app.GetOrm().Context.Update(&user, cols...)
	info = user
	return
}
