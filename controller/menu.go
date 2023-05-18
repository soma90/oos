package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//메뉴 신규 등록
func (p *Controller) InsertMenu(c *gin.Context) {
	type Menu struct {
    Name  string `bson:"name"`
    Price int    `bson:"price"`
	}
	var menu Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, upsertedID := p.md.InsertMenu(menu.Name, menu.Price); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "upsertedID": upsertedID})
		return
	}	
}

//메뉴 삭제
func (p *Controller) DelMenu(c *gin.Context) {
	type Menu struct {
    Name  string `bson:"name"`
	}
	var menu Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, modifiedCount := p.md.DelMenu(menu.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "modifiedCount": modifiedCount})
		return
	}	
}

//메뉴 추천 or 기본
func (p *Controller) RecommendMenu(c *gin.Context) {
	type Menu struct {
    Name  string `bson:"name"`
		Recommend bool `bson:"recommend"`
	}
	var menu Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, modifiedCount := p.md.RecommendMenu(menu.Name, menu.Recommend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "modifiedCount": modifiedCount})
		return
	}	
}