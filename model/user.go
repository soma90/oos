package model

import (
	"context"
	"fmt"
	"time"
)

//유저 추가
func (p *Model) InsertUser(user *User) (error, interface{}){
	newUser := User{
		Name: user.Name, 
		CreatedAt: time.Now(),
	}
	result, err := p.colUser.InsertOne(context.TODO(), newUser)
	if err != nil {
		return err, nil
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil, result.InsertedID
} 