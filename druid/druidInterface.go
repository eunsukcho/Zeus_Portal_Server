package druid

import (
	"fmt"
	"os"
	"zeus/models"

	"github.com/spf13/viper"
)

type DruidInterface interface {
	GetColumnValue(string, string, string, chan string, chan []map[string]string) ([]map[string]string, error)
	GetLogValue(models.LogSearchObj, string, string) ([]map[string]string, error)
}

type ClientInfo struct {
	Host     string
	Port     string
	Endpoint string
}

func NewClientInfo() *ClientInfo {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	host := viper.GetString("druid.ope.host")
	port := viper.GetString("druid.ope.port")
	endpoint := viper.GetString("druid.ope.endpoint")

	clientInfo := ClientInfo{
		Host:     host,
		Port:     port,
		Endpoint: endpoint,
	}
	return &clientInfo
}
