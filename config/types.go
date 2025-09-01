package config

type Config struct {
	DBUser       string `config:"DB_USER"`
	DBPassword   string `config:"DB_PASSWORD"`
	DBName       string `config:"DB_NAME"`
	SSLMode      string `config:"DB_SSLMODE"`
	DBPort       string `config:"DB_PORT"`
	Port         string `config:"PORT"`
	Host         string `config:"HOST"`
	Protocol     string `config:"PROTOCOL"`
	LPort        string `config:"LPORT"`
	ProductsPort string `config:"PRODUCTS_PORT"`
}
