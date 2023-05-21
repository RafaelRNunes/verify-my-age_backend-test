package controllers

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application/dto"
	"github.com/RafaelRNunes/verify-my-age_backend-test/config"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userService = config.NewUserService()

func FindUsers(ctx *gin.Context) {
	users := userService.FindAll()
	ctx.JSON(200, users)
}

func FindUserById(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user := userService.FindById(userId)

	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "User not found")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var user dto.UserInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userCreated := userService.Create(user)
	ctx.JSON(http.StatusCreated, userCreated)
}

func UpdateUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	var user dto.UserInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userUpdated := userService.Update(userId, user)
	ctx.JSON(http.StatusOK, userUpdated)
}

func DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	wasDeleted := userService.Delete(userId)

	if !wasDeleted {
		ctx.JSON(http.StatusNotFound, "User not found to delete")
		return
	}

	ctx.JSON(http.StatusOK, "User deleted successfully")
}
