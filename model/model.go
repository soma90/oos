package model

import (
	"context"
	conf "oos/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewModel() (*Model, error) {
	r := &Model{}
	
	config := conf.GetConfig("./config/config.toml")	
	mgUrl := config.DB["user"]["host"].(string)
	dbName := config.DB["user"]["name"].(string)
	var err error

	if r.client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mgUrl)); err != nil {
		return nil, err
	} else if err := r.client.Ping(context.Background(), nil); err != nil {
		return nil, err
	} else {
		db := r.client.Database(dbName)
		r.colUser = db.Collection("user")
		r.colMenu = db.Collection("menu")
		r.colOrder = db.Collection("order")
		r.colReview = db.Collection("review")
	}

	return r, nil
}