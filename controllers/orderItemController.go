package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrderItems() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func GetOrderItemByOrder() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	return OrderItems, err
}

func GetOrderItem() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(context *gin.Context) {

	}
}
