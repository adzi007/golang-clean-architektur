package mongo_repo

import (
	"cobasatu/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductModel interface {
	// ConnectMongo()
	GetProduct(string, string) ([]*model.Product, int64)
	GetProductById(string) (model.Product, error)
	Insert(model.Product) (*mongo.InsertOneResult, error)
	Update(string, model.Product) error
	Delete(string) (*mongo.DeleteResult, error)
	RemoveGalery(string, model.ImgFile) error
}
