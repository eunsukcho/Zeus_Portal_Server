package httpd

import (
	"fmt"
	"net/http"
	"zeus/models"

	"github.com/gin-gonic/gin"
)

type K8SNamespaceInterface interface {
	BindingModel(c *gin.Context)
	CreateRequestProject(c *gin.Context)
	CreateRequestResourceQuota(c *gin.Context)
	CreateServiceAccount(c *gin.Context)
	CreateRole(c *gin.Context)
	CreateRoleBinding(c *gin.Context)
}

func (h *Handler) BindingModel(c *gin.Context) {
	var requestData models.K8SRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Set("RequestData", requestData)
	c.Next()
}

func (h *Handler) CreateRequestProject(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)

	namespaceEndpoint := h.k8s.NamespaceEndpoint
	fmt.Println("Namespace Endpoint : ", namespaceEndpoint)

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var k8SProjcet models.K8SProjcet
	requestData.ApiVersion = "v1"
	requestData.Kind = "Namespace"
	k8SProjcet.SettingK8SPj(requestData)

	c.JSON(http.StatusOK, k8SProjcet)
}

func (h *Handler) CreateRequestResourceQuota(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)
	requestData.Name = requestData.Name + "-resource" // resourceQuota Name(UserId-resource)

	reousceEndpoint := h.k8s.NamespaceEndpoint + "/" + requestData.Namespace + "/resourcequotas"
	fmt.Println("Namespace_Resource_Endpoint : ", reousceEndpoint)

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var K8SProjcet models.K8SProjcet
	requestData.ApiVersion = "v1"
	requestData.Kind = "ResourceQuota"
	K8SProjcet.SettingK8SPj(requestData)

	var k8sResourceHardSpec models.Hard
	k8sResourceHardSpec.SettingSpecHard(requestData)
	requestData.Hard = k8sResourceHardSpec

	var k8sResourceSpec models.Spec
	k8sResourceSpec.SettingResourceSpec(requestData)
	requestData.Spec = k8sResourceSpec

	var k8SResource models.K8SResource
	k8SResource.SettingSpecResource(requestData, K8SProjcet)

	c.JSON(http.StatusOK, k8SResource)
}

func (h *Handler) CreateServiceAccount(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)

	serviceEndpoint := h.k8s.NamespaceEndpoint + "/" + requestData.Namespace + "/serviceaccounts"
	fmt.Println("Namespace_ServiceAccount_Endpoint : ", serviceEndpoint)

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var k8SProjcet models.K8SProjcet
	requestData.ApiVersion = "v1"
	requestData.Kind = "ServiceAccount"
	k8SProjcet.SettingK8SPj(requestData)

	c.JSON(http.StatusOK, k8SProjcet)
}

func (h *Handler) CreateRole(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)
	requestData.Name = requestData.Name + "-role" // Role Name(UserId-role)

	rolesEndpoint := h.k8s.NamespaceEndpoint + "/" + requestData.Namespace + "/roles"
	fmt.Println("Namespace_Role_Endpoint : ", rolesEndpoint)

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var k8SProjcet models.K8SProjcet
	requestData.ApiVersion = "rbac.authorization.k8s.io/v1"
	requestData.Kind = "Role"
	k8SProjcet.SettingK8SPj(requestData)

	var rulesArray models.RulesArray
	rulesArray.SettingValue()
	requestData.RulesArray = rulesArray

	var rules models.K8SRole
	rules.SettingK8SSetting(requestData, k8SProjcet)

	c.JSON(http.StatusOK, rules)
}

func (h *Handler) CreateRoleBinding(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)
	name := requestData.Name
	requestData.Name = name + "-roleBiding" // RoleBinding Name(UserId-role)

	rolesBindingEndpoint := h.k8s.NamespaceEndpoint + "/" + requestData.Namespace + "/rolebindings"
	fmt.Println("Namespace_RoleBinding_Endpoint : ", rolesBindingEndpoint)

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var k8SProjcet models.K8SProjcet
	requestData.ApiVersion = "rbac.authorization.k8s.io/v1"
	requestData.Kind = "RoleBinding"
	k8SProjcet.SettingK8SPj(requestData)

	var roleRef models.RoleBindingBaseObject
	requestData.Name = name + "-role" // RoleBinding Name(UserId-role)
	roleRef.SettingRoleRef(requestData)
	requestData.RoleRef = roleRef

	requestData.Name = name
	roleRef.SettingSubject(requestData)
	requestData.SettingRequest(roleRef)

	var roleBinding models.K8SRoleBinding
	roleBinding.SettingRoleBinding(requestData, k8SProjcet)

	c.JSON(http.StatusOK, roleBinding)
}
