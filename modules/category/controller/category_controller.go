package transaction

import (
	"gogo/common"
	"gogo/modules/category/database"
	"gogo/modules/category/model"
	"gogo/modules/category/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.CategoryCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := database.NewSQLStore(db)

		business := service.GetCategoryService(store)

		//Create category
		new, err := business.CreateCategory(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(new))
	}
}

func GetCategories(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		store := database.NewSQLStore(db)
		business := service.GetCategoryService(store)
		data, err := business.GetCategories(&paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}

func GetCategoryById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetCategoryService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := business.GetCategoryById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}

func UpdateCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetCategoryService(store)

		var data model.CategoryUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedData, err := business.UpdateCategory(id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(updatedData))
	}
}

func DeleteCategory(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetCategoryService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := business.DeleteCategory(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
