package client

import "time"

var (
	httpServerReadTimeout = 60 * time.Second
	httpServerWriteTimeout = 60 * time.Second
)

func (srv *Service) RunAdminServer(address string) (err error){
	router := mux.new
}