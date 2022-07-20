package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
	"reproxy/pkg/auth"
	"strings"
)

type ClientCommonConf struct {
	auth.ClientConfig `ini:",extends"`

	ServerAddr string `ini:"server_addr" json:"server_addr"`
	ServerPort int    `ini:"server_port" json:"server_port"`

	DialServerTimeout   int64 `ini:"dial_server_timeout" json:"dial_server_timeout"`
	DialServerKeepAlive int64 `ini:"dial_server_keepalive" json:"dial_server_keepalive"`

	ConnectServerLocalIP string `ini:"connect_server_local_ip" json:"connect_server_local_ip"`
	HTTPProxy            string `ini:"http_proxy" json:"http_proxy"`

	LogFile         string `ini:"log_file" json:"log_file"`
	LogWay          string `ini:"log_way" json:"log_way"`
	LogLevel        string `ini:"log_level" json:"log_level"`
	LogMaxDays      int64  `ini:"log_max_days" json:"log_max_days"`
	DisableLogColor bool   `ini:"disable_log_color" json:"disable_log_color"`

	AdminAddr string `ini:"admin_addr" json:"admin_addr"`
	AdminPort int    `ini:"admin_port" json:"admin_port"`
	AdminUser string `ini:"admin_user" json:"admin_user"`
	AdminPwd  string `ini:"admin_pwd" json:"admin_pwd"`

	AssetsDir string `ini:"assets_dir" json:"assets_dir"`
	PoolCount int    `ini:"pool_count" json:"pool_count"`

	TCPMux                  bool   `ini:"tcp_mux" json:"tcp_mux"`
	TCPMuxKeepaliveInterval int64  `ini:"tcp_mux_keepalive_interval" json:"tcp_mux_keepalive_interval"`
	User                    string `ini:"user" json:"user"`
	DNSServer               string `ini:"dns_server" json:"dns_server"`

	LoginFailExit bool `ini:"login_fail_exit" json:"login_fail_exit"`

	Start                     []string          `ini:"start" json:"start"`
	Protocol                  string            `ini:"protocol" json:"protocol"`
	TLSEnable                 bool              `ini:"tls_enable" json:"tls_enable"`
	TLSCertFile               string            `ini:"tls_cert_file" json:"tls_cert_file"`
	TLSKeyFile                string            `ini:"tls_key_file" json:"tls_key_file"`
	TLSTrustedCaFile          string            `ini:"tls_trusted_ca_file" json:"tls_trusted_ca_file"`
	TLSServerName             string            `ini:"tls_server_name" json:"tls_server_name"`
	DisableCustomTLSFirstByte bool              `ini:"disable_custom_tls_first_byte" json:"disable_custom_tls_first_byte"`
	HeartbeatInterval         int64             `ini:"heartbeat_interval" json:"heartbeat_interval"`
	HeartbeatTimeout          int64             `ini:"heartbeat_timeout" json:"heartbeat_timeout"`
	Metas                     map[string]string `ini:"metas" json:"metas"`
	UDPPacketSize             int64             `ini:"udp_packet_size" json:"udp_packet_size"`
	IncludeConfigFiles        []string          `ini:"includes" json:"includes"`
	PprofEnable               bool              `ini:"pprof_enable" json:"pprof_enable"`
}

func UnmarshalClientConfFromIni(source interface{}) (ClientCommonConf, error) {
	f, err := ini.LoadSources(ini.LoadOptions{
		Insensitive:         false,
		InsensitiveSections: false,
		InsensitiveKeys:     false,
		IgnoreInlineComment: true,
		AllowBooleanKeys:    true,
	}, source)

	if err != nil {
		return ClientCommonConf{},err
	}

	s,err := f.GetSection("common")
	if err != nil {
		return ClientCommonConf{}, fmt.Errorf("invalid configuration file,not found [common] section")
	}

	common := GetDefaultClientConf()
	err = s.MapTo(&common)
	if err != nil {
		return ClientCommonConf{},err
	}

	common.Metas = GetMapWithoutPrefix(s.KeysHash(),"meta_")
	common.ClientConfig.OidcAdditionalEndpointParams = GetMapWithoutPrefix(s.KeysHash(),"oidc_additional_")

	return common,nil
}

func GetDefaultClientConf() ClientCommonConf {
	return ClientCommonConf{
		ClientConfig:            auth.GetDefaultClientConf(),
		ServerAddr:              "0.0.0.0",
		ServerPort:              7000,
		DialServerTimeout:       10,
		DialServerKeepAlive:     7200,
		HTTPProxy:               os.Getenv("http_proxy"),
		LogFile:                 "console",
		LogWay:                  "console",
		LogLevel:                "info",
		LogMaxDays:              3,
		DisableLogColor:         false,
		AdminAddr:               "127.0.0.1",
		AdminPort:               0,
		AdminUser:               "",
		AdminPwd:                "",
		AssetsDir:               "",
		PoolCount:               1,
		TCPMux:                  true,
		TCPMuxKeepaliveInterval: 60,
		User:                    "",
		DNSServer:               "",
		LoginFailExit:           true,
		Start: make([]string,0),
		Protocol: "tcp",
		TLSEnable: false,
		TLSCertFile: "",
		TLSKeyFile: "",
		TLSTrustedCaFile: "",
		HeartbeatInterval: 30,
		HeartbeatTimeout: 90,
		Metas: make(map[string]string),
		UDPPacketSize: 1500,
		IncludeConfigFiles: make([]string,0),
		PprofEnable: false,
	}
}

func (cfg *ClientCommonConf) Complete(){
	if cfg.LogFile == "console"{
		cfg.LogWay = "console"
	}else{
		cfg.LogWay = "file"
	}
}

func (cfg *ClientCommonConf) Validate() error{
	if cfg.HeartbeatTimeout >0 && cfg.HeartbeatInterval>0 {
		if cfg.HeartbeatTimeout < cfg.HeartbeatInterval{
			return fmt.Errorf("invalid heartbeat_timeout,heartbeat_timeout is less than heartbeat_interval")
		}
	}

	if cfg.TLSEnable == false{
		if cfg.TLSCertFile != ""{
			fmt.Println("WARNING! tls_cert_file is invalid when tls_enable is false")
		}

		if cfg.TLSKeyFile != ""{
			fmt.Println("WARNING! tls_key_file is invalid when tls_enable is false")
		}

		if cfg.TLSTrustedCaFile != ""{
			fmt.Println("WARNING! tls_trusted_ca_file is invalid when tls_enable is false")
		}
	}

	if cfg.Protocol != "tcp" && cfg.Protocol != "kcp" && cfg.Protocol != "websocket" {
		return fmt.Errorf("invalid protocol")
	}

	for _,f := range cfg.IncludeConfigFiles{
		absDir,err := filepath.Abs(filepath.Dir(f))
		if err != nil {
			return fmt.Errorf("include: parse directory of %s failed: %v",f,absDir)
		}
		if _,err := os.Stat(absDir); os.IsNotExist(err){
			return fmt.Errorf("include: directory of %s not exist",f)
		}
	}
	return nil
}


func LoadAllProxyConfsFromIni(prefix string,source interface{},start []string) (map[string]ProxyConf,map[string]VisitorConf,error){
	f,err := ini.LoadSources(ini.LoadOptions{
		Insensitive: false,
		InsensitiveSections: false,
		InsensitiveKeys: false,
		IgnoreInlineComment: true,
		AllowBooleanKeys: true,
	},source)
	if err != nil {
		return nil,nil,err
	}

	proxyConfs := make(map[string]ProxyConf)
	visitorConfs := make(map[string]VisitorConf)

	if prefix != "" {
		prefix += "."
	}

	startProxy := make(map[string]struct{})
	for _,s := range start {
		startProxy[s] = struct{}{}
	}

	startAll := true
	if len(startProxy) >0 {
		startAll = false
	}

	rangeSections := make([]*ini.Section,0)
	for _,section := range f.Sections(){
		name := section.Name()

		if name == ini.DefaultSection || name == "common" || strings.HasPrefix(name,"range:"){
			continue
		}
		_,shouldStart := startProxy[name]
		if !startAll && !shouldStart {
			continue
		}
		roleType := section.Key("role").String()

		if roleType == ""{
			roleType = "server"
		}
		switch roleType {
		case "server":
			newConf,newErr := NewPro
		}
	}

}

