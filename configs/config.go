package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	JWTSecretKey string `mapstructure:"jwt_secret_key"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Println("Config loaded successfully")
}
