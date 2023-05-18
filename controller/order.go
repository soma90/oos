package controller

import (
	"net/http"
	"oos/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//주문 신규 등록
func (p *Controller) InsertOrder(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, _id := p.md.InsertOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "_id": _id})
		return
	}	
}

//주문 내역 변경
func (p *Controller) UpdateOrder(c *gin.Context) {
	//주문id에 해당하는 주문 정보 가져오기
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//주문변경 상태인지 체크
	if err, result := p.md.FindOneOrder(order.ID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else if(result["status"].(int32) > 1) {
		c.JSON(409, gin.H{"error": "주문변경할 수 없는 상태입니다"})
		return
	}
	//업데이트 주문
	if err, upsertedID, modifiedCount := p.md.UpdateOrder(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "modifiedCount": modifiedCount, "upsertedID": upsertedID})
		return
	}	
}

//주문 상태 변경
func (p *Controller) UpdateOrderStatus(c *gin.Context) {
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, modifiedCount := p.md.UpdateOrderStatus(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "modifiedCount": modifiedCount})
		return
	}	
}

//주문id에 해당하는 주문 정보 가져오기
func (p *Controller) FindOneOrder(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err, result := p.md.FindOneOrder(_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "result": result})
		return
	}
}

//주문자의 모든 주문 내역 가져오기
func (p *Controller) FindOrder(c *gin.Context) {
	id := c.Param("id")
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err, result := p.md.FindOrder(_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "result": *result})
		return
	}
}