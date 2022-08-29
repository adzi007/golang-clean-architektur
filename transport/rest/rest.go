package rest

import (
	// "cobasatu/service"

	"cobasatu/entity"
	"cobasatu/model"
	"cobasatu/service"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/nfnt/resize"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RestServer() *gin.Engine {

	r := gin.Default()

	r.GET("/ping", HandlePing)
	r.GET("/product", HandleGetProduct)
	r.POST("/product", HandleInputProduct)

	return r

}

func HandlePing(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func HandleGetProduct(ctx *gin.Context) {

	resuldData, total := service.ProductService.GetProduct(ctx.Query("page"), ctx.Query("s"))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"total":   total,
		"data":    resuldData,
	})

}

func HandleInputProduct(ctx *gin.Context) {

	var pInput entity.ProductInput

	ctx.ShouldBind(&pInput)

	validate := validator.New()
	err := validate.Struct(pInput)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error_validation": err.Error()})
		return
	}

	// result, err := service.ProductService.InsertProduct(pInput)

	var savedImgGalery []model.ImageGalery

	for _, galeryItem := range pInput.Galery {

		// Retrieve file information
		extension := filepath.Ext(galeryItem.Filename)

		// Generate random file name for the new uploaded file
		newFileName := uuid.New().String() + extension

		err := ctx.SaveUploadedFile(galeryItem, "assets/"+newFileName)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "unknown error")
			return
		}

		createThumbnail(newFileName)

		imgGalery := model.ImageGalery{

			Fullsize:  newFileName,
			Thumbnail: "thumb-" + newFileName,
		}

		savedImgGalery = append(savedImgGalery, imgGalery)

	}

	product := model.Product{
		Id:          primitive.NewObjectID(),
		Name:        pInput.Name,
		Description: pInput.Description,
		Tags:        pInput.Tags,
		Galery:      savedImgGalery,
		Created_at:  time.Now().UTC(),
		Updated_at:  time.Now().UTC(),
	}

	result, err := service.ProductService.InsertProduct(product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "create product success",
		"data":    result,
	})

}

func createThumbnail(filename string) {

	file, err := os.Open("assets/" + filename)

	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(300, 0, img, resize.Lanczos3)

	out, err := os.Create("assets/thumb-" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)

}
