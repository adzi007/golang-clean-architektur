package service

import (
	"cobasatu/model"
	mongodb "cobasatu/repository/mongo_repo"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

type productServiceObj struct{}

var ProductService = productServiceObj{}

func (p productServiceObj) GetProduct(page string, search string) ([]*model.Product, int64) {

	resultData, totalna := mongodb.PModel.GetProduct(page, search)

	return resultData, totalna

}

func (p productServiceObj) InsertProduct(pInput model.Product) (*mongo.InsertOneResult, error) {

	validate := validator.New()

	err := validate.Struct(pInput)

	if err != nil {
		return nil, err
	}

	result, err := mongodb.PModel.Insert(pInput)

	// var savedImgGalery []model.ImageGalery

	return result, err
}
