package transaction

import (
	"gogo/common"
	"gogo/modules/order/database"
	"gogo/modules/order/model"
	"gogo/modules/order/service"
	"gogo/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		userId, err := utils.GetUserIdFromContext(c)
		if err != nil {
			return
		}

		var data model.OrderCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data.UserId = userId

		store := database.NewSQLStore(db)

		business := service.GetOrderService(store)

		//Create order
		new, err := business.CreateOrder(&data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(new))
	}
}

func GetOrders(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		store := database.NewSQLStore(db)
		business := service.GetOrderService(store)
		data, err := business.GetOrders(&paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, nil))
	}
}

func GetOrderById(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetOrderService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		data, err := business.GetOrderById(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}
}

func UpdateOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetOrderService(store)

		var data model.OrderUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedData, err := business.UpdateOrder(id, &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(updatedData))
	}
}

func DeleteOrder(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		store := database.NewSQLStore(db)
		business := service.GetOrderService(store)

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if _, err := business.DeleteOrder(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(nil))
	}
}
