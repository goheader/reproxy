package client

import (
	"github.com/gorilla/mux"
	"time"
)

var (
	httpServerReadTimeout = 60 * time.Second
	httpServerWriteTimeout = 60 * time.Second
)

func (srv *Service) RunAdminServer(address string) (err error){
	//url router
	router := mux.NewRouter()

	router.HandleFunc("/healthz",srv.)

}