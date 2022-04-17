package config

type ServerConfig struct {
	ServerPort  int `mapstructure:"SERVER_PORT"`
	MySQLConfig `mapstructure:",squash"`
}

type MySQLConfig struct {
	Username string `mapstructure:"DB_USERNAME"`
	Password string `mapstructure:"DB_PASSWORD"`
	DbPORT   int    `mapstructure:"DB_PORT"`
	DbName   string `mapstructure:"DB_NAME"`
	DbHost   string `mapstructure:"DB_HOST"`
}
