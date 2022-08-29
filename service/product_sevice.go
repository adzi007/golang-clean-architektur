package service

import (
	"cobasatu/model"
	mongodb "cobasatu/repository/mongo_repo"
	"fmt"
	"os"
	"time"

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

	return result, err
}

func (p productServiceObj) UpdateProduct(id string, pUpdate model.Product) error {

	pUpdate.Updated_at = time.Now().UTC()

	err := mongodb.PModel.Update(id, pUpdate)

	return err

}

func (p productServiceObj) DeleteImgGalery(id string, img model.ImgFile) error {

	err := mongodb.PModel.RemoveGalery(id, img)

	return err
}

func (p productServiceObj) DeleteProduct(id string) error {

	product, err := mongodb.PModel.GetProductById(id)

	if err != nil {

		return err
	}

	if product.Galery != nil {

		for _, item := range product.Galery {

			// fmt.Println(item.Thumbnail)
			//proses delete galery

			fullsizeimage := "assets/" + item.Fullsize
			thumbsizeimage := "assets/" + item.Thumbnail

			err1 := os.Remove(fullsizeimage)

			if err1 != nil {

				fmt.Println(err1)
				return err1
			}

			err2 := os.Remove(thumbsizeimage)

			if err2 != nil {

				fmt.Println(err2)
				return err2
			}

		}
	}

	result, err := mongodb.PModel.Delete(id)

	if result.DeletedCount != 1 {

		return err
	}

	return nil
}
