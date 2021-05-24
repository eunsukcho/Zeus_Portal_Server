package k8s

import (
	"fmt"
	"os"
	"zeus/models"

	"github.com/spf13/viper"
)

type K8SInterface interface {
	CreateProject(models.K8SProjcet) (string, int, error)
	CreateResource(models.K8SRequestData, models.K8SResource) (string, int, error)
	CreateServiceAccount(models.K8SRequestData, models.K8SProjcet) (string, int, error)
	CreateRole(models.K8SRequestData, models.K8SRole) (string, int, error)
	CreateRoleBinding(models.K8SRequestData, models.K8SRoleBinding) (string, int, error)

	GetUserSecretName(models.K8SRequestData) (string, int, error)
	GetUserToken(models.K8SRequestData) (string, int, error)

	DeleteNamespace(string) (string, int, error)
}

type K8SInfo struct {
	NamespaceEndpoint string
	AuthEndpoint      string
	Token             string
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
	auth := viper.GetString("k8s." + env + ".k8sAuthEndpoint")
	token := viper.GetString("k8s." + env + ".k8sToken")

	k8sInfo := K8SInfo{
		NamespaceEndpoint: namespaceEndpoint,
		AuthEndpoint:      auth,
		Token:             token,
	}
	return &k8sInfo
}
