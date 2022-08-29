package entity

import (
	"mime/multipart"
)

type ProductInput struct {
	Id          int                     `uri:"id"`
	Name        string                  `form:"name" binding:"required" validate:"required"`
	Tags        []string                `form:"tags" bson:"tags" binding:"required" validate:"required"`
	Description string                  `form:"description" binding:"required" validate:"required"`
	Galery      []*multipart.FileHeader `form:"galery" binding:"required" validate:"required"`
}

// type Product struct {
// 	Id          primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
// 	Name        string             `json:"name" bson:"name"`
// 	Description string             `json:"description" bson:"description"`
// 	Galery      []ImageGalery      `json:"galery" bson:"galery"`
// 	Tags        []string           `json:"tags" bson:"tags"`
// 	Created_at  time.Time          `json:"created_at,omitempty"`
// 	Updated_at  time.Time          `json:"updated_at,omitempty"`
// }
