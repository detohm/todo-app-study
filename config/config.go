package config

import (
	"os"
)

type Config struct {
	MySQLConn string
	Port      string
}

func LoadConfig() Config {
	config := Config{}
	config.MySQLConn = os.Getenv("MYSQL_CONN")
	config.Port = os.Getenv("PORT")
	return config
}
