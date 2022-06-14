package middleware

import (
	"cloud-disk/core/helper"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqeustPath := r.RequestURI
		logx.Infof("请求地址:%s", reqeustPath)
		if "/userLogin" == reqeustPath {
			// Passthrough to next handler if need
			next(w, r)
			return
		}
		auth := r.Header.Get("Authorization")
		// 为空则返回未授权
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		//解析token
		uc, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		r.Header.Set("UserId", string(rune(uc.Id)))
		r.Header.Set("UserIdentity", uc.Identity)
		r.Header.Set("UserName", uc.Name)

		next(w, r)
	}
}
