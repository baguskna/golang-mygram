package controller

import (
	"golang-mygram/helpers"
	"golang-mygram/model/domain"
	"golang-mygram/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (c *userController) GetUsers(ctx *gin.Context) {
	users, err := c.userService.FindAll()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (c *userController) PostUser(ctx *gin.Context) {
	var userRequest domain.User

	err := ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err := c.userService.Create(userRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	var userReq domain.UserUpdate

	err := ctx.ShouldBindJSON(&userReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := c.userService.Update(id, userReq)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (c *userController) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := c.userService.Delete(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete the data",
	})
}

func (c *userController) GetUserById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)

	user, err := c.userService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (c *userController) LoginUser(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	user := domain.User{}

	if contentType == "application/json" {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	getUser, err := c.userService.Login(user)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(getUser.Password), []byte(user.Password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(uint(getUser.ID), user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
