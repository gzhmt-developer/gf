package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
)

type RegisterReq struct {
	Name  string `p:"username"  v:"required|length:6,30#请输入账号|账号长度为:min到:max位"`
	Pass  string `p:"password1" v:"required|length:6,30#请输入密码|密码长度不够"`
	Pass2 string `p:"password2" v:"required|length:6,30|same:password1#请确认密码|两次密码不一致"`
}

type RegisterRes struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func main() {
	s := g.Server()
	s.BindHandler("/register", func(r *ghttp.Request) {
		var req *RegisterReq
		if err := r.Parse(&req); err != nil {
			// Validation error.
			if v, ok := err.(gvalid.Error); ok {
				r.Response.WriteJsonExit(RegisterRes{
					Code:  1,
					Error: v.FirstString(),
				})
			}
			// Other error.
			r.Response.WriteJsonExit(RegisterRes{
				Code:  1,
				Error: err.Error(),
			})
		}
		// ...
		r.Response.WriteJsonExit(RegisterRes{
			Data: req,
		})
	})
	s.SetPort(8199)
	s.Run()
}
