package env

import (
	"fmt"
	"github.com/spf13/viper"
)

type DatabaseConfiguration struct {
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbName     string `mapstructure:"DB_NAME"`
}

type Configuration struct {
	AppEnv   string                `mapstructure:"APP_ENV"`
	Database DatabaseConfiguration `mapstructure:",squash"`
}

var configuration Configuration

func init() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		viper.SetConfigFile("../../../../../.env")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into map, %v", err)
	}

}

func GetConfiguration() Configuration {
	return configuration
}
