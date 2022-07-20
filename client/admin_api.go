package client

import (
	"github.com/fatedier/beego/logs"
	"net/http"
	"reproxy/pkg/config"
)

func (svr *Service) healthz(w http.ResponseWriter,r *http.Request){
	w.WriteHeader(200)
}

func (svr *Service) apiReload(w http.ResponseWriter,r *http.Request){
	res := GeneralResponse{Code: 200}

	logs.Info("api request [/api/reload]")
	defer func() {
		logs.Info("api response [/api/reload],code [%d]",res.Code)
		w.WriteHeader(res.Code)
		if len(res.Msg)>0{
			w.Write([]byte(res.Msg))
		}
	}()
	_,pxyCfgs,visitorCfgs,err := config.ParseClientConfig(svr.cfgFile)

	if
}


type GeneralResponse struct{
	Code int
	Msg string
}
