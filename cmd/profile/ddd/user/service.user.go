package user

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/passport-go/internal/db/model"
)

func GetUserInfoByCode(code string) (info model.UserInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.UserInfo)).Filter("user_code", code).One(&info)
	return
}

func UpUserInfo(ori model.UserInfo, cols ...string) (info model.UserInfo, err error) {
	_, err = app.GetOrm().Context.Update(&ori, cols...)
	info = ori
	return
}
func UpdateUserInfo(user model.UserInfo, cols ...string) (info model.UserInfo, err error) {
	_, err = app.GetOrm().Context.Update(&user, cols...)
	info = user
	return
}
