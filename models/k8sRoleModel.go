package models

type K8SRole struct {
	K8SProjcet
	Rules RulesArray `json:"rules"`
}
type RulesArray struct {
	ApiGroups []string `json:"apiGroups"`
	Resources []string `json:"resources"`
	Verbs     []string `json:"verbs"`
}

func (rule *K8SRole) SettingK8SSetting(request K8SRequestData, baseProjectInfo K8SProjcet) {
	rule.K8SProjcet = baseProjectInfo
	rule.Rules = request.RulesArray
}

func (rule *RulesArray) SettingValue() {
	apiGroups := []string{"", "apps"}
	resources := []string{"pods", "service", "statefulset", "deployments"}
	verbs := []string{"get", "list", "edit", "create", "update", "patch", "delete"}

	rule.ApiGroups = apiGroups
	rule.Resources = resources
	rule.Verbs = verbs
}
