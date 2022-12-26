package controller

import (
	"golang-mygram/model/domain"
	"golang-mygram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaService service.SocialMediaService
}

func NewSocialMediaController(s service.SocialMediaService) *socialMediaController {
	return &socialMediaController{s}
}

func (c *socialMediaController) GetSocialMedias(ctx *gin.Context) {
	socialMedias, err := c.socialMediaService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": socialMedias})
}

func (c *socialMediaController) PostSocialMedia(ctx *gin.Context) {
	var socialMediaRequest domain.SocialMedia

	err := ctx.ShouldBindJSON(&socialMediaRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	socialMedia, err := c.socialMediaService.Create(socialMediaRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": socialMedia,
	})
}

func (c *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	var socialMediaReq domain.SocialMediaUpdate

	err := ctx.ShouldBindJSON(&socialMediaReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	socialMedia, err := c.socialMediaService.Update(id, socialMediaReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": socialMedia,
	})
}

func (c *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	socialMedia, err := c.socialMediaService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": socialMedia,
	})
}

func (c *socialMediaController) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	socialMedia, err := c.socialMediaService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": socialMedia,
	})
}
