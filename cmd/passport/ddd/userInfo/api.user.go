package userInfo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/passport-go/internal/db/model"
)

type UserResp struct {
	app.Response
	Item UserInfo `json:"item"`
}

type UserInfo struct {
	Id       int    `json:"id,omitempty"`
	UserCode string `json:"userCode,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

func GetUserInfo(ctx server.Context) {
	var resp UserResp
	ui := ctx.C.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("code:%s", jwt.Uid)
	userOne, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Item.Id = userOne.Id
	resp.Item.UserCode = userOne.UserCode
	resp.Item.Username = userOne.Username
	resp.Item.Password = userOne.Password
	resp.Item.Email = userOne.Email
	resp.Item.Phone = userOne.Phone
	ctx.Json(resp)
	return
}

type userInfoReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func BindUser(ctx server.Context) {
	var req userInfoReq
	var resp app.Response
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	ui := ctx.C.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("code:%s", jwt.Uid)
	userOne, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "未查到记录"
		ctx.Json(resp)
		return
	}
	userOne.Username = req.Username
	userOne.Password = req.Password
	_, err = UpdateUserInfo(userOne, "username", "password")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "绑定失败"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type BindPhoneReq struct {
	Phone string `json:"phone,omitempty"`
	Code  string `json:"code,omitempty"`
}

func BindPhone(ctx server.Context) {
	var req BindPhoneReq
	var resp app.Response
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	ui := ctx.C.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("userCode:%s", jwt.Uid)
	userOne, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户不存在"
		ctx.Json(resp)
		return
	}
	key := string(model.SmsBind) + req.Phone
	var value string
	err = app.GetCache().Get(key, &value)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码未发送，请重新发送！"
		ctx.Json(resp)
		return
	}
	log.Info("code:%s,%s", value, req.Code)
	if value != req.Code {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码不正确"
		ctx.Json(resp)
		return
	}
	var cols []string
	userOne.Phone = req.Phone
	cols = append(cols, "Phone")

	_, err = UpdateUserInfo(userOne, cols...)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "绑定失败"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type BindEmailReq struct {
	Email string `json:"email,omitempty"`
	Code  string `json:"code,omitempty"`
}

func BindEmail(ctx server.Context) {
	var req BindEmailReq
	var resp app.Response
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "json解析失败"
		ctx.Json(resp)
		return
	}
	ui := ctx.C.Values().Get(auth.UserInfoKey)
	jwt, ok := ui.(token.JwtPayload)
	if !ok {
		return
	}
	log.Debug("userCode:%s", jwt.Uid)
	userOne, err := GetUserInfoByCode(jwt.Uid)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "用户不存在"
		ctx.Json(resp)
		return
	}
	key := string(model.EmailBind) + req.Email
	var value string
	err = app.GetCache().Get(key, &value)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码未发送，请重新发送！"
		ctx.Json(resp)
		return
	}
	log.Info("code:%s,%s", value, req.Code)
	if value != req.Code {
		resp.Code = tool.RespCodeError
		resp.Message = "验证码不正确"
		ctx.Json(resp)
		return
	}
	var cols []string
	userOne.Email = req.Email
	cols = append(cols, "Email")

	_, err = UpdateUserInfo(userOne, cols...)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "绑定失败"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
