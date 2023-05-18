package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//메뉴 신규 등록
/* func (p *Model) InsertMenu(name string, price int) (error){
	newMenu := Menu{
		Name: name,
		Price: price,
		Recommend: false,
		View: true,
		CreatedAt: time.Now(),
	}
	result, err := p.colMenu.InsertOne(context.TODO(), newMenu)
	if err != nil {
		//panic(err)
		return err
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil
} */

//메뉴 신규 등록
func (p *Model) InsertMenu(name string, price int) (error, interface{}){
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	update := bson.M{"$set": bson.D{
		primitive.E{Key: "price", Value: price},
		primitive.E{Key: "recommend", Value: false},
		primitive.E{Key: "view", Value: true},
		primitive.E{Key: "createdAt", Value: time.Now()},
	}}	

	result, err := p.colMenu.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err, nil
	}
	//fmt.Printf("Document upserted with ID: %s\n", result.UpsertedID)
	return nil, result.UpsertedID
}

//메뉴 삭제
func (p *Model) DelMenu(name string) (error, int64){
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	update := bson.M{"$set": bson.D{
		primitive.E{Key: "view", Value: false},
	}}	
	result, err := p.colMenu.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err, 0
	}
	//fmt.Printf("Modified documents: %d\n", result.ModifiedCount)
	return nil, result.ModifiedCount
} 

//메뉴 추천 or 기본
func (p *Model) RecommendMenu(name string, recommend bool) (error, int64){
	filter := bson.D{primitive.E{Key: "name", Value: name}}
	update := bson.M{"$set": bson.D{
		primitive.E{Key: "recommend", Value: recommend},
	}}	
	result, err := p.colMenu.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err, 0
	}
	//fmt.Printf("Modified documents: %d\n", result.ModifiedCount)
	return nil, result.ModifiedCount
} 