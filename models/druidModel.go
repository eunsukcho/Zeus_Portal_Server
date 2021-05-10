package models

type LogSearchObj struct {
	Table          string `json:"table"`
	Container_name string `json:"container_name"`
	Log_Level      string `json:"loglevel"`
	Hostname       string `json:"hostname"`
	Namespace      string `json:"namespace"`
	Pod_name       string `json:"pod_name"`
	Log            string `json:"log"`
	/*StartDt        string              `json:"startDt"`
	EndDt          string              `json:"endDt"`*/
	Host    string              `json:"host"`
	Message string              `json:"message"`
	Process string              `json:"process"`
	DateArr []map[string]string `json:"dateArr"`
}

func (log LogSearchObj) GetValue(key string) {

}
