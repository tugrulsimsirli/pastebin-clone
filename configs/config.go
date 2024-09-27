package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type DBConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Port     string `mapstructure:"port"`
	SSLMode  string `mapstructure:"sslmode"`
}

type Config struct {
	JWTSecretKey string   `mapstructure:"jwt_secret_key"`
	DBConfig     DBConfig `mapstructure:"db"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/root")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// JWT Secret Key'i .env'den alıyoruz
	AppConfig.JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	if AppConfig.JWTSecretKey == "" {
		log.Fatalf("JWT_SECRET_KEY is not set in environment variables")
	}
}
