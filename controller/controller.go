package controller

import (
	"oos/model"
)

type Controller struct{
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{md : rep}
	return r, nil
}