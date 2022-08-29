package mongo_repo

import (
	"cobasatu/config"
	"cobasatu/model"
	"context"
	"log"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var productCollection *mongo.Collection = config.GetCollection(config.DBConnection, "product")

type ProductImpl struct{}

var PModel ProductModel = ProductImpl{}

func ConnectMongo() {
	config.MongoDBConfig.ConnectDB()
}

func (p ProductImpl) GetProduct(page string, search string) (result []*model.Product, total int64) {

	filter := bson.M{}
	findOptions := options.Find()
	var episodes []*model.Product

	pageFilter, _ := strconv.Atoi(page)
	var perPage int64 = 10

	if search != "" {

		filter = bson.M{
			"$or": []bson.M{
				{
					"name": bson.M{
						"$regex": primitive.Regex{
							Pattern: search,
							Options: "i",
						},
					},
				},
				{
					"description": bson.M{
						"$regex": primitive.Regex{
							Pattern: search,
							Options: "i",
						},
					},
				},
			},
		}

	}

	findOptions.SetSkip((int64(pageFilter) - 1) * perPage)
	findOptions.SetLimit(perPage)

	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := productCollection.Find(context.Background(), filter, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	if err := cursor.All(context.Background(), &episodes); err != nil {
		log.Fatal(err)
	}

	var totalna, _ = productCollection.CountDocuments(context.Background(), filter)

	return episodes, totalna

}

func (p ProductImpl) Insert(product model.Product) (*mongo.InsertOneResult, error) {

	return productCollection.InsertOne(context.Background(), product)
}

func (p ProductImpl) Update(id string, product model.Product) error {

	objId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objId}

	product.Id = objId

	_, err := productCollection.UpdateOne(context.Background(), filter, bson.M{"$set": product})

	return err
}

func (p ProductImpl) GetProductById(id string) (model.Product, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}

	var product model.Product
	err := productCollection.FindOne(context.Background(), filter).Decode(&product)
	return product, err

}

func (p ProductImpl) Delete(id string) (*mongo.DeleteResult, error) {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}
	return productCollection.DeleteOne(context.Background(), filter)
}

func (p ProductImpl) RemoveGalery(id string, img model.ImgFile) error {

	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objId}

	_, err := productCollection.UpdateOne(context.Background(), filter, bson.M{
		"$pull": bson.M{
			"galery": bson.M{
				"full_size": img.Filename,
			},
		},
	})

	return err

}
