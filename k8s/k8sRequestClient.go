package k8s

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"zeus/models"
)

func (k8sInfo *K8SInfo) CreateProject(reqJson models.K8SProjcet) (string, int, error) {
	url := k8sInfo.NamespaceEndpoint
	fmt.Println("CreateProject : ", url)

	reqJsonTxt, err := json.Marshal(reqJson)
	if err != nil {
		return "", 0, err
	}
	respBody, statusCode, err := HTTPK8S("POST", url, bytes.NewBuffer([]byte(reqJsonTxt)), k8sInfo.Token, "")
	if err != nil {
		return "", 0, err
	}
	fmt.Println("Response : ", string(respBody))
	fmt.Println("StatusCode :", statusCode)
	return string(respBody), statusCode, nil
}

func (k8sInfo *K8SInfo) CreateResource(requestData models.K8SRequestData, reqJson models.K8SResource) (string, int, error) {
	url := k8sInfo.NamespaceEndpoint + "/" + requestData.Namespace + "/resourcequotas"
	fmt.Println("CreateResource : ", url)

	reqJsonTxt, err := json.Marshal(reqJson)
	if err != nil {
		return "", 0, err
	}
	respBody, statusCode, err := HTTPK8S("POST", url, bytes.NewBuffer([]byte(reqJsonTxt)), k8sInfo.Token, "")
	if err != nil {
		return "", 0, err
	}
	fmt.Println("Response : ", string(respBody))
	fmt.Println("StatusCode :", statusCode)
	return string(respBody), statusCode, nil
}

func (k8sInfo *K8SInfo) CreateServiceAccount(requestData models.K8SRequestData, reqJson models.K8SProjcet) (string, int, error) {
	url := k8sInfo.NamespaceEndpoint + "/" + requestData.Namespace + "/serviceaccounts"
	fmt.Println("CreateServiceAccount : ", url)

	reqJsonTxt, err := json.Marshal(reqJson)
	if err != nil {
		return "", 0, err
	}
	respBody, statusCode, err := HTTPK8S("POST", url, bytes.NewBuffer([]byte(reqJsonTxt)), k8sInfo.Token, "")
	if err != nil {
		return "", 0, err
	}
	fmt.Println("Response : ", string(respBody))
	fmt.Println("StatusCode :", statusCode)
	return string(respBody), statusCode, nil

}

func (k8sInfo *K8SInfo) CreateRole(requestData models.K8SRequestData, reqJson models.K8SRole) (string, int, error) {
	url := k8sInfo.AuthEndpoint + "/" + requestData.Namespace + "/roles"
	fmt.Println("CreateRole : ", url)

	reqJsonTxt, err := json.Marshal(reqJson)
	if err != nil {
		return "", 0, err
	}
	respBody, statusCode, err := HTTPK8S("POST", url, bytes.NewBuffer([]byte(reqJsonTxt)), k8sInfo.Token, "")
	if err != nil {
		return "", 0, err
	}
	fmt.Println("Response : ", string(respBody))
	fmt.Println("StatusCode :", statusCode)
	return string(respBody), statusCode, nil
}

func (k8sInfo *K8SInfo) CreateRoleBinding(requestData models.K8SRequestData, reqJson models.K8SRoleBinding) (string, int, error) {
	url := k8sInfo.AuthEndpoint + "/" + requestData.Namespace + "/rolebindings"
	fmt.Println("CreateRoleBinding : ", url)

	reqJsonTxt, err := json.Marshal(reqJson)
	if err != nil {
		return "", 0, err
	}
	respBody, statusCode, err := HTTPK8S("POST", url, bytes.NewBuffer([]byte(reqJsonTxt)), k8sInfo.Token, "")
	if err != nil {
		return "", 0, err
	}
	fmt.Println("Response : ", string(respBody))
	fmt.Println("StatusCode :", statusCode)
	return string(respBody), statusCode, nil
}

func (k8sInfo *K8SInfo) DeleteNamespace(namespace string) (string, int, error) {
	url := k8sInfo.NamespaceEndpoint + "/" + namespace
	fmt.Println("CreateProject : ", url)

	respBody, statusCode, err := HTTPK8S("DELETE", url, nil, k8sInfo.Token, "")
	if err != nil {
		return "", 0, err
	}
	fmt.Println("Delete Response : ", string(respBody))
	fmt.Println("Delete StatusCode :", statusCode)
	return string(respBody), statusCode, nil
}

func (k8sInfo *K8SInfo) GetUserSecretName(requestData models.K8SRequestData) (string, int, error) {
	url := k8sInfo.NamespaceEndpoint + "/" + requestData.Namespace + "/serviceaccounts/" + requestData.Name
	fmt.Println("GetUsserSecretName : ", url)
	respBody, statusCode, err := HTTPK8S("GET", url, nil, k8sInfo.Token, "user-secret")
	if err != nil {
		return "", 0, err
	}

	var secret map[string]interface{}
	if err = json.Unmarshal(respBody, &secret); err != nil {
		return "", 0, err
	}
	val := secret["secrets"]

	/***********/

	tmpValue, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}

	/***********/
	var secretValue []map[string]interface{}
	if err = json.Unmarshal(tmpValue, &secretValue); err != nil {
		fmt.Println(err.Error())

		return "", 0, err
	}
	userSecreat := secretValue[0]["name"]
	fmt.Println("userSecret : ", userSecreat)
	/***********/

	return userSecreat.(string), statusCode, nil
}

func (k8sInfo *K8SInfo) GetUserToken(requestData models.K8SRequestData) (string, int, error) {
	url := k8sInfo.NamespaceEndpoint + "/" + requestData.Namespace + "/secrets/" + requestData.Name
	fmt.Println("GetUsserSecretName : ", url)
	respBody, statusCode, err := HTTPK8S("GET", url, nil, k8sInfo.Token, "user-secret")
	if err != nil {
		return "", 0, err
	}

	var userToken map[string]interface{}
	if err = json.Unmarshal(respBody, &userToken); err != nil {
		return "", 0, err
	}
	val := userToken["data"]

	/***********/
	tmpValue, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	/***********/

	var userTokenData map[string]interface{}
	if err = json.Unmarshal(tmpValue, &userTokenData); err != nil {
		fmt.Println(err.Error())

		return "", 0, err
	}
	token := userTokenData["token"]
	fmt.Println("userTokenData : ", token)
	/***********/

	return token.(string), statusCode, nil
}

func HTTPK8S(method string, url string, reqJson *bytes.Buffer, token string, secret string) (respBody []byte, statusCode int, err error) {
	var req *http.Request

	if method == "DELETE" || secret == "user-secret" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, reqJson)
	}

	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", token)

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if secret != "" && secret == "user-secret" {
		req.Header.Add("Content-Type", "text")
	}

	if err != nil {
		return nil, http.StatusRequestTimeout, err
	}

	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, resp.StatusCode, err
	}

	return respBody, resp.StatusCode, nil
}
