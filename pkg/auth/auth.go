package auth

import "reproxy/pkg/msg"

type Setter interface {
	SetLogin(*msg.Login) error
	SetPing(*msg.Ping) error
	SetNewWorkConn(*msg.NewWorkConn) error
}



type ClientConfig struct {
	BaseConfig `ini:",extends"`
	OidcClientConfig `ini:",extends"`
	TokenConfig `ini:",extends"`
}


type BaseConfig struct {
	AuthenticationMethod string `ini:"authentication_method" json:"authentication_method"`
	AuthenticateHeartBeats bool `ini:"authenticate_heartbeats" json:"authenticate_heartbeats"`
	AuthenticateNewWorkConns bool `ini:"authenticate_new_work_conns" json:"authenticate_new_work_conns"`

}


