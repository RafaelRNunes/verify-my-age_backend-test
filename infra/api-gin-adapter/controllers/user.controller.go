package controllers

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/application/dto"
	"github.com/RafaelRNunes/verify-my-age_backend-test/application/validation"
	"github.com/RafaelRNunes/verify-my-age_backend-test/config"
	"github.com/RafaelRNunes/verify-my-age_backend-test/infra/api-gin-adapter/presenter"
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
	ctx.JSON(200, presenter.Response(true, users, ""))
}

// FindUserById godoc
// @Summary Return a user
// @Schemes
// @Description Return a user and address
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
		ctx.JSON(http.StatusNotFound, presenter.Response(false, "", "user not found"))
		return
	}

	ctx.JSON(http.StatusOK, presenter.Response(true, user, ""))
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
		ctx.JSON(http.StatusBadRequest, presenter.Response(false, "", err.Error()))
		return
	}

	if err := validation.ValidateUser(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, presenter.Response(false, "", err.Error()))
		return
	}

	userCreated := userService.Create(user)
	ctx.JSON(http.StatusCreated, presenter.Response(true, *userCreated, ""))
}

// UpdateUser godoc
// @Summary Update a user
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
		ctx.JSON(http.StatusBadRequest, presenter.Response(false, "", err.Error()))
		return
	}

	userUpdated := userService.Update(userId, user)

	if userUpdated.Id == 0 {
		ctx.JSON(http.StatusNotFound, presenter.Response(false, "", "user not found"))
		return
	}

	ctx.JSON(http.StatusOK, presenter.Response(true, userUpdated, ""))
}

// DeleteUser godoc
// @Summary Delete a user
// @Schemes
// @Description Delete a user and address
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User Id"
// @Success 204 {string} string
// @Failure 404 {string} httputil.HTTPError
// @Router /users/{id} [delete]
func DeleteUser(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Param("id"))
	wasDeleted := userService.Delete(userId)

	if !wasDeleted {
		ctx.JSON(http.StatusNotFound, presenter.Response(false, "", "user not found"))
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
