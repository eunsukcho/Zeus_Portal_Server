package druid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
	"zeus/models"
)

func (urlInfo *ClientInfo) GetColumnValue(column string, table string, tableDiv string, key chan string, c chan []map[string]string) (returnColumn []map[string]string, error error) {

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

func metaData(anything interface{}, tableDiv string) string {

	var whereQuery bytes.Buffer
	if tableDiv == "container" {
		whereQuery.WriteString(` WHERE \"hostname\" IS NOT NULL`)
	}
	if tableDiv == "syslog" {
		whereQuery.WriteString(` WHERE \"host\" IS NOT NULL`)
	}

	target := reflect.ValueOf(anything)
	elements := target.Elem()

	fmt.Printf("Type: %s\n", target.Type())

	for i := 0; i < elements.NumField(); i++ {
		mValue := elements.Field(i)
		mType := elements.Type().Field(i)
		tag := mType.Tag

		if mType.Name == "Table" {
			continue
		}

		value := fmt.Sprintf("%v", mValue.Interface())
		if value != "" {
			if mType.Name == "DateArr" {
				dateMap := mValue.Interface().([]map[string]string)

				startDt := dateMap[0]["startDt"]
				endDt := dateMap[1]["endDt"]

				if startDt != "" && endDt != "" {
					whereQuery.WriteString(` AND \"__time\" BETWEEN TIMESTAMP '`)
					whereQuery.WriteString(startDt)
					whereQuery.WriteString("' AND TIMESTAMP '")
					whereQuery.WriteString(endDt)
					whereQuery.WriteString("'")
				}
				if startDt != "" && endDt == "" {
					whereQuery.WriteString(` AND \"__time\" >= TIMESTAMP '`)
					whereQuery.WriteString(startDt)
					whereQuery.WriteString("'")
				}
				if startDt == "" && endDt != "" {
					whereQuery.WriteString(` AND \"__time\" <= TIMESTAMP '`)
					whereQuery.WriteString(endDt)
					whereQuery.WriteString("'")
				}

				continue
			}
			if mType.Name == "Log" || mType.Name == "Process" || mType.Name == "Message" {
				whereQuery.WriteString(` AND `)
				whereQuery.WriteString(`\"`)
				whereQuery.WriteString(tag.Get("json"))
				whereQuery.WriteString(`\"`)
				whereQuery.WriteString(` LIKE `)
				whereQuery.WriteString("'%")
				whereQuery.WriteString(value)
				whereQuery.WriteString("%'")
				continue
			}
			whereQuery.WriteString(` AND `)
			whereQuery.WriteString(`\"`)
			whereQuery.WriteString(tag.Get("json"))
			whereQuery.WriteString(`\"`)
			whereQuery.WriteString(`=`)
			whereQuery.WriteString("'")
			whereQuery.WriteString(value)
			whereQuery.WriteString("'")

		}
	}
	return whereQuery.String()
}

func (urlInfo *ClientInfo) GetLogValue(where models.LogSearchObj, table string, tableDiv string) (rst interface{}, error error) {
	url := urlInfo.Host + ":" + urlInfo.Port + urlInfo.Endpoint

	whereQuery := metaData(&where, tableDiv)
	fmt.Println(whereQuery)

	var query bytes.Buffer
	var str string
	if tableDiv == "container" {
		str = `SELECT \"__time\" AS collectDt, \"container_name\", \"hostname\", \"namespace\", \"pod_name\", \"loglevel\", \"log\" AS logMessage `
	}
	if tableDiv == "syslog" {
		str = `SELECT \"__time\" AS collectDt, \"host\", \"loglevel\", \"message\" AS logMessage , \"process\"`
	}

	query.WriteString(str)
	query.WriteString(`FROM \"`)
	query.WriteString(table)
	query.WriteString(`\" `)
	query.WriteString(whereQuery)

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
