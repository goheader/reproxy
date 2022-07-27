package config

import plugin "reproxy/pkg/plugin/server"



type ServerCommonConf struct {

	BindAddr string `ini:"bind_addr" json:"bind_addr"`
	BindPort int `ini:"bind_port" json:"bind_port"`
	BindUDPPort int `ini:"bind_udp_port" json:"bind_udp_port"`
	KCPBindPort int `ini:"kcp_bind_port" json:"kcp_bind_port"`
	ProxyBindPort int `ini:"proxy_bind_port" json:"proxy_bind_port"`
	VhostHTTPPort int `ini:"vhost_http_port" json:"vhost_http_port"`
	VhostHTTPSPort int `ini:"vhost_https_port" json:"vhost_https_port"`
	TCPMuxHTTPConnectPort int `ini:"tcpmux_httpconnect_port" json:"tcpmux_httpconnect_port" validate:"gte=0,lte=65535"`
	TCPMuxPassthrough bool `ini:"tcpmux_passthrough" json:"tcpmux_passthrough"`
	VhostHTTPTimeout int64 `ini:"vhost_http_timeout" json:"vhost_http_timeout"`
	DashboardAddr string `ini:"dashboard_addr" json:"dashboard_addr"`
	DashboardPort int `ini:"dashboard_port" json:"dashboard_port"`
	DashboardUser string `ini:"dashboard_user" json:"dashboard_user"`
	DashboardPwd string `ini:"dashboard_pwd" json:"dashboard_pwd"`
	EnablePrometheus bool `ini:"enable_prometheus" json:"enable_prometheus"`
	AssetsDir string `ini:"assets_dir" json:"assets_dir"`
	LogFile string `ini:"log_file" json:"log_file"`
	LogWay string `ini:"log_way" json:"log_way"`
	LogLevel string `ini:"log_level" json:"log_level"`
	LogMaxDays int64 `ini:"log_max_days" json:"log_max_days"`
	DisableLogColor bool `ini:"disable_log_color" json:"disable_log_color"`
	DetailedErrorsToClient bool `ini:"detailed_errors_to_client" json:"detailed_errors_to_client"`
	SubDomainHost string `ini:"subdomain_host" json:"subdomain_host"`
	TCPMux bool `ini:"tcp_mux" json:"tcp_mux"`
	TCPMuxKeepaliveInterval int64 `ini:"tcp_mux_keepalive_interval" json:"tcp_mux_keepalive_interval"`
	TCPKeepAlive int64 `ini:"tcp_keep_alive" json:"tcp_keep_alive"`
	Custom404Page string `ini:"custom_404_page" json:"custom_404_page"`
	AllowPorts map[int]struct{} `ini:"-" json:"-"`
	MaxPoolCount int64 `ini:"max_pool_count" json:"max_pool_count"`
	MaxPortsPerClient int64 `ini:"max_ports_per_client" json:"max_ports_per_client"`
	TLSOnly bool `ini:"tls_only" json:"tls_only"`
	TLSCertFile string `ini:"tls_cert_file" json:"tls_cert_file"`
	TLSKeyFile string `ini:"tls_key_file" json:"tls_key_file"`
	TLSTrustedCaFile string `ini:"tls_trusted_ca_file" json:"tls_trusted_ca_file"`
	HeartbeatTimeout int64 `ini:"heartbeat_timeout" json:"heartbeat_timeout"`
	UserConnTimeout int64 `ini:"user_conn_timeout" json:"user_conn_timeout"`
	HTTPPlugins map[string]plugin.HTTPPluginOptions `ini:"-" json:"http_plugins"`
	UDPPacketSize int64 `ini:"udp_packet_size" json:"udp_packet_size"`
	PprofEnable bool `ini:"pprof_enable" json:"pprof_enable"`




}
