package utils

import "github.com/spf13/viper"

type EnvironmentVariables struct {
	RABBITMQ_URL        string `mapstructure:"RABBITMQ_URL"` // connection string to rabbitmq, have host, port (or url), username, password, vhost (default vhost is '/')
	HTTP_SERVER_ADDRESS string `mapstructure:"HTTP_SERVER_ADDRESS"`
	MODE_ENV            string `mapstructure:"MODE"`
}

func LoadEnvironmentVariables(pathname string) (vars EnvironmentVariables, err error) {
	viper.AddConfigPath(pathname)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&vars)
	return
}
