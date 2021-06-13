package util

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	ServerGrpcAddr    string `mapstructure:"SERVER_GRPC_ADDR"`
	ServerApiAddr     string `mapstructure:"SERVER_API_ADDR"`
	ServerGraphqlAddr string `mapstructure:"SERVER_GRAPHQL_ADDR"`
	DbHost            string `mapstructure:"DB_HOST"`
	DbPort            string `mapstructure:"DB_PORT"`
	DbUsername        string `mapstructure:"DB_USERNAME"`
	DbPassword        string `mapstructure:"DB_PASSWORD"`
	DbDatabase        string `mapstructure:"DB_DATABASE"`
	DbSslmode         string `mapstructure:"DB_SSLMODE"`
}

// Returns the db connection url
func (config *Config) GetDbConnectionUrl() string {
	return fmt.Sprintf("postgresql://%v:%v@%v:%v/%v?sslmode=%v",
		config.DbUsername,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbDatabase,
		config.DbSslmode)
}

func LoadConfig(path string) (Config, error) {
	var c Config
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return Config{}, err
	}
	viper.AutomaticEnv()
	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return Config{}, err
	}
	return c, nil
}
