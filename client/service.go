package client

import (
	"context"
	"math/rand"
	"reproxy/pkg/auth"
	"reproxy/pkg/config"
	"sync"
	"time"

	"github.com/fatedier/golib/crypto"
)

func init() {
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())
}

type Service struct {
	runID      string
	ctl        *Control
	ctlMu      sync.RWMutex
	authSetter auth.Setter

	cfg         config.ClientCommonConf
	pxyCfgs     map[string]config.ProxyConf
	visitorCfgs map[string]config.VisitorConf

	cfgMu sync.RWMutex

	cfgFile       string
	serverUDPPort int
	exit          uint32 //0 means not exit

	ctx    context.Context
	cancel context.CancelFunc
}

func (svr *Service) ReloadConf(pxyCfgs map[string]config.ProxyConf,visitorCfgs map[string]config.VisitorConf) error {
	svr.cfgMu.Lock()
	svr.pxyCfgs = pxyCfgs
	svr.visitorCfgs = visitorCfgs
	svr.cfgMu.Unlock()

	svr.ctlMu.Lock()
	ctl := svr.ctl
	svr.ctlMu.Unlock()

	if ctl != nil {
		return svr.ctl.
	}
}


func (ctl *Control) ReloadConf(pxyCfgs map[string]config.ProxyConf,visitorCfgs map[string]config.VisitorConf) error{
	ctl.v
}
