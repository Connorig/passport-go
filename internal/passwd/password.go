package passwd

import (
	"github.com/lishimeng/passport-go/internal/db/model"
	"github.com/lishimeng/x/digest"
)

// 生成密码
func Generate(plaintext string, ds model.UserInfo) (r string) {
	r = digest.Generate(plaintext, ds.CreateTime.Unix()/1000)
	return
}

// 校验密码
func Verify(plaintext string, ds model.UserInfo) (r bool) {
	r = digest.Verify(plaintext, ds.Password, ds.CreateTime.Unix()/1000)
	return
}
