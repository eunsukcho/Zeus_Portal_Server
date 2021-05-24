package models

type K8SRequestData struct {
	ApiVersion string `json:"apiVersion"`
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	Namespace  string `json:"namespace"`
	Cpu        string `json:"cpu"`
	Memory     string `json:"memory"`
	UserID     string `json:"userId"`

	Hard       Hard                    `json:"hard"`
	Spec       Spec                    `json:"spec"`
	MetaData   MetaData                `json:"metadata"`
	RulesArray []RulesArray            `json:"rules"`
	Subjects   []RoleBindingBaseObject `json:"subjects"`
	RoleRef    RoleBindingBaseObject   `json:"roleRef"`
}

func (request *K8SRequestData) SettingRequest(ref RoleBindingBaseObject) {
	request.Subjects = append(request.Subjects, ref)
}

func (request *K8SRequestData) SettingRuleRequest(rules RulesArray) {
	request.RulesArray = append(request.RulesArray, rules)
}
