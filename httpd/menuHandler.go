package httpd

import (
	"fmt"
	"net/http"
	"zeus/models"

	"github.com/gin-gonic/gin"
)

type MenuHandler interface {

	//menu setting
	GetTopMenuInfoByName(c *gin.Context)
	GetMainView(c *gin.Context)
	GetTopMenuData(c *gin.Context)
	SubTopMenuData(c *gin.Context)
	CkDuplicateTopMenu(c *gin.Context)
	SaveTopMenu(c *gin.Context)
	CkDuplicateSubMenu(c *gin.Context)
	SaveSubMenu(c *gin.Context)
	DeleteTopMenu(c *gin.Context)
	DeleteSubMenu(c *gin.Context)
	GetIcon(c *gin.Context)
	CkDuplicateIsMain(c *gin.Context)
	SaveUrlLink(c *gin.Context)
	SaveUrlSubLink(c *gin.Context)
	DeleteTopMenuUrl(c *gin.Context)
	DeleteSubMenuByTopCodeUrl(c *gin.Context)
	DeleteSubMenuUrl(c *gin.Context)
	GetMenuTargetUrl(c *gin.Context)
	GetTopMenuTargetUrl(c *gin.Context)
	UpdateTopMenuInfo(c *gin.Context)
	UpdateSubMenuInfo(c *gin.Context)
	UpdateSubMenuTopCodeName(c *gin.Context)
}

//menu setting
func (h *Handler) GetTopMenuInfoByName(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.TopCodeName)

	top_menu, err := h.db.GetTopMenuInfoByName(uri.TopCodeName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   top_menu,
	})
}

func (h *Handler) GetMainView(c *gin.Context) {
	main_view, err := h.db.GetMainView()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, main_view)
}

func (h *Handler) GetTopMenuData(c *gin.Context) {

	top_menu, err := h.db.GetAllTopMenu()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(top_menu))
	c.JSON(http.StatusOK, top_menu)
}
func (h *Handler) SubTopMenuData(c *gin.Context) {

	sub_menu, err := h.db.GetAllSubMenu()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(sub_menu))
	c.JSON(http.StatusOK, sub_menu)
}
func (h *Handler) CkDuplicateTopMenu(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("userID : ", uri.TopCode)

	rst, err := h.db.CkDuplicateTopMenu(uri.TopCode)
	order, err := h.db.CkDuplicateTopMenuOrder(uri.Order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"topCodeLen": rst,
		"topOrder":   order,
	})
}
func (h *Handler) SaveTopMenu(c *gin.Context) {
	var topMenu models.TopMenuInfo
	err := c.ShouldBindJSON(&topMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	fmt.Printf("topMenu :", topMenu)

	rst, err := h.db.SaveTopMenuInfo(topMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}

func (h *Handler) CkDuplicateSubMenu(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("CkDuplicateSubMenu uri : ", uri)

	rst, err := h.db.CkDuplicateSubMenu(uri.TopCode, uri.SubCode)
	order, err := h.db.CkDuplicateSubMenuOrder(uri.TopCode, uri.SubCode, uri.Order)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"subCodeLen": rst,
		"subOrder":   order,
	})
}

func (h *Handler) SaveSubMenu(c *gin.Context) {
	var subMenu models.SubMenuInfo

	err := c.ShouldBindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.SaveSubMenuInfo(subMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) DeleteTopMenu(c *gin.Context) {
	var topMenu models.TopMenuInfo

	err := c.BindJSON(&topMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.DeleteTopMenuInfo(topMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) DeleteSubMenu(c *gin.Context) {
	var subMenu models.SubMenuInfo

	err := c.BindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.DeleteSubMenuInfo(subMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) GetIcon(c *gin.Context) {

	icon, err := h.db.GetAllIcon()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, icon)
}

func (h *Handler) CkDuplicateIsMain(c *gin.Context) {
	var topMenu models.TopMenuInfo
	err := c.ShouldBindJSON(&topMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.CkDuplicateIsMain(topMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"body":   rst,
	})
}

func (h *Handler) SaveUrlLink(c *gin.Context) {

	var topMenu models.TopMenuInfo
	err := c.ShouldBindJSON(&topMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	fmt.Println("topMenu : ", topMenu)
	rst, err := h.db.SaveUrlLink(topMenu)
	fmt.Println(rst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})

}
func (h *Handler) SaveUrlSubLink(c *gin.Context) {
	var subMenu models.SubMenuInfo
	err := c.ShouldBindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.SaveUrlSubLink(subMenu)
	fmt.Println(rst)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})

}

func (h *Handler) DeleteTopMenuUrl(c *gin.Context) {
	var topMenu models.TopMenuInfo
	err := c.BindJSON(&topMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}

	rst, err := h.db.DeleteTopMenuUrl(topMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}

func (h *Handler) DeleteSubMenuByTopCodeUrl(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("CkDuplicateSubMenu uri : ", uri)

	rst, err := h.db.DeleteSubMenuByTopCodeUrl(uri.TopCodeName)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}

func (h *Handler) DeleteSubMenuUrl(c *gin.Context) {
	var subMenu models.SubMenuInfo
	err := c.BindJSON(&subMenu)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		return
	}
	rst, err := h.db.DeleteSubMenuUrl(subMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
func (h *Handler) GetMenuTargetUrl(c *gin.Context) {
	var menuCode models.SubMenuInfo
	err := c.ShouldBindJSON(&menuCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	fmt.Println("MenuCode :", menuCode)

	rst, err := h.db.GetMenuTargetUrl(menuCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println("targetUrl : ", rst)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst.Sub_Menu_Target_Url,
	})

}
func (h *Handler) GetTopMenuTargetUrl(c *gin.Context) {
	var menuCode models.TopMenuInfo
	err := c.ShouldBindJSON(&menuCode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	fmt.Println("MenuCode :", menuCode)

	rst, err := h.db.GetTopMenuTargetUrl(menuCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println("targetUrl : ", rst)
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst.Top_Menu_Target_Url,
	})

}

func (h *Handler) UpdateTopMenuInfo(c *gin.Context) {
	var topMenu models.TopMenuInfo
	err := c.ShouldBindJSON(&topMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.UpdateTopMenuInfo(topMenu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})

}

func (h *Handler) UpdateSubMenuInfo(c *gin.Context) {
	var subMenu models.SubMenuInfo
	err := c.ShouldBindJSON(&subMenu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"isOK":   0,
			"error":  err,
		})
		fmt.Println(err)
		return
	}
	rst, err := h.db.UpdateSubMenuInfo(subMenu)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}

func (h *Handler) UpdateSubMenuTopCodeName(c *gin.Context) {
	var uri models.Uri
	if err := c.ShouldBindUri(&uri); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	fmt.Println("UpdateSubMenuTopCodeName uri : ", uri)
	rst, err := h.db.UpdateSubMenuTopCodeName(uri.TopCode, uri.TopCodeName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"isOK":   1,
		"data":   rst,
	})
}
