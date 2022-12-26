package controller

import (
	"golang-mygram/model/domain"
	"golang-mygram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentService service.CommentService
}

func NewCommentController(s service.CommentService) *commentController {
	return &commentController{s}
}

func (c *commentController) GetComments(ctx *gin.Context) {
	comments, err := c.commentService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func (c *commentController) PostComment(ctx *gin.Context) {
	var commentReq domain.Comment

	err := ctx.ShouldBindJSON(&commentReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	comment, err := c.commentService.Create(commentReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func (c *commentController) UpdateComment(ctx *gin.Context) {
	var updateCommentReq domain.CommentUpdate

	err := ctx.ShouldBindJSON(&updateCommentReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	comment, err := c.commentService.Update(id, updateCommentReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func (c *commentController) DeleteComment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	comment, err := c.commentService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}

func (c *commentController) GetCommentBtId(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	comment, err := c.commentService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": comment,
	})
}
