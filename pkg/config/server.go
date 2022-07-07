package config

type ServerCommonConf struct {

	BindAddr string `ini:"bind_addr" json:"bind_addr"`
	BindPort int `ini:"bind_port" json:"bind_port"`
	BindUDPPort int `ini:"bind_udp_port" json:"bind_udp_port"`
	KCPBindPort int `ini:"kcp_bind_port" json:"kcp_bind_port"`
	ProxyBindPort int `ini:"proxy_bind_port" json:"proxy_bind_port"`
	VhostHTTPPort int `ini:"vhost_http_port" json:"vhost_http_port"`
	VhostHTTPSPort int `ini:"vhost_https_port" json:"vhost_https_port"`
	TCPMuxHTTPConnectPort int `ini:"tcpmux_httpconnect_port" json:"tcpmux_httpconnect_port" validate:"gte=0,lte=65535"`


}
