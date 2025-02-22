package handler

import (
	"log"
	"net/http"

	"github.com/fauzan264/transaction-api-service/helper"
	"github.com/fauzan264/transaction-api-service/user"
	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c echo.Context) error {
	var input user.RegisterUserinput

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}

		c.Set("error", errorMessage)
		response := helper.APIResponse(false, "Register account failed.", errorMessage)
		return c.JSON(http.StatusBadRequest, response)
	}

	createUser, err := h.userService.RegisterUser(input)
	if err != nil {
		if helper.IsDatabaseError(err) {
			// If the issue is related to the database (e.g., constraint violation or query issues)
			errorMessage := map[string]interface{}{"errors": err.Error()}
			c.Set("error", errorMessage)
			response := helper.APIResponse(false, "Internal server error during user registration.", errorMessage)
			return c.JSON(http.StatusInternalServerError, response)
		}

		errorMessage := map[string]interface{}{"errors": err.Error()}
		c.Set("error", errorMessage)
		response := helper.APIResponse(false, "Register account failed.", errorMessage)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse(true, "Account has been registered", user.FormatUserCreate(createUser))
	return c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetBalance(c echo.Context) error {
	getNumberBalance := c.Param("number_balance")
	log.Println(getNumberBalance)
	
	userBalance, err := h.userService.GetBalance(getNumberBalance)
	if err != nil {
		errorMessage := map[string]interface{}{"errors": err.Error()}
		c.Set("error", errorMessage)
		response := helper.APIResponse(false, "Failed to check user balance", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	formatter := user.FormatUserBalance(userBalance)
	response := helper.APIResponse(true, "Success to check user balance", formatter)
	return c.JSON(http.StatusOK, response)
}