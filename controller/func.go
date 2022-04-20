package controller

import (
	"day8-project/config"
	structs "day8-project/struct"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MakeItem(c *gin.Context) {
	var (
		Item  []structs.Items
		order structs.CreateOrders
	)

	config := config.GetDB()

	if rr := c.ShouldBindJSON(&order); rr != nil {
		c.AbortWithError(http.StatusBadRequest, rr)
		return
	}

	PutOrder := structs.Orders{
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
	}
	config.Create(&PutOrder)
	orderID := PutOrder.ID
	for _, v := range order.Item {
		item := structs.Items{
			ItemCode:    v.ItemCode,
			Quantity:    v.Quantity,
			Description: v.Description,
			OrderId:     PutOrder.ID,
		}
		Item = append(Item, item)
	}
	result := config.Create(&Item)
	log.Println(orderID, result.RowsAffected)

	respon := structs.Orders{
		ID:           Item[0].OrderId,
		CustomerName: order.CustomerName,
		OrderedAt:    time.Now(),
		Item:         Item,
	}
	c.JSON(http.StatusOK, respon)

}

func TakeItem(c *gin.Context) {
	config := config.GetDB()

	orders := []structs.Orders{}
	rr := config.Preload("Item").Find(&orders).Error

	if rr != nil {
		fmt.Println(rr.Error())
		return
	}
	c.JSON(http.StatusOK, orders)
}

func TakeDown(c *gin.Context) {
	config := config.GetDB()
	strID := c.Param("orderId")
	id, _ := strconv.Atoi(strID)
	order := structs.Orders{}
	item := structs.Items{}
	rr := config.Where("order_id = ?", id).Delete(&item).Error
	rowefc := config.Where("ID=?", id).Delete(&order).RowsAffected

	if rowefc == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"messege": "id not found",
		})
		return
	}

	if rr != nil {
		log.Println(rr.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messege": "data erased",
	})
}
func UpdateItem(c *gin.Context) {
	config := config.GetDB()

	stId := c.Param("orderID")
	id, _ := strconv.Atoi(stId)
	var order = structs.Orders{}

	if rr := c.ShouldBindJSON(&order); rr != nil {
		c.AbortWithError(http.StatusBadRequest, rr)
		return
	}

	updtOrder := structs.Orders{
		ID:           uint(id),
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
		Item:         order.Item,
	}
	config.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updtOrder)
	c.JSON(http.StatusOK, updtOrder)
}
