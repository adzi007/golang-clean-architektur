package rest

import (
	// "cobasatu/service"

	"cobasatu/entity"
	"cobasatu/model"
	"cobasatu/service"
	"fmt"
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
	r.PUT("/product/:id", HandleUpdateProduct)
	r.DELETE("/product/:id", HandleDeleteProduct)
	r.DELETE("/product-img/:id", handleDeleteImage)

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

func HandleUpdateProduct(ctx *gin.Context) {

	var pInput entity.ProductInput

	ctx.ShouldBind(&pInput)

	validate := validator.New()
	err := validate.StructExcept(pInput, "Galery")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if pInput.Galery == nil {

		updateProducData := model.Product{
			Name:        pInput.Name,
			Description: pInput.Description,
			Tags:        pInput.Tags,
		}

		err := service.ProductService.UpdateProduct(ctx.Param("id"), updateProducData)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message_error1": err.Error()})
			return
		}

	} else {

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

		updateProducData := model.Product{
			Name:        pInput.Name,
			Description: pInput.Description,
			Tags:        pInput.Tags,
			Galery:      savedImgGalery,
		}

		err := service.ProductService.UpdateProduct(ctx.Param("id"), updateProducData)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message_error": err.Error()})
			return
		}

	}

	ctx.JSON(http.StatusOK, gin.H{

		"message": "proses update",
		"data":    pInput,
	})

}

func handleDeleteImage(ctx *gin.Context) {
	var img model.ImgFile

	if err := ctx.ShouldBindJSON(&img); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := service.ProductService.DeleteImgGalery(ctx.Param("id"), img)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message_error": err.Error()})
		return
	}

	fullsizeimage := "assets/" + img.Filename
	thumbsizeimage := "assets/thumb-" + img.Filename

	err1 := os.Remove(fullsizeimage)

	if err1 != nil {

		fmt.Println(err1)
		return
	}

	err2 := os.Remove(thumbsizeimage)

	if err2 != nil {

		fmt.Println(err2)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{

		"message": "berhasil delete image",
		"data":    img,
	})
}

func HandleDeleteProduct(ctx *gin.Context) {

	err := service.ProductService.DeleteProduct(ctx.Param("id"))

	if err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{

			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{

		"message": "success delete product",
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
