package config

import "reproxy/pkg/auth"

type ClientCommonConf struct {
	auth.ClientConfig `ini:",extends"`

	ServerAddr string `ini:"server_addr" json:"server_addr"`
	ServerPort int `ini:"server_port" json:"server_port"`

	DialServerTimeout int64 `ini:"dial_server_timeout" json:"dial_server_timeout"`
	DialServerKeepAlive int64 `ini:"dial_server_keepalive" json:"dial_server_keepalive"`

	ConnectServerLocalIP string `ini:"connect_server_local_ip" json:"connect_server_local_ip"`
	HTTPProxy string `ini:"http_proxy" json:"http_proxy"`

	LogFile string `ini:"log_file" json:"log_file"`
	LogWay string `ini:"log_way" json:"log_way"`
	LogLevel string `ini:"log_level" json:"log_level"`
	LogMaxDays int64 `ini:"log_max_days" json:"log_max_days"`
	DisableLogColor bool `ini:"disable_log_color" json:"disable_log_color"`

	AdminAddr string `ini:"admin_addr" json:"admin_addr"`
	AdminPort int `ini:"admin_port" json:"admin_port"`
	AdminUser string `ini:"admin_user" json:"admin_user"`
	AdminPwd string `ini:"admin_pwd" json:"admin_pwd"`


	AssetsDir string `ini:"assets_dir" json:"assets_dir"`
	PoolCount int `ini:"pool_count" json:"pool_count"`

	TCPMux bool `ini:"tcp_mux" json:"tcp_mux"`
	TCPMuxKeepaliveInterval int64 `ini:"tcp_mux_keepalive_interval" json:"tcp_mux_keepalive_interval"`
	User string `ini:"user" json:"user"`
	DNSServer string `ini:"dns_server" json:"dns_server"`

	LoginFailExit bool `ini:"login_fail_exit" json:"login_fail_exit"`


}
