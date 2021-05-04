package httpd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DruidHandler interface {
	GetColumnSearchInfo(c *gin.Context)
	GetLogValue(c *gin.Context)
}

func (h *Handler) GetColumnSearchInfo(c *gin.Context) {
	var searchInfoInit = make(map[string]interface{})

	columnKey := make(chan string)
	columnValue := make(chan []map[string]string)
	columns := []string{"hostname", "namespace_name", "pod_name", "container_name"}
	for _, column := range columns {
		go h.druid.GetColumnValue(column, "processed-containerlog", columnKey, columnValue)
		//fmt.Println(<-columnKey)
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
	var druidJson map[string]string
	if err := c.ShouldBindJSON(&druidJson); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Binding Error"})
		return
	}
	rst, err := h.druid.GetLogValue(druidJson, "processed-containerlog")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"body":   rst,
	})
}
