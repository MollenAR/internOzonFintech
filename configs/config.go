package configs

import (
	"flag"

	"github.com/spf13/viper"
)

const BdTarantool = true

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

func ServerConfig() (string, bool, error) {
	bdType := flag.Bool("m", false, "использовать базу данных tarantool, иначе postgresql")
	flag.Parse()

	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return "", false, err
	}

	var port string
	switch {
	case *bdType == BdTarantool:
		port = viper.GetString("server.tarantool_port")
	default:
		port = viper.GetString("server.psql_port")
	}

	address := viper.GetString("server.address")
	return address + ":" + port, *bdType, nil
}

func TarantoolConfig() (string, string, string, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		return "", "", "", err
	}

	user := viper.GetString("tarantool.user")
	pass := viper.GetString("tarantool.pass")
	addr := viper.GetString("tarantool.addr")
	return user, pass, addr, nil
}
