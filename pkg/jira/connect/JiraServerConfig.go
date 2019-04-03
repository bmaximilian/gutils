package connect

type JiraServerConfig struct {
	Url        string
	APIVersion int
	TlsConfig  TLSConfig
	Token      string
	Tempo      bool
	UserName   string
	Password   string
}
