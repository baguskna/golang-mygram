package controller

import (
	"golang-mygram/model/domain"
	"golang-mygram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *photoController {
	return &photoController{photoService}
}

func (c *photoController) GetPhotos(ctx *gin.Context) {
	photos, err := c.photoService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": photos})
}

func (c *photoController) PostPhoto(ctx *gin.Context) {
	var photoRequest domain.Photo

	err := ctx.ShouldBindJSON(&photoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	photo, err := c.photoService.Create(photoRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

func (c *photoController) UpdatePhoto(ctx *gin.Context) {
	var photoReq domain.PhotoUpdate

	err := ctx.ShouldBindJSON(&photoReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	photo, err := c.photoService.Update(id, photoReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

func (c *photoController) DeletePhoto(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	photo, err := c.photoService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

func (c *photoController) GetPhotoById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	photo, err := c.photoService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}
