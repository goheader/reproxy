package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"reflect"
	"reproxy/pkg/consts"
)

var (
	visitorConfTypeMap = map[string]reflect.Type{
		consts.STCPProxy: reflect.TypeOf(STCPVisitorConf{}),
		consts.XTCPProxy: reflect.TypeOf(XTCPVisitorConf{}),
		consts.SUDPProxy: reflect.TypeOf(SUDPVisitorConf{}),
	}
)

type SUDPVisitorConf struct {
	BaseVisitorConf `ini:",extends"`
}

type STCPVisitorConf struct {
	BaseVisitorConf `ini:",extends"`
}

type XTCPVisitorConf struct {
	BaseVisitorConf `ini:",extends"`
}

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

func NewVisitorConfFromIni(prefix string,name string, section *ini.Section) (VisitorConf,error){
	visitorType := section.Key("type").String()

	if visitorType == ""{
		return nil,fmt.Errorf("visitor [%s] type shouldn't be empty",name)
	}

	conf := DefaultVisitorConf(visitorType)
	if conf == nil{
		return nil,fmt.Errorf("visitor [%s] type [%s] error",name,visitorType)
	}
	if err := conf.UnmarshalFromIni(prefix,name,section);err != nil{
		return nil,fmt.Errorf("visitor [%s] type [%s] error",name,visitorType)
	}
	if err := conf.Check();err != nil{
		return nil,err
	}
	return conf,nil
}

func DefaultVisitorConf(visitorType string) VisitorConf{
	v,ok := visitorConfTypeMap[visitorType]
	if !ok {
		return nil
	}
	return reflect.New(v).Interface().(VisitorConf)
}
