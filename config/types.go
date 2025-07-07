package config

type Config struct {
	User     string `config:"USER"`
	Password string `config:"PASSWORD"`
	DBName   string `config:"DBNAME"`
	SSLMode  string `config:"SSLMODE"`
	Port     string `config:"PORT"`
	Host     string `config:"HOST"`
}
