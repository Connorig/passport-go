package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/send"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/signin"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/signup"
	"github.com/lishimeng/passport-go/cmd/passport/ddd/userInfo"
)

func Route(app server.Router) {
	route(app.Path("/api"))
}

func route(root server.Router) {
	signin.Route(root.Path("/signin"))
	signup.Route(root.Path("/signup"))
	send.Route(root.Path("/send"))
	userInfo.Route(root.Path("/userInfo"))
}
