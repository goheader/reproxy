package config

import "gopkg.in/ini.v1"

type VisitorConf interface {
	GetBaseInfo() *BaseVisitorConf
	Compare(cmp VisitorConf) bool
	UnmarshalFromIni(prefix string, name string, section *ini.Section) error
	Check() error
}

type BaseVisitorConf struct {
	ProxyName     string `ini:"proxy_name" json:"proxy_name"`
	ProxyType     string `ini:"proxy_type" json:"proxy_type"`
	UseEncryption bool   `ini:"use_encryption" json:"use_encryption"`
	Role          string `ini:"role" json:"role"`
	Sk            string `ini:"sk" json:"sk"`
	ServerName    string `ini:"server_name" json:"server_name"`
	BindAddr      string `ini:"bind_addr" json:"bind_addr"`
	BindPort      string `ini:"bind_port" json:"bind_port"`
}
