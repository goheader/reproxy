package config

import (
	"gopkg.in/ini.v1"
	"reflect"
	"reproxy/pkg/consts"
	"reproxy/pkg/msg"
)

var (
	proxyConfTypeMap = map[string]reflect.Type{
		consts.TCPProxy:    reflect.TypeOf(TCPProxyConf{}),
		consts.TCPMuxProxy: reflect.TypeOf(TCPProxyConf{}),
		consts.UDPProxy:    reflect.TypeOf(UDPProxyConf{}),
		consts.STCPProxy:   reflect.TypeOf(STCPProxyConf{}),
		consts.HTTPProxy:   reflect.TypeOf(HTTPProxyConf{}),
		consts.HTTPSProxy:  reflect.TypeOf(HTTPSProxyConf{}),
		consts.STCPProxy:   reflect.TypeOf(STCPProxyConf{}),
		consts.XTCPProxy:   reflect.TypeOf(XTCPProxyConf{}),
		consts.SUDPProxy:   reflect.TypeOf(SUDPProxyConf{}),
	}
)

type ProxyConf interface {
	GetBaseInfo() *BaseProxyConf
	UnmarshalFromMsg(*msg.NewProxy)
	UnmarshalFromIni(string,string,*ini.Section) error
	MarshalToMsg(*msg.NewProxy)
	CheckForCli() error
	CheckForSvr(ServerCommonConf) error
	Compare(ProxyConf) bool

}


//TCP
type TCPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	RemotePort    int `ini:"remote_port" json:"remote_port"`
}

//UDP
type UDPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	RemotePort    int `ini:"remote_port" json:"remote_port"`
}

//SUDP
type SUDPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	Role          string `ini:"role" json:"role"`
	Sk            string `ini:"sk" json:"sk"`
}

//XTCP

type XTCPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	Role          string `ini:"role" json:"role"`
	Sk            string `ini:"sk" json:"sk"`
}

//STCP
type STCPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	Role          string `ini:"role" json:"role"`
	Sk            string `ini:"sk" json:"sk"`
}

//HTTP
type HTTPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	DomainConf    `ini:",extends"`

	Locations         []string          `ini:"locations" json:"locations"`
	HTTPUser          string            `ini:"http_user" json:"http_user"`
	HTTPPwd           string            `ini:"http_pwd" json:"http_pwd"`
	HostHeaderRewrite string            `ini:"host_header_rewrite" json:"host_header_rewrite"`
	Headers           map[string]string `ini:"-" json:"headers"`
	RouteByHTTPUser   string            `ini:"route_by_http_user" json:"route_by_http_user"`
}

//HTTPS
type HTTPSProxyConf struct {
	BaseProxyConf `ini:",extends"`
	DomainConf    `ini:",extends"`
}

//TCPMux
type TCPMuxProxyConf struct {
	BaseProxyConf   `ini:",extends`
	DomainConf      `ini:",extends"`
	RouteByHTTPUser string `ini:"route_by_http_user" json:"route_by_http_user"`

	Multiplexer string `ini:"multiplexer"`
}

type BaseProxyConf struct {
	ProxyName      string `ini:"name" json:"name"`
	ProxyType      string `ini:"type" json:"type"`
	UseEncryption  bool   `ini:"use_encryption" json:"use_encryption"`
	UseCompression bool   `ini:"use_compression" json:"use_compression"`
	Group          string `ini:"group" json:"group"`
	GroupKey       string `ini:"groupkey" json:"groupkey"`

	ProxyProtocolVersion string `ini:"proxy_protocol_version" json:"proxy_protocol_version"`

	BandWidthLimit BandWidthQuantity `ini:"bandwidth_limit" json:"bandwidth_limit"`

	Metas map[string]string `ini:"-" json:"metas"`

	LocalSvrConf    `ini:",extends"`
	HealthCheckConf `ini:",extends"`
}

type LocalSvrConf struct {
	LocalIP      string            `ini:"local_ip" json:"local_ip"`
	LocalPort    string            `ini:"local_port" json:"local_port"`
	Plugin       string            `ini:"plugin" json:"plugin"`
	PluginParams map[string]string `ini:"-"`
}

type HealthCheckConf struct {
	HealthCheckType      string `ini:"health_check_type" json:"health_check_type"`
	HealthCheckTimeoutS  int    `ini:"health_check_timeout_s" json:"health_check_timeout_s"`
	HealthCheckMaxFailed int    `ini:"health_check_max_failed" json:"health_check_max_failed"`
	HealthCheckIntervalS int    `ini:"health_check_interval_s" json:"health_check_interval_s"`
	HealthCheckURL       string `ini:"health_check_url" json:"health_check_url"`
	HealthCheckAddr      string `ini:"health_check_addr" json:"health_check_addr"`
}

type DomainConf struct {
	CustomDomains []string `ini:"custom_domains" json:"custom_domains"`
	SubDomain     string   `ini:"sub_domain" json:"sub_domain"`
}


func defaultBaseProxyConf(proxyType string) BaseProxyConf{
	return BaseProxyConf{
		ProxyType: proxyType,
		LocalSvrConf: LocalSvrConf{
			LocalIP: "127.0.0.1",
		},
	}
}



func DefaultProxyConf(proxyType string) ProxyConf{
	var conf ProxyConf
	switch proxyType {
	case consts.TCPProxy:
		conf = &TCPProxyConf{
			BaseProxyConf: defaultBaseProxyConf(proxyType),
		}
	case consts.TCPMuxProxy:
		conf = &TCPMuxProxyConf{
			BaseProxyConf: defaultBaseProxyConf(proxyType),
		}
	case consts.UDPProxy:
		conf = &UDPProxyConf{
			BaseProxyConf: defaultBaseProxyConf(proxyType),
		}
	default:
		return nil
	}

	return conf

	}
}





func NewProxyConfFromIni(prefix,name string,section *ini.Section) (ProxyConf,error){
	proxyType := section.Key("type").String()
	if proxyType == "" {
		proxyType = consts.TCPProxy
	}
	conf := DefaultPro
}
