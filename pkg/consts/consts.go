package consts

var (
	Idle    string = "idle"
	Working string = "working"
	Closed  string = "closed"
	Online  string = "online"
	Offline string = "offline"

	TCPProxy    string = "tcp"
	UDPProxy    string = "udp"
	TCPMuxProxy string = "tcpmux"
	HTTPProxy   string = "http"
	HTTPSProxy  string = "https"
	STCPProxy   string = "stcp"
	XTCPProxy   string = "xtcp"
	SUDPProxy   string = "sudp"

	TokenAuthMethod string = "token"
	OidcAuthMethod  string = "oidc"

	HTTPConnectTCPMultiplexer string = "httpconnect"
)
