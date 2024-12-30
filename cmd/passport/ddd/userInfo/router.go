package userInfo

import (
	"github.com/lishimeng/app-starter/server"
)

func Route(root server.Router) {
	root.Get("/getUserInfo", GetUserInfo)
	root.Post("/bindUser", BindUser)   //绑定用户名密码
	root.Post("/bindPhone", BindPhone) //换绑手机
	root.Post("/bindEmail", BindEmail) //换绑邮箱
}
