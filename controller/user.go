package controller

import (
	"net/http"
	"oos/model"

	"github.com/gin-gonic/gin"
)

//유저 신규 등록
func (p *Controller) InsertUser(c *gin.Context) {
	/* type User struct {
    Name string `bson:"name"`
	}
	var user User */
	
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, _id := p.md.InsertUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		c.JSON(200, gin.H{"msg": "ok", "_id": _id})
		return
	}	
}