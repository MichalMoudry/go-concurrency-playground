package config

import "github.com/spf13/viper"

type Config struct {
	Port         int
	DbConnStr    string
	RedisConnStr string
}

func LoadConfigFromFile(path string) (Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	return Config{
		Port:         viper.GetInt("port"),
		DbConnStr:    viper.GetString("db_conn_str"),
		RedisConnStr: viper.GetString("redis"),
	}, nil
}
