package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"net"
	"reflect"
	"reproxy/pkg/consts"
	"reproxy/pkg/msg"
	"strconv"
	"strings"
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

func (cfg *TCPProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.unmarshalFromMsg(pMsg)
}

func (cfg *TCPProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err !=nil {
		return err
	}
	//Add custom logic unmarshal if exists
	return nil
}

func (cfg *TCPProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	pMsg.RemotePort = cfg.RemotePort
}

func (cfg *TCPProxyConf) CheckForCli() (err error) {
	if err := cfg.BaseProxyConf.checkForCli();err != nil{
		return
	}
	return
}

func (cfg *TCPProxyConf) CheckForSvr(serverCfg ServerCommonConf) error {
	return nil
}


func (cfg *TCPProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf,ok := cmp.(*TCPProxyConf)
	if !ok {
		return false
	}
	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf) {
		return false
	}
	if cfg.RemotePort != cmpConf.RemotePort {
		return false
	}
	return true
}

func preUnmarshalFromIni(cfg ProxyConf,prefix string, name string, section *ini.Section) error{
	err := section.MapTo(cfg)
	if err !=nil {
		return err
	}

	err = cfg.GetBaseInfo().decorate(prefix,name,section)
	if err !=nil {
		return err
	}
	return nil
}

func (cfg *BaseProxyConf) checkForCli() (err error) {
	if cfg.ProxyProtocolVersion != "" {
		if cfg.ProxyProtocolVersion != "v1" && cfg.ProxyProtocolVersion != "v2" {
			return fmt.Errorf("no support proxy protocol version: %s",cfg.ProxyProtocolVersion)
		}
	}
	if err = cfg.LocalSvrConf.checkForCli();err != nil {
		return
	}
	if err = cfg.HealthCheckConf.checkForCli(); err != nil {
		return
	}
	return nil
}

func (cfg *HealthCheckConf) checkForCli() (err error) {
	if cfg.HealthCheckType != "" && cfg.HealthCheckType != "tcp" && cfg.HealthCheckType != "http" {
		return fmt.Errorf("unsupport health check type")
	}
	if cfg.HealthCheckType != "" {
		if cfg.HealthCheckType == "http" && cfg.HealthCheckURL == "" {
			return fmt.Errorf("health_check_url is required for health check type 'http'")
		}
	}
	return nil
}


func (cfg *LocalSvrConf) checkForCli() (err error) {
	if cfg.Plugin == "" {
		if cfg.LocalIP == "" {
			err = fmt.Errorf("local ip or plugin is required")
			return
		}
		if cfg.LocalPort <= 0 {
			err = fmt.Errorf("error local_port")
			return
		}
	}
	return
}

//UDP
type UDPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	RemotePort    int `ini:"remote_port" json:"remote_port"`
}

func (cfg *UDPProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.unmarshalFromMsg(pMsg)
	cfg.RemotePort = pMsg.RemotePort
}

func (cfg *UDPProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *UDPProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	pMsg.RemotePort = cfg.RemotePort
}

func (cfg *UDPProxyConf) CheckForCli() (err error) {
	if err = cfg.BaseProxyConf.checkForCli();err != nil {
		return
	}
	return
}

func (cfg *UDPProxyConf) CheckForSvr(conf ServerCommonConf) error {
	return nil
}

func (cfg *UDPProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf,ok := cmp.(*UDPProxyConf)
	if !ok {
		return false
	}
	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf){
		return false
	}
	if cfg.RemotePort != cmpConf.RemotePort{
		return false
	}
	return true
}

//SUDP
type SUDPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	Role          string `ini:"role" json:"role"`
	Sk            string `ini:"sk" json:"sk"`
}
// Only for role server
func (cfg *SUDPProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.unmarshalFromMsg(pMsg)
	cfg.Sk = pMsg.Sk

}

func (cfg *SUDPProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err != nil{
		return err
	}

	return nil
}

func (cfg *SUDPProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	pMsg.Sk = cfg.Sk
}

func (cfg *SUDPProxyConf) CheckForCli() error {
	if err := cfg.BaseProxyConf.checkForCli();err != nil{
		return err
	}
	if cfg.Role != "server" {
		return fmt.Errorf("role should be 'server' ")
	}
	return nil
}

func (cfg *SUDPProxyConf) CheckForSvr(serverCfg ServerCommonConf) error {
	return nil
}

func (cfg *SUDPProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf,ok := cmp.(*SUDPProxyConf)
	if !ok {
		return false
	}
	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf) {
		return false
	}
	if cfg.Role != cmpConf.Role ||
		cfg.Sk != cmpConf.Role{
		return false
	}
	return true

}

//XTCP

type XTCPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	Role          string `ini:"role" json:"role"`
	Sk            string `ini:"sk" json:"sk"`
}

func (cfg *XTCPProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.unmarshalFromMsg(pMsg)
	cfg.Sk = pMsg.Sk
}

func (cfg *XTCPProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err != nil{
		return err
	}
	//Add custom login equal if exists
	if cfg.Role == ""{
		cfg.Role = "server"
	}
	return nil
}

func (cfg *XTCPProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	pMsg.Sk = cfg.Sk
}

func (cfg *XTCPProxyConf) CheckForCli() (err error) {
	if err = cfg.BaseProxyConf.checkForCli(); err != nil{
		return
	}
	if cfg.Role != "server"{
		return fmt.Errorf("role should be 'server' ")
	}
	return
}

func (cfg *XTCPProxyConf) CheckForSvr(serverCfg ServerCommonConf) error {
	return nil
}

func (cfg *XTCPProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf, ok := cmp.(*XTCPProxyConf)
	if !ok {
		return false
	}
	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf){
		return false
	}
	if cfg.Role != cmpConf.Role ||
		cfg.Sk != cmpConf.Sk{
		return false
	}
	return true
}

//STCP
type STCPProxyConf struct {
	BaseProxyConf `ini:",extends"`
	Role          string `ini:"role" json:"role"`
	Sk            string `ini:"sk" json:"sk"`
}

func (cfg *STCPProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	cfg.Sk = pMsg.Sk

}

func (cfg *STCPProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err != nil {
		return err
	}
	if cfg.Role != ""{
		cfg.Role = "server"
	}
	return nil
}

func (cfg *STCPProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	pMsg.Sk = cfg.Sk
}

func (cfg *STCPProxyConf) CheckForCli() (err error) {
	if err = cfg.BaseProxyConf.checkForCli();err != nil {
		return
	}

	if cfg.Role != "server"{
		return fmt.Errorf("role should be 'server'")
	}
	return
}

func (cfg *STCPProxyConf) CheckForSvr(serverCfg ServerCommonConf) error {
	return nil
}

func (cfg *STCPProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf,ok := cmp.(*STCPProxyConf)
	if !ok {
		return false
	}

	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf){
		return false
	}

	// Add custom logic equal if exists
	if cfg.Role != cmpConf.Role || cfg.Sk != cmpConf.Sk {
		return false
	}
	return true

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

func (cfg *HTTPProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.unmarshalFromMsg(pMsg)
	cfg.CustomDomains = pMsg.CustomDomains
	cfg.SubDomain = pMsg.SubDomain
	cfg.Locations = pMsg.Locations
	cfg.HostHeaderRewrite = pMsg.HostHeaderRewrite
	cfg.HTTPUser = pMsg.HTTPUser
	cfg.HTTPPwd = pMsg.HTTPPwd
	cfg.Headers = pMsg.Headers
	cfg.RouteByHTTPUser = pMsg.RouteByHTTPUser
}

func (cfg *HTTPProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err != nil {
		return err
	}
	cfg.Headers = GetMapWithoutPrefix(section.KeysHash(),"header_")
	return nil
}

func (cfg *HTTPProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	pMsg.CustomDomains = cfg.CustomDomains
	pMsg.SubDomain = cfg.SubDomain
	pMsg.Locations = cfg.Locations
	pMsg.HostHeaderRewrite = cfg.HostHeaderRewrite
	pMsg.HTTPUser = cfg.HTTPUser
	pMsg.HTTPPwd = cfg.HTTPPwd
	pMsg.Headers = cfg.Headers
	pMsg.RouteByHTTPUser = cfg.RouteByHTTPUser

}

func (cfg *HTTPProxyConf) CheckForCli() (err error) {
	if err = cfg.BaseProxyConf.checkForCli();err != nil {
		return
	}
	if err = cfg.DomainConf.checkForCli();err != nil {
		return
	}
	return
}

func (cfg *HTTPProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf,ok := cmp.(*HTTPProxyConf)
	if !ok {
		return false
	}

	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf){
		return false
	}

	if !reflect.DeepEqual(cfg.DomainConf,cmpConf.DomainConf){
		return false
	}

	if !reflect.DeepEqual(cfg.Locations,cmpConf.Locations) ||
		cfg.HTTPUser != cmpConf.HTTPUser ||
		cfg.HTTPPwd != cmpConf.HTTPPwd ||
		cfg.HostHeaderRewrite != cmpConf.HostHeaderRewrite ||
		cfg.RouteByHTTPUser != cmpConf.RouteByHTTPUser ||
		!reflect.DeepEqual(cfg.Headers,cmpConf.Headers){
		return false
	}
	return true
}

//HTTPS
type HTTPSProxyConf struct {
	BaseProxyConf `ini:",extends"`
	DomainConf    `ini:",extends"`
}

func (cfg *HTTPSProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.unmarshalFromMsg(pMsg)
	cfg.CustomDomains = pMsg.CustomDomains
	cfg.SubDomain = pMsg.SubDomain
}

func (cfg *HTTPSProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err != nil{
		return err
	}

	return nil
}

func (cfg *HTTPSProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)
	pMsg.CustomDomains = cfg.CustomDomains
	pMsg.SubDomain = cfg.SubDomain
}

func (cfg *HTTPSProxyConf) CheckForCli() (err error) {
	if err = cfg.BaseProxyConf.checkForCli();err != nil{
		return
	}

	if err = cfg.DomainConf.checkForCli();err != nil{
		return
	}
	return
}

func (cfg *HTTPSProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf,ok := cmp.(*HTTPSProxyConf)
	if !ok {
		return false
	}

	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf) {
		return false
	}
	if !reflect.DeepEqual(cfg.DomainConf,cmpConf.DomainConf){
		return false
	}
	return true
}

func (cfg *HTTPSProxyConf) CheckForSvr(serverCfg ServerCommonConf) (err error) {
	if serverCfg.VhostHTTPSPort ==0{
		return fmt.Errorf("type [https] not support when vhost_https_port is not set")
	}
	if err = cfg.DomainConf.CheckForSvr(serverCfg);err != nil {
		err = fmt.Errorf("proxy [%s] domain conf check error: %v",cfg.ProxyName,err)
		return
	}
	return
}

//TCPMux
type TCPMuxProxyConf struct {
	BaseProxyConf   `ini:",extends`
	DomainConf      `ini:",extends"`
	RouteByHTTPUser string `ini:"route_by_http_user" json:"route_by_http_user"`

	Multiplexer string `ini:"multiplexer"`
}

func (cfg *TCPMuxProxyConf) UnmarshalFromMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.unmarshalFromMsg(pMsg)
	cfg.CustomDomains = pMsg.CustomDomains
	cfg.SubDomain = pMsg.SubDomain
	cfg.Multiplexer = pMsg.Multiplexer
	cfg.RouteByHTTPUser = pMsg.RouteByHTTPUser
}

func (cfg *TCPMuxProxyConf) UnmarshalFromIni(prefix string, name string, section *ini.Section) error {
	err := preUnmarshalFromIni(cfg,prefix,name,section)
	if err != nil {
		return err
	}

	//Add custom logic unmarshal if exists
	return nil
}


func (cfg *TCPMuxProxyConf) MarshalToMsg(pMsg *msg.NewProxy) {
	cfg.BaseProxyConf.marshalToMsg(pMsg)

	pMsg.CustomDomains = cfg.CustomDomains
	pMsg.SubDomain = cfg.SubDomain
	pMsg.Multiplexer = cfg.Multiplexer
	pMsg.RouteByHTTPUser = cfg.RouteByHTTPUser
}

func (cfg *TCPMuxProxyConf) CheckForCli()  (err error) {
	if err = cfg.BaseProxyConf.checkForCli();err != nil{
		return
	}
	if err = cfg.DomainConf.checkForCli();err != nil {
		return
	}
	if cfg.Multiplexer != consts.HTTPConnectTCPMultiplexer{
		return fmt.Errorf("parse conf error: incorrect  multiplexer [%s]",cfg.Multiplexer)
	}
	return
}

func (cfg *TCPMuxProxyConf) CheckForSvr(serverCfg ServerCommonConf) (err error) {
	if cfg.Multiplexer != consts.HTTPConnectTCPMultiplexer{
		return fmt.Errorf("proxy [%s] incorrect multiplexer [%s]",cfg.ProxyName,cfg.Multiplexer)
	}
	if cfg.Multiplexer == consts.HTTPConnectTCPMultiplexer && serverCfg.TCPMuxHTTPConnectPort ==0{
		return fmt.Errorf("proxy [%s] type [tcpmux] with multiplexer [httpconnect] requires tcpmux_httpconnect_port configuration",cfg.ProxyName)
	}
	if err = cfg.DomainConf.CheckForSvr(serverCfg);err != nil{
		err = fmt.Errorf("proxy [%s] domain conf check error: %v",cfg.ProxyName,err)
		return
	}
	return
}

func (cfg *TCPMuxProxyConf) Compare(cmp ProxyConf) bool {
	cmpConf,ok := cmp.(*TCPMuxProxyConf)
	if !ok {
		return false
	}

	if !cfg.BaseProxyConf.compare(&cmpConf.BaseProxyConf){
		return false
	}
	if !reflect.DeepEqual(cfg.DomainConf,cmpConf.DomainConf){
		return false
	}
	if cfg.Multiplexer != cmpConf.Multiplexer || cfg.RouteByHTTPUser != cmpConf.RouteByHTTPUser {
		return false
	}
	return true
}

func (cfg *DomainConf) checkForCli() (err error){
	if err = cfg.check();err != nil{
		return
	}
	return
}

func (cfg *DomainConf) check() (err error){
	if len(cfg.CustomDomains) == 0&& cfg.SubDomain == "" {
		err = fmt.Errorf("custom_domains and subdomain should set at least one of them")
		return
	}
	return
}

func (cfg *DomainConf) CheckForSvr(serverCfg ServerCommonConf) (err error){
	if err = cfg.check();err != nil{
		return
	}

	for _, domain := range cfg.CustomDomains{
		if serverCfg.SubDomainHost != "" && len(strings.Split(serverCfg.SubDomainHost,".")) < len(strings.Split(domain,".")){
			if strings.Contains(domain,serverCfg.SubDomainHost){
				return fmt.Errorf("custom domain [%s] should not belong to subdomain_host [%s]",domain,serverCfg.SubDomainHost)
			}
		}
	}
	if cfg.SubDomain != ""{
		if serverCfg.SubDomainHost == ""{
			return fmt.Errorf("subdomain is not supported because this feature is not enabled in remote frps")
		}
		if strings.Contains(cfg.SubDomain,".") || strings.Contains(cfg.SubDomain,"*"){
			return fmt.Errorf("'.' and '*' is not supported in subdomain")
		}
	}
	return nil

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
	LocalPort    int            `ini:"local_port" json:"local_port"`
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
	case consts.HTTPProxy:
		conf = &HTTPProxyConf{
		BaseProxyConf: defaultBaseProxyConf(proxyType),
		}
	case consts.HTTPSProxy:
		conf = &HTTPSProxyConf{
		BaseProxyConf: defaultBaseProxyConf(proxyType),
		}
	case consts.STCPProxy:
		conf = &STCPProxyConf{
		BaseProxyConf: defaultBaseProxyConf(proxyType),
		Role: "server",
		}
	case consts.XTCPProxy:
		conf = &XTCPProxyConf{
		BaseProxyConf: defaultBaseProxyConf(proxyType),
		Role: "server",
		}
	case consts.SUDPProxy:
		conf = &SUDPProxyConf{
		BaseProxyConf: defaultBaseProxyConf(proxyType),
		Role: "server",
		}
	default:
		return nil
	}
	return conf
	}





func NewProxyConfFromIni(prefix,name string,section *ini.Section) (ProxyConf,error){
	proxyType := section.Key("type").String()
	if proxyType == "" {
		proxyType = consts.TCPProxy
	}
	conf := DefaultProxyConf(proxyType)
	if conf == nil {
		return nil,fmt.Errorf("proxy %s has invalid type [%s]",name,proxyType)
	}
	if err := conf.UnmarshalFromIni(prefix,name,section);err != nil{
		return nil,err
	}
	if err := conf.CheckForCli();err != nil{
		return nil,err
	}
	return conf,nil
}


func (cfg *BaseProxyConf) GetBaseInfo() *BaseProxyConf{
	return cfg
}

func (cfg *BaseProxyConf) compare(cmp *BaseProxyConf) bool{
	if cfg.ProxyName != cmp.ProxyName || cfg.ProxyType != cmp.ProxyType || cfg.UseEncryption != cmp.UseEncryption ||
		cfg.UseCompression != cmp.UseCompression || cfg.Group != cmp.Group || cfg.GroupKey != cmp.GroupKey ||
		cfg.ProxyProtocolVersion != cmp.ProxyProtocolVersion || !cfg.BandWidthLimit.Equal(&cmp.BandWidthLimit) ||
		!reflect.DeepEqual(cfg.Metas, cmp.Metas) {
		return false
	}
	if !reflect.DeepEqual(cfg.LocalSvrConf,cmp.LocalSvrConf){
		return false
	}
	if !reflect.DeepEqual(cfg.HealthCheckConf,cmp.HealthCheckConf){
		return false
	}
	return true
}

func (cfg *BaseProxyConf) decorate(prefix,name string,section *ini.Section) error {
	cfg.ProxyName = prefix + name
	cfg.Metas = GetMapWithoutPrefix(section.KeysHash(), "meta_")

	if bandwidth, err := section.GetKey("bandwidth_limit"); err == nil {
		cfg.BandWidthLimit, err = NewBandwidthQuantity(bandwidth.String())
		if err != nil {
			return err
		}
	}

	cfg.LocalSvrConf.PluginParams = GetMapByPrefix(section.KeysHash(),"plugin_")

	if cfg.HealthCheckType == "tcp" && cfg.Plugin == ""{
		cfg.HealthCheckAddr = cfg.LocalIP + fmt.Sprintf(":%d",cfg.LocalPort)
	}

	if cfg.HealthCheckType == "http" && cfg.Plugin == "" && cfg.HealthCheckURL != "" {
		s := "http://" + net.JoinHostPort(cfg.LocalIP,strconv.Itoa(cfg.LocalPort))
		if !strings.HasPrefix(cfg.HealthCheckURL,"/"){
			s += "/"
		}
		cfg.HealthCheckURL = s + cfg.HealthCheckURL
	}
	return nil
}


func (cfg *BaseProxyConf) marshalToMsg(pMsg *msg.NewProxy){
	pMsg.ProxyName = cfg.ProxyName
	pMsg.ProxyType = cfg.ProxyType
	pMsg.UseEncryption = cfg.UseEncryption
	pMsg.UseCompression = cfg.UseCompression
	pMsg.Group = cfg.Group
	pMsg.GroupKey = cfg.GroupKey
	pMsg.Metas = cfg.Metas
}

func (cfg *BaseProxyConf) unmarshalFromMsg(pMsg *msg.NewProxy){
	cfg.ProxyName = pMsg.ProxyName
	cfg.ProxyType = pMsg.ProxyType
	cfg.UseEncryption = pMsg.UseEncryption
	cfg.UseCompression = pMsg.UseCompression
	cfg.Group  = pMsg.Group
	cfg.GroupKey = pMsg.GroupKey
	cfg.Metas = pMsg.Metas
}



