package server

type HTTPPluginOptions struct {
	Name string `ini:"name"`
	Addr string `ini:"addr"`
	Path string `ini:"path"`
	Ops []string `ini:"ops"`
	TLSVerify bool `ini:"tls_verify"`
}
