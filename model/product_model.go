package model

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UploadImageGalery struct {
	Images []*multipart.FileHeader `form:"images" binding:"required"`
}

type ImageGalery struct {
	Fullsize  string `json:"full_size" bson:"full_size"`
	Thumbnail string `json:"tumbnail" bson:"tumbnail"`
}

type ImgFile struct {
	Filename string `json:"filename"`
}

type ProductInput struct {
	Id          int                     `uri:"id"`
	Name        string                  `form:"name" binding:"required" validate:"required"`
	Tags        []string                `form:"tags" bson:"tags" binding:"required" validate:"required"`
	Description string                  `form:"description" binding:"required" validate:"required"`
	Galery      []*multipart.FileHeader `form:"galery" binding:"required" validate:"required"`
}

type Product struct {
	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name        string             `json:"name" bson:"name" binding:"required" validate:"required"`
	Description string             `json:"description" bson:"description" binding:"required" validate:"required"`
	Galery      []ImageGalery      `json:"galery" bson:"galery" binding:"required" validate:"required"`
	Tags        []string           `json:"tags" bson:"tags" binding:"required" validate:"required"`
	Created_at  time.Time          `json:"created_at,omitempty"`
	Updated_at  time.Time          `json:"updated_at,omitempty"`
}
