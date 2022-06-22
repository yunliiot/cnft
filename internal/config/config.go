package config

var Config AppConfig

type AppConfig struct {
	AppConfigServer `json:"server"`
}

type AppConfigServer struct {
	UseChain33 bool   `json:"useChain33"`
	Chain33Url string `json:"chain33Url"`
}
