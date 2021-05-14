package models

type K8SResource struct {
	K8SProjcet
	Spec `json:"spec"`
}
type Spec struct {
	Hard Hard `json:"hard"`
}

type Hard struct {
	Cpu    string `json:"limits.cpu"`
	Memory string `json:"limits.memory"`
}

func (resource *K8SResource) SettingSpecResource(request K8SRequestData, baseProjectInfo K8SProjcet) {
	resource.K8SProjcet = baseProjectInfo
	resource.Spec = request.Spec
}

func (spec *Spec) SettingResourceSpec(request K8SRequestData) {
	spec.Hard = request.Hard
}
func (hardSpec *Hard) SettingSpecHard(request K8SRequestData) {
	hardSpec.Cpu = request.Cpu
	hardSpec.Memory = request.Memory
}
