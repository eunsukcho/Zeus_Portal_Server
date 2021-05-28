package dblayer

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type dbInfo struct {
	db_name     string
	db_port     string
	db_host     string
	db_username string
	db_password string
}

func settingDB() *dbInfo {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	db_name := viper.GetString("app.db_name")
	db_port := viper.GetString("app.db_port")
	db_host := viper.GetString("app.db_host")
	db_username := viper.GetString("app.db_username")
	db_password := viper.GetString("app.db_password")

	fmt.Println("settingDB Path Init")

	dbinfo := dbInfo{
		db_name:     db_name,
		db_port:     db_port,
		db_host:     db_host,
		db_username: db_username,
		db_password: db_password,
	}
	return &dbinfo
}
