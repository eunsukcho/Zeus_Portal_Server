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

	GetUserToken(c *gin.Context)

	DeleteNamespace(c *gin.Context)
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

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var k8SProjcet models.K8SProjcet
	requestData.ApiVersion = "v1"
	requestData.Kind = "Namespace"
	k8SProjcet.SettingK8SPj(requestData)

	rstTxt, statusCode, err := h.k8s.CreateProject(k8SProjcet)
	fmt.Println("CreateRequestProject Error :", err)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":        statusCode,
			"error message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"status": statusCode,
		"data":   rstTxt,
	})
}

func (h *Handler) CreateRequestResourceQuota(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)
	requestData.Name = requestData.Name + "-resource" // resourceQuota Name(UserId-resource)

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

	rstTxt, statusCode, err := h.k8s.CreateResource(requestData, k8SResource)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":        statusCode,
			"error message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"status": statusCode,
		"data":   rstTxt,
	})
}

func (h *Handler) CreateServiceAccount(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var k8SProjcet models.K8SProjcet
	requestData.ApiVersion = "v1"
	requestData.Kind = "ServiceAccount"
	k8SProjcet.SettingK8SPj(requestData)

	rstTxt, statusCode, err := h.k8s.CreateServiceAccount(requestData, k8SProjcet)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":        statusCode,
			"error message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"status": statusCode,
		"data":   rstTxt,
	})
}

func (h *Handler) CreateRole(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)
	requestData.Name = requestData.Name + "-role" // Role Name(UserId-role)

	var k8SMetaData models.MetaData
	k8SMetaData.SettingMetaData(requestData)
	requestData.MetaData = k8SMetaData

	var k8SProjcet models.K8SProjcet
	requestData.ApiVersion = "rbac.authorization.k8s.io/v1"
	requestData.Kind = "Role"
	k8SProjcet.SettingK8SPj(requestData)

	var rulesArray models.RulesArray
	rulesArray.SettingValue()
	requestData.SettingRuleRequest(rulesArray)

	var rules models.K8SRole
	rules.SettingK8SSetting(requestData, k8SProjcet)

	rstTxt, statusCode, err := h.k8s.CreateRole(requestData, rules)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":        statusCode,
			"error message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"status": statusCode,
		"data":   rstTxt,
	})
}

func (h *Handler) CreateRoleBinding(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)
	name := requestData.Name
	requestData.Name = name + "-roleBiding" // RoleBinding Name(UserId-role)

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

	rstTxt, statusCode, err := h.k8s.CreateRoleBinding(requestData, roleBinding)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":        statusCode,
			"error message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"status": statusCode,
		"data":   rstTxt,
	})
}
func (h *Handler) GetUserToken(c *gin.Context) {
	requestData := c.MustGet("RequestData").(models.K8SRequestData)
	rstTxt, statusCode, err := h.k8s.GetUserSecretName(requestData)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":        statusCode,
			"error message": err.Error(),
		})
		return
	}
	fmt.Println("User Secret Name :", rstTxt)
	requestData.Name = rstTxt //유저 시크릿명

	rstTxtToken, statusCodeToken, errToken := h.k8s.GetUserToken(requestData)
	if errToken != nil {
		c.AbortWithStatusJSON(statusCodeToken, gin.H{
			"status":        statusCodeToken,
			"error message": errToken.Error(),
		})
		return
	}
	fmt.Println("User Token Value : ", rstTxtToken)
	c.JSON(statusCodeToken, gin.H{
		"status": statusCodeToken,
		"data":   rstTxtToken,
	})
}

type UriParameter struct {
	Namespace string `uri:"namespace"`
}

func (h *Handler) DeleteNamespace(c *gin.Context) {
	var namespace UriParameter
	if err := c.ShouldBindUri(&namespace); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", namespace.Namespace)

	rstTxt, statusCode, err := h.k8s.DeleteNamespace(namespace.Namespace)
	if err != nil {
		c.AbortWithStatusJSON(statusCode, gin.H{
			"status":        statusCode,
			"error message": err.Error(),
		})
		return
	}
	c.JSON(statusCode, gin.H{
		"status": statusCode,
		"data":   rstTxt,
	})
}
