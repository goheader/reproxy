package client

import (
	"github.com/fatedier/golib/crypto"
	"math/rand"
	"reproxy/pkg/auth"
	"sync"
	"time"
)

func init(){
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())
}


type Service struct {
	runID string
	ctl *Control
	ctlMu sync.RWMutex
	authSetter auth.Setter

	cfg config.Client
}