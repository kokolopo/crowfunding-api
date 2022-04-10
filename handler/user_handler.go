package handler

import (
	"crowfunding_api/helper"
	"crowfunding_api/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.InputRegistration

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errMessage := gin.H{"errors": errors}

		response := helper.ResponseApi("account registration failed", http.StatusUnprocessableEntity, "failed", errMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, errUser := h.userService.RegisterUser(input)
	if errUser != nil {
		response := helper.ResponseApi("account registration failed", http.StatusBadRequest, "failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := user.FormatUser(newUser, "")

	response := helper.ResponseApi("account registration successful", http.StatusCreated, "success", data)

	c.JSON(http.StatusCreated, response)
}
