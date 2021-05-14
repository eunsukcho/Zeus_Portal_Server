package models

type K8SRoleBinding struct {
	K8SProjcet
	Subjects []RoleBindingBaseObject `json:"subjects"`
	RoleRef  RoleBindingBaseObject   `json:"roleRef"`
}
type RoleBindingBaseObject struct {
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	ApiGroup string `json:"apiGroup"`
}

func (roleBinding *K8SRoleBinding) SettingRoleBinding(request K8SRequestData, baseProjectInfo K8SProjcet) {
	roleBinding.K8SProjcet = baseProjectInfo
	roleBinding.Subjects = request.Subjects
	roleBinding.RoleRef = request.RoleRef
}

func (baseObj *RoleBindingBaseObject) SettingSubject(request K8SRequestData) {
	baseObj.Kind = "ServiceAccount"
	baseObj.ApiGroup = ""
	baseObj.Name = request.Name
}
func (baseObj *RoleBindingBaseObject) SettingRoleRef(request K8SRequestData) {
	baseObj.Kind = "Role"
	baseObj.ApiGroup = "rbac.authorization.k8s.io"
	baseObj.Name = request.Name
}
