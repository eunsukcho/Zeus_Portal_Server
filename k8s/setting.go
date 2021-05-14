package k8s

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type K8SInfo struct {
	NamespaceEndpoint string
}

func SettingK8SInfo(env string) *K8SInfo {
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
	namespaceEndpoint := viper.GetString("k8s." + env + ".k8sNamespaceEndpoint")

	k8sInfo := K8SInfo{
		NamespaceEndpoint: namespaceEndpoint,
	}
	return &k8sInfo
}

func (k8s K8SInfo) GetK8SInfo() *K8SInfo {
	return &k8s
}
