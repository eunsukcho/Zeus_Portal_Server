package models

type K8SProjcet struct {
	ApiVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	MetaData   MetaData `json:"metadata"`
}
type MetaData struct {
	Name      string `json:"name"` // 네임스페이스명
	Namespace string `json:"namespace,omitempty"`
}

func (k8sbase *K8SProjcet) SettingK8SPj(request K8SRequestData) {
	k8sbase.ApiVersion = request.ApiVersion
	k8sbase.Kind = request.Kind
	k8sbase.MetaData = request.MetaData
}
func (metadata *MetaData) SettingMetaData(request K8SRequestData) {
	metadata.Name = request.Name
	metadata.Namespace = request.Namespace
}
