package connect

type TempoOptions struct {
	Enabled    bool
	ApiVersion int
}

type JiraServerConfig struct {
	Url        string
	APIVersion int
	TlsConfig  *TLSConfig
	Token      string
	Tempo      *TempoOptions
	UserName   string
	Password   string
}
