package configs

import (
	"github.com/spf13/viper"
)

func PostgresConfig() (string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	user := viper.GetString("postgres.user")
	dbname := viper.GetString("postgres.dbname")
	password := viper.GetString("postgres.password")
	host := viper.GetString("postgres.host")
	sslmode := viper.GetString("postgres.sslmode")

	return "user=" + user + " dbname=" + dbname + " password=" + password + " host=" + host + " sslmode=" + sslmode, nil
}

func AddresConfig() (string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return "", err
	}

	address := viper.GetString("server.address")
	port := viper.GetString("server.port")
	return address + ":" + port, nil
}