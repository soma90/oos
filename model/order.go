package model

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//주문 추가
func (p *Model) InsertOrder(order *Order) (error, interface{}){
	newOrder := Order{
		MenuList: order.MenuList, 
		Status: 0,
		UserID: order.UserID,
		Phone: order.Phone,
		Addr: order.Addr,
		Payment_Method: order.Payment_Method,
		CreatedAt: time.Now(),
	}
	result, err := p.colOrder.InsertOne(context.TODO(), newOrder)
	if err != nil {
		return err, nil
	}
	//fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
	return nil, result.InsertedID
} 

//주문 변경
func (p *Model) UpdateOrder(order *Order) (error, interface{}, int64){
	filter := bson.D{primitive.E{Key: "_id", Value: order.ID}}
	update := bson.M{"$set": bson.D{
		primitive.E{Key: "menuList", Value: order.MenuList},
	}}	

	result, err := p.colOrder.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return err, nil, 0
	}
	//fmt.Printf("Document upserted with ID: %s\n", result.UpsertedID)
	return nil, result.UpsertedID, result.ModifiedCount
}

//주문 상태 변경
func (p *Model) UpdateOrderStatus(order *Order) (error, int64){
	filter := bson.D{primitive.E{Key: "_id", Value: order.ID}}
	update := bson.M{"$set": bson.D{
		primitive.E{Key: "status", Value: order.Status},
	}}	

	result, err := p.colOrder.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err, 0
	}
	return nil, result.ModifiedCount
}

//주문Id에 해당하는 주문 정보 가져오기
func (p *Model) FindOneOrder(orderID primitive.ObjectID) (error, primitive.M) {
	filter := bson.D{primitive.E{Key: "_id", Value: orderID}}
	var result bson.M
	err := p.colOrder.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return err, nil
	}
	return nil, result
}

//주문자의 모든 주문 내역 가져오기
func (p *Model) FindOrder(userID primitive.ObjectID) (error, *[]bson.M) {
	filter := bson.D{primitive.E{Key: "userID", Value: userID}}
	cursor, err := p.colOrder.Find(context.TODO(), filter)
	if err != nil {
		return err, nil
	}

	var result []bson.M
	if err = cursor.All(context.TODO(), &result); err != nil {
		return err, nil
	}

	return nil, &result
	//fmt.Println(result)   //law data 확인
	// if jsonData, err := json.MarshalIndent(result, "", "    "); err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Printf("%s\n", jsonData)
	// }
}
