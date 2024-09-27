package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	JWTSecretKey string `mapstructure:"jwt_secret_key"`
}

var AppConfig Config

// Config dosyasını yükleyen fonksiyon
func LoadConfig() {
	viper.SetConfigName("config") // config dosyasının adı
	viper.SetConfigType("yaml")   // config dosyasının türü
	viper.AddConfigPath(".")      // config dosyasının yolunu belirle (. = kök dizin)

	// Config dosyasını okuma
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Config dosyasını yapılandırmaya map etme
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Println("Config loaded successfully")
}
