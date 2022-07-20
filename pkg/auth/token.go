package auth


type TokenConfig struct {
	 
	Token string `ini:"token" json:"token"`

}

func getDefaultTokenConf() TokenConfig{
	return TokenConfig{
		Token: "",
	}
}
