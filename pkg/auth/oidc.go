package auth



type OidcClientConfig struct {
	OidcClientID string `ini:"oidc_client_id" json:"oidc_client_id"`
	OidcClientSecret string `ini:"oidc_client_secret" json:"oidc_client_secret"`
	OidcAudience string `ini:"oidc_audience" json:"oidc_audience"`
	OidcTokenEndpointURL string `ini:"oidc_token_endpoint_url" json:"oidc_token_endpoint_url"`

	OidcAdditionalEndpointParams map[string]string `ini:"-" json:"oidc_additional_endpoint_params"`

}


func getDefaultOidcClientConf() OidcClientConfig{
	return OidcClientConfig{
		OidcClientID: "",
		OidcClientSecret: "",
		OidcAudience: "",
		OidcTokenEndpointURL: "",
		OidcAdditionalEndpointParams: make(map[string]string),
	}
}