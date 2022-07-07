package client

import (
	"github.com/fatedier/golib/crypto"
	"math/rand"
	"time"
)

func init(){
	crypto.DefaultSalt = "frp"
	rand.Seed(time.Now().UnixNano())
}


type Service struct {
	runID string
	ctl *Control
}