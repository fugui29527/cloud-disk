package middleware

import (
	"cloud-disk/core/helper"
	"cloud-disk/core/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type AuthMiddleware struct {
	Config config.Config
}

func NewAuthMiddleware(c config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		Config: c,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqeustPath := r.RequestURI
		logx.Infof("请求地址:%s", reqeustPath)
		for _, s := range m.Config.IgnoreUrl {
			if "/userLogin" == s {
				// Passthrough to next handler if need
				next(w, r)
				return
			}
		}
		auth := r.Header.Get("Authorization")
		// 为空则返回未授权
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			//helper.NewFailResult(helper.FailAuthCode,"")
			return
		}
		//解析token
		uc, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		r.Header.Set("UserId", string(rune(uc.Id)))
		r.Header.Set("UserIdentity", uc.Identity)
		r.Header.Set("UserName", uc.Name)

		next(w, r)
	}
}
