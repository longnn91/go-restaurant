package transaction

import (
	"gogo/common"
	"gogo/modules/menu/database"
	"gogo/modules/menu/model"
	"gogo/modules/menu/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMenu(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.MenuCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := database.NewSQLStore(db)

		business := service.GetMenuService(store)

		//Create menu
		new, err := business.CreateMenu(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(new))
	}
}

func GetMenus(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		store := database.NewSQLStore(db)
		business := service.GetMenuService(store)
		data, err := business.GetMenus(&paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}

func GetMenuById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetMenuService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := business.GetMenuById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}

func UpdateMenu(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetMenuService(store)

		var data model.MenuUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedData, err := business.UpdateMenu(id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(updatedData))
	}
}

func DeleteMenu(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetMenuService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := business.DeleteMenu(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
