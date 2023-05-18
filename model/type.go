package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model struct {
	client *mongo.Client
	colUser *mongo.Collection
	colMenu *mongo.Collection
	colOrder *mongo.Collection
	colReview *mongo.Collection
}

type Menu struct {
	Name string `bson:"name"`
	Price int `bson:"price"`
	Recommend	bool `bson:"recommend"`
	View bool `bson:"view"`
	CreatedAt time.Time `bson:"createdAt"`
}

type User struct {
	Name string `bson:"name"`
	CreatedAt time.Time `bson:"createdAt"`
}

type MenuItem struct {
	MenuID   primitive.ObjectID `bson:"menuID"`
	Quantity int `bson:"quantity"`
}

type Order struct {
	ID primitive.ObjectID `bson:"_id"`
	MenuList []MenuItem `bson:"menuList"`
	Status int `bson:"status"` //0: 접수중 1: 접수 2: 접수취소 3: 조리중 4: 배달중
	UserID primitive.ObjectID `bson:"userID"`
	Phone string `bson:"phone"`
	Addr string `bson:"addr"`
	Payment_Method string `bson:"payment_method"`
	CreatedAt time.Time `bson:"createdAt"`
}

type Review struct {
	OrderID primitive.ObjectID `bson:"orderID"`
	MenuID primitive.ObjectID `bson:"menuID"`
	UserID primitive.ObjectID `bson:"userID"`
	Rating float32 `bson:"rating"`
	Content string `bson:"content"`
	CreatedAt time.Time `bson:"createdAt"`
}