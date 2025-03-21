package config
import (
	"github.com/spf13/viper"
)

type Config struct{
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDB       string `mapstructure:"POSTGRES_DB"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
}

func ProvideConfig() *Config {
	config := &Config{}
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic("Error reading config file "+err.Error())
	}

	if err := viper.Unmarshal(config); err != nil {
		panic("Unable to decode into struct "+err.Error())
	}

	return config
}