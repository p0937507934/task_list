package config

import "github.com/spf13/viper"

var GlobalConfig ServerConfig

func InitConfig() {
	v := viper.New()
	v.SetDefault("SERVER_PORT", "8000")
	v.SetDefault("DB_USERNAME", "root")
	v.SetDefault("DB_PASSWORD", "secret")
	v.SetDefault("DB_PORT", "3306")
	v.SetDefault("DB_NAME", "task_list")
	v.SetDefault("DB_HOST", "127.0.0.1")
	v.AutomaticEnv()
	if err := v.Unmarshal(&GlobalConfig); err != nil {
		panic(err)
	}

}
