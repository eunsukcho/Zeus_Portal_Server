package requestLayer

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type KeycloakApiInfo struct {
	UserEndpoint 	string
	GroupEndpoint	string
}
func SettingKeycloakInfo(env string) *KeycloakApiInfo {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	userEndpoint := viper.GetString("keycloak."+env+".userEndpoint")
	groupEndpoint := viper.GetString("keycloak."+env+".groupEndpoint")

	keycloakApiInfo := KeycloakApiInfo{
		UserEndpoint:     userEndpoint,
		GroupEndpoint:     groupEndpoint,
	}
	return &keycloakApiInfo
}

func (keyclock KeycloakApiInfo) getKeyclock() *KeycloakApiInfo {
	return &keyclock
}