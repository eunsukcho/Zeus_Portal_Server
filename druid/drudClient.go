package druid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

func ConvertColumnNameAs(column string) (returnAs string) {
	switch column {
	case "kubernetes.host":
		returnAs = "hostname"
	case "kubernetes.namespace_name":
		returnAs = "namespace"
	case "kubernetes.pod_name":
		returnAs = "podname"
	case "kubernetes.container_name":
		returnAs = "container_id"
	}
	return returnAs
}
func ConvertValueName(column string) (returnAs string) {
	switch column {
	case "hostname":
		returnAs = "kubernetes.host"
	case "namespace":
		returnAs = "kubernetes.namespace_name"
	case "podname":
		returnAs = "kubernetes.pod_name"
	case "container_id":
		returnAs = "kubernetes.container_name"
	case "startDt":
		returnAs = `\"__time\" > TIMESTAMP `
	case "endDt":
		returnAs = `\"__time\" <= TIMESTAMP `
	}
	return returnAs
}

func (urlInfo *ClientInfo) GetColumnValue(column string, table string, key chan string, c chan []map[string]string) (returnColumn []map[string]string, error error) {

	url := urlInfo.Host + ":" + urlInfo.Port + urlInfo.Endpoint

	var query bytes.Buffer
	query.WriteString(`SELECT \"`)
	query.WriteString(column)
	query.WriteString(`\" FROM \"`)
	query.WriteString(table)
	query.WriteString(`\" GROUP BY \"`)
	query.WriteString(column)
	query.WriteString(`\"`)

	reqTxt := `{"query" : "` + query.String() + `"}`

	var sqlJson = bytes.NewBuffer([]byte(reqTxt))

	respBody, err := HTTPDruid(url, sqlJson)
	if err == nil {
		error = json.Unmarshal(respBody, &returnColumn)
		if error != nil {
			fmt.Println(error)
			return nil, error
		}

		key <- column
		c <- returnColumn
		return returnColumn, nil
	}
	return nil, nil
}

func (urlInfo *ClientInfo) GetLogValue(where map[string]string, table string) (rst interface{}, error error) {
	url := urlInfo.Host + ":" + urlInfo.Port + urlInfo.Endpoint

	var whereQuery bytes.Buffer
	whereQuery.WriteString(` WHERE \"hostname\" IS NOT NULL`)

	keys := reflect.ValueOf(where).MapKeys()
	for _, key := range keys {
		convertKey := ConvertValueName(key.String())
		value := where[key.String()]

		if value != "" && (key.String() != "startDt" || key.String() != "endDt") {
			if key.String() == "startDt" || key.String() == "endDt" {
				fmt.Println("Date Column")
				whereQuery.WriteString(` AND `)
				whereQuery.WriteString(convertKey)
				whereQuery.WriteString("'")
				whereQuery.WriteString(value)
				whereQuery.WriteString("'")
				continue
			}
			whereQuery.WriteString(` AND `)
			whereQuery.WriteString(`\"`)
			whereQuery.WriteString(key.String())
			whereQuery.WriteString(`\"`)
			whereQuery.WriteString(`=`)
			whereQuery.WriteString("'")
			whereQuery.WriteString(value)
			whereQuery.WriteString("'")
		}
	}
	var query bytes.Buffer
	str := `SELECT \"__time\" AS collectDt, \"container_name\", \"hostname\", \"namespace_name\", \"pod_name\", \"loglevel\", \"log\" AS logMessage `
	query.WriteString(str)
	query.WriteString(`FROM \"`)
	query.WriteString(table)
	query.WriteString(`\" `)
	query.WriteString(whereQuery.String())
	query.WriteString(` LIMIT 100`)

	reqTxt := `{"query" : "` + query.String() + `"}`
	fmt.Println("Request JSON : ", reqTxt)
	var sqlJson = bytes.NewBuffer([]byte(reqTxt))

	respBody, err := HTTPDruid(url, sqlJson)
	var m interface{}
	if err == nil {
		error = json.Unmarshal(respBody, &m)
		if error != nil {
			fmt.Println(error)
			return nil, error
		}

		return m, nil
	}
	return nil, nil

}

func HTTPDruid(url string, sqlJson *bytes.Buffer) (respBody []byte, err error) {

	req, err := http.NewRequest("POST", url, sqlJson)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err = ioutil.ReadAll(resp.Body)
	return respBody, nil
}
