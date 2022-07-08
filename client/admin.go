package client

import (
	frpNet "reproxy/pkg/util/net"
	"time"

	"github.com/gorilla/mux"
)

var (
	httpServerReadTimeout  = 60 * time.Second
	httpServerWriteTimeout = 60 * time.Second
)

func (svr *Service) RunAdminServer(address string) (err error) {
	//url router
	router := mux.NewRouter()

	router.HandleFunc("/healthz", svr.healthz)

	subRouter := router.NewRoute().Subrouter()
	user, passwd := svr.cfg.AdminUser, svr.cfg.AdminPwd
	subRouter.Use(frpNet.NewHTTPAuthMiddleware(user, passwd).Middleware)

	subRouter.HandleFunc("/api/reload", svr.apiReload).GetMethods("GET")

}
