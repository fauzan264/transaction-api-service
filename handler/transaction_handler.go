package handler

import (
	"net/http"

	"github.com/fauzan264/transaction-api-service/helper"
	"github.com/fauzan264/transaction-api-service/transaction"
	"github.com/labstack/echo/v4"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService}
}

func (h *transactionHandler) WithdrawTransaction(c echo.Context) error {
	var input transaction.TransactionInput

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}

		c.Set("error", errorMessage)
		response := helper.APIResponse(false, "Failed to withdraw amount", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	createTransaction, err := h.transactionService.WithdrawTransaction(input)
	if err != nil {
		if helper.IsDatabaseError(err) {
			errorMessage := map[string]interface{}{"errors": err.Error()}
			c.Set("error", errorMessage)
			response := helper.APIResponse(false, "Internal server error", errorMessage)
			return c.JSON(http.StatusInternalServerError, response)
		}

		errorMessage := map[string]interface{}{"errors": err.Error()}
		c.Set("error", errorMessage)
		response := helper.APIResponse(false, "Failed to withdraw amount", errorMessage)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse(true, "Successfully withdraw amount", transaction.FormatTransactionBalance(createTransaction))
	return c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) SavingTransaction(c echo.Context) error {
	var input transaction.TransactionInput

	err := c.Bind(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := map[string]interface{}{"errors": errors}

		c.Set("error", errorMessage)
		response := helper.APIResponse(false, "Failed to deposit amount", errorMessage)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	createTransaction, err := h.transactionService.SavingTransaction(input)
	if err != nil {
		if helper.IsDatabaseError(err) {
			errorMessage := map[string]interface{}{"errors": err.Error()}

			c.Set("errors", errorMessage)
			response := helper.APIResponse(false, "Internal server error", errorMessage)
			return c.JSON(http.StatusInternalServerError, response)
		}

		errorMessage := map[string]interface{}{"errors": err.Error()}
		c.Set("error", errorMessage)
		response := helper.APIResponse(false, "Failed to saving amount", errorMessage)
		return c.JSON(http.StatusBadRequest, response)
	}

	response := helper.APIResponse(true, "Successfully saving amount", transaction.FormatTransactionBalance(createTransaction))
	return c.JSON(http.StatusOK, response)
}