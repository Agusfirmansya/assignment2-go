package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"time"
	"net/http"
	"fmt"
)
type Orders struct {
	OrderID string
	CustomerName string
	OrderedAt time.Time
	Item []Items
}

type Items struct {
	ItemID string
	ItemCode string
	Description string
	Quantity uint
	OrderID string
}

var OrderDatas = []Orders{}

func GetAllOrder(ctx *gin.Context) {
	if len(OrderDatas) < 1 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H {
			"error_status": "No Data",
			"error_message": "No Data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": OrderDatas,
	})
}
func CreateOrder(ctx *gin.Context) {
	var newOrder Orders

	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newOrder.OrderID = fmt.Sprintf("%s", uuid.New())
	newOrder.OrderedAt = time.Now()

	for i,_ := range newOrder.Item {
		newOrder.Item[i].ItemID = fmt.Sprintf("%s", uuid.New())
		newOrder.Item[i].OrderID = newOrder.OrderID
	}
	OrderDatas = append(OrderDatas, newOrder)

	ctx.JSON(http.StatusCreated, gin.H{
		"order": newOrder,
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	condition := false
	var updateOrder Orders
	var itemID []string

	if err := ctx.ShouldBindJSON(&updateOrder); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}



	for i, order := range OrderDatas {
		if orderID == order.OrderID {
			for idx, _ := range OrderDatas[i].Item {
				itemID = append(itemID, OrderDatas[i].Item[idx].ItemID)
			}
			condition = true
			OrderDatas[i] = updateOrder
			OrderDatas[i].OrderID = orderID
			for idx,_ := range OrderDatas[i].Item {
				OrderDatas[i].Item[idx].ItemID = itemID[idx]
				OrderDatas[i].Item[idx].OrderID = orderID
			}
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "data not found",
			"error_message": orderID+", not found",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "update success",
	})
}

func GetOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	condition := false
	var orderData Orders

	for i, order := range OrderDatas {
		if orderID == order.OrderID {
			condition = true
			orderData = OrderDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "data not found",
			"error_message": orderID+", not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"order": orderData,
	})
}

func DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("orderID")
	condition := false
	var orderIndex int

	for i, order := range OrderDatas {
		if orderID == order.OrderID {
			condition = true
			orderIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "data not found",
			"error_message": orderID+", not found",
		})
		return
	}

	copy(OrderDatas[orderIndex:], OrderDatas[orderIndex+1:])
	OrderDatas[len(OrderDatas)-1] = Orders{}
	OrderDatas = OrderDatas[:len(OrderDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": orderID+" deleted",
	})
}