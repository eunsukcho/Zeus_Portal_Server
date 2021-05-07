package httpd

import (
	"fmt"
	"net/http"
	"zeus/models"

	"github.com/gin-gonic/gin"
)

type DruidHandler interface {
	GetColumnSearchInfo(c *gin.Context)
	GetLogValue(c *gin.Context)
}

func (h *Handler) GetColumnSearchInfo(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	var tableNm = uri.Table
	var columns []string
	if uri.Table == "container" {
		tableNm = "processed-containerlog"
		columns = append(columns, "hostname", "namespace_name", "pod_name", "container_name")
	}
	if uri.Table == "syslog" {
		tableNm = "processed-syslog"
		columns = append(columns, "host")
	}

	var searchInfoInit = make(map[string]interface{})
	columnKey := make(chan string)
	columnValue := make(chan []map[string]string)

	for _, column := range columns {
		go h.druid.GetColumnValue(column, tableNm, uri.Table, columnKey, columnValue)
		searchInfoInit[<-columnKey] = <-columnValue
	}

	loglevel, err := h.db.GetLogCode()
	fmt.Println(loglevel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	searchInfoInit["loglevel"] = loglevel

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"body":   searchInfoInit,
	})
}

func (h *Handler) GetLogValue(c *gin.Context) {
	var druidJson models.LogSearchObj
	if err := c.ShouldBindJSON(&druidJson); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Binding Error"})
		return
	}
	var tableNm string
	if druidJson.Table == "container" {
		tableNm = "processed-containerlog"
	}
	if druidJson.Table == "syslog" {
		tableNm = "processed-syslog"
	}
	rst, err := h.druid.GetLogValue(druidJson, tableNm, druidJson.Table)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"body":   rst,
	})
}
