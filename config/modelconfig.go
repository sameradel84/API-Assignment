package config

type Config struct {
	Db       string `json:"db"`
	Server   string `json:"server"`
	User     string `json:"user"`
	Password string `json:"password"`
}
