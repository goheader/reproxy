package client

import "reproxy/pkg/config"

type Control struct {
	runID string
	pxyCfgs map[string]config.ProxyConf
	pm *proxy.
}