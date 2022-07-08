package net

import "net/http"

type HTTPAuthWraper struct {
	h      http.Handler
	user   string
	passwd string
}

type HTTPAuthMiddleware struct {
	user   string
	passwd string
}

func NewHTTPAuthMiddleware(user, passwd string) *HTTPAuthMiddleware {
	return &HTTPAuthMiddleware{
		user:   user,
		passwd: passwd,
	}
}

func (aw *HTTPAuthWraper) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	user, passwd, hasAuth := r.BasicAuth()
	if (aw.user == "" && aw.passwd == "") || (hasAuth && user == aw.user && passwd == aw.passwd) {
		aw.h.ServeHTTP(w, r)
	} else {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func (authMid *HTTPAuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqUser, reqPasswd, hasAuth := r.BasicAuth()
		if (authMid.user == "" && authMid.passwd == "") || (hasAuth && reqUser == authMid.user && reqPasswd == authMid.passwd) {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set("WWWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	})
}
