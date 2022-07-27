package client

import (
	"net/http"
	"reproxy/pkg/config"
	"reproxy/pkg/util/log"
)

func (svr *Service) healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (svr *Service) apiReload(w http.ResponseWriter, r *http.Request) {
	res := GeneralResponse{Code: 200}

	log.Info("api request [/api/reload]")
	defer func() {
		log.Info("api response [/api/reload],code [%d]", res.Code)
		w.WriteHeader(res.Code)
		if len(res.Msg) > 0 {
			w.Write([]byte(res.Msg))
		}
	}()
	_, pxyCfgs, visitorCfgs, err := config.ParseClientConfig(svr.cfgFile)

	if err != nil {
		res.Code = 400
		res.Msg = err.Error()
		log.Warn("reload frpc proxy config error: %s",res.Msg)
		return
	}

	if err = svr.Reload
}

type GeneralResponse struct {
	Code int
	Msg  string
}
