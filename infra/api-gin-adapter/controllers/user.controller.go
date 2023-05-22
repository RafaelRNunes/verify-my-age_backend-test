package controllers

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application/dto"
	"github.com/RafaelRNunes/verify-my-age_backend-test/config"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	"strconv"
)

var userService = config.NewUserService()

// FindUsers godoc
// @Summary Return a list of users
// @Schemes
// @Description Return a list of users and address
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object} []dto.UserOutput
// @Router /users [get]
func FindUsers(ctx *gin.Context) {
	users := userService.FindAll()
	ctx.JSON(200, users)
}

// FindUserById godoc
// @Summary Return a list of users
// @Schemes
// @Description Return a list of users and address
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User Id"
// @Success 200 {object} dto.UserOutput
// @Failure 404 {string} httputil.HTTPError
// @Router /users/{id} [get]
func FindUserById(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	user := userService.FindById(userId)

	if user.Id == 0 {
		ctx.JSON(http.StatusNotFound, "User not found")
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// CreateUser godoc
// @Summary Create a user
// @Schemes
// @Description Create a user
// @Tags User
// @Accept json
// @Produce json
// @Param user body dto.UserInput true "User Body"
// @Success 201 {object} dto.UserOutput
// @Failure 400 {string} httputil.HTTPError
// @Router /users [post]
func CreateUser(ctx *gin.Context) {
	var user dto.UserInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userCreated := userService.Create(user)
	ctx.JSON(http.StatusCreated, userCreated)
}

// UpdateUser godoc
// @Summary Update a users
// @Schemes
// @Description Update user and address
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User Id"
// @Param user body dto.UserInput true "User Body"
// @Success 200 {object} dto.UserOutput
// @Failure 400 {string} httputil.HTTPError
// @Router /users/{id} [patch]
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

	if userUpdated.Id == 0 {
		ctx.JSON(http.StatusNotFound, "User not found to update")
		return
	}

	ctx.JSON(http.StatusOK, userUpdated)
}

// DeleteUser godoc
// @Summary Delete a user
// @Schemes
// @Description Delete a user and address
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User Id"
// @Success 200 {string} string
// @Failure 404 {string} httputil.HTTPError
// @Router /users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	wasDeleted := userService.Delete(userId)

	if !wasDeleted {
		ctx.JSON(http.StatusNotFound, "User not found to delete")
		return
	}

	ctx.JSON(http.StatusOK, "User deleted successfully")
}
