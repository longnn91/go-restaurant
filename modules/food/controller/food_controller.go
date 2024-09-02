package transaction

import (
	"gogo/common"
	"gogo/modules/food/database"
	"gogo/modules/food/model"
	"gogo/modules/food/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFood(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.FoodCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := database.NewSQLStore(db)

		business := service.GetFoodService(store)

		//Create food
		new, err := business.CreateFood(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(new))
	}
}

func GetFoods(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		store := database.NewSQLStore(db)
		business := service.GetFoodService(store)
		data, err := business.GetFoods(&paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}

func GetFoodById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetFoodService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := business.GetFoodById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}

func UpdateFood(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetFoodService(store)

		var data model.FoodUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedData, err := business.UpdateFood(id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(updatedData))
	}
}

func DeleteFood(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetFoodService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := business.DeleteFood(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
