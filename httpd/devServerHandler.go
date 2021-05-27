package httpd

import (
	"fmt"
	"net/http"
	"zeus/models"

	"github.com/gin-gonic/gin"
)

type DevServerHandler interface {
	GetAllDevServerInfoData(c *gin.Context)
	GetDevServerInfoDataById(c *gin.Context)

	SaveDevServerInfo(c *gin.Context)

	UpdateDevServerInfo(c *gin.Context)
	DeleteDevServerInfo(c *gin.Context)
}

func (h *Handler) GetAllDevServerInfoData(c *gin.Context) {
	devInfo, err := h.db.GetAllDevServerInfoData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(devInfo))
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   devInfo,
		"len":    len(devInfo),
	})
}

func (h *Handler) GetDevServerInfoDataById(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.ReqId)
	devInfoById, err := h.db.GetDevServerInfoDataById(uri.ReqId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, devInfoById)
}

func (h *Handler) SaveDevServerInfo(c *gin.Context) {
	var dev models.DevServerModel
	err := c.ShouldBindJSON(&dev)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	saveDevInfo, err := h.db.SaveDevServerInfo(dev)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   saveDevInfo,
	})
}

func (h *Handler) UpdateDevServerInfo(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	var dev models.DevServerModel
	err := c.ShouldBindJSON(&dev)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusBadRequest,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	updateDevInfo, err := h.db.UpdateDevServerInfo(dev, uri.ReqId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   updateDevInfo,
	})
}

func (h *Handler) DeleteDevServerInfo(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	err := h.db.DeleteDevServerInfo(uri.ReqId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": http.StatusText(http.StatusBadRequest),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   http.StatusText(http.StatusOK),
	})
}
