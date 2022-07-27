package msg

const (
	TypeLogin = 'o'
	TypeLoginResp = '1'
	TypeNewProxy = 'p'
	TypeNewProxyResp='2'
)

var (
	msgTypeMap = map[byte]interface{}{
		TypeLogin: Login{},
		TypeLoginResp: LoginResp{},
		TypeNewProxy: NewProxy{},
		TypeNewProxyResp:
	}
)


type NewProxy struct {
	ProxyName string `json:"proxy_name,omitempty"`
	ProxyType string `json:"proxy_type,omitempty"`
	UseEncryption bool `json:"use_encryption,omitempty"`
	UseCompression bool `json:"use_compression,omitempty"`
	Group string `json:"group,omitempty"`
	GroupKey string `json:"group_key,omitempty"`
	Metas map[string]string `json:"metas,omitempty"`

	//tcp and udp only
	RemotePort int `json:"remote_port,omitempty"`

	//http and https only
	CustomDomains []string `json:"custom_domains,omitempty"`
	SubDomain string `json:"sub_domain,omitempty"`
	Locations []string `json:"locations,omitempty"`
	HTTPUser string `json:"http_user,omitempty"`
	HTTPPwd string `json:"http_pwd,omitempty"`
	HostHeaderRewrite string `json:"host_header_rewrite,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	RouteByHTTPUser string `json:"route_by_http_user,omitempty"`

	//stcp
	Sk string `json:"sk,omitempty"`

	//tcpmux
	Multiplexer string `json:"multiplexer,omitempty"`

}



type Login struct {
	Version string  `json:"vsersion,omitempty"`
	Hostname string  `json:"hostname,omitempty"`
	Os string `json:"os,omitempty"`
	Arch string `json:"arch,omitempty"`
	User string `json:"user,omitempty"`
	PrivilegeKey string `json:"privilege_key,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
	RunID string `json:"run_id,omitempty"`
	Metas map[string]string `json:"metas,omitempty"`
	
	//Some global configures
	PoolCount int `json:"pool_count,omitempty"`
}

type LoginResp struct {
	Version string `json:"version,omitempty"`
	RunID string `json:"run_id,omitempty"`
	ServerUDPPort int `json:"server_udp_port,omitempty"`
	Error string `json:"error,omitempty"`
}


type Ping struct {
	PrivilegeKey string `json:"privilege_key,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}



type Pong struct {
	Error string `json:"error,omitempty"`
}



type NewWorkConn struct {
	RunID string `json:"run_id,omitempty"`
	PrivilegeKey string `json:"privilege_key,omitempty"`
	Timestamp int64 `json:"timestamp,omitempty"`
}

type NewProxyResp struct{
	ProxyName string `json:"proxy_name,omitempty"`
	RemoteAddr string `json:"remote_addr,omitempty"`
	Error string `json:"error,omitempty"`
}

type ReqWorkConn struct {
	
}

type CloseProxy struct {
	ProxyName string `json:"proxy_name,omitempty"`
}





