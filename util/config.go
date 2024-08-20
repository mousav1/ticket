package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	APPPORT      string `mapstructure:"APP_PORT"`
	APPNAME      string `mapstructure:"APP_NAME"`
	APPENV       string `mapstructure:"APP_ENV"`
	APPDEBUG     string `mapstructure:"APP_DEBUG"`
	DBCONNECTION string `mapstructure:"DB_CONNECTION"`
	DBHOST       string `mapstructure:"DB_HOST"`
	DBPORT       string `mapstructure:"DB_PORT"`
	DBDATABASE   string `mapstructure:"DB_DATABASE"`
	DBUSERNAME   string `mapstructure:"DB_USERNAME"`
	DBPASSWORD   string `mapstructure:"DB_PASSWORD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
