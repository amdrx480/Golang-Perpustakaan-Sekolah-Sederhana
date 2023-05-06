package controllers

import (
	"net/http"
	"perpustakaan/models"
	"perpustakaan/services"

	"github.com/labstack/echo/v4"
)

type TransactionController struct {
	service services.TransactionService
}

func InitTransactionController() TransactionController {
	return TransactionController{
		service: services.InitTransactionService(),
	}
}

func (tc *TransactionController) GetAll(c echo.Context) error {
	transactions, err := tc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch transaction data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Transaction]{
		Status:  "success",
		Message: "all transactions",
		Data:    transactions,
	})
}

func (tc *TransactionController) GetByID(c echo.Context) error {
	transactionID := c.Param("id")

	transaction, err := tc.service.GetByID(transactionID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "transaction not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Transaction]{
		Status:  "success",
		Message: "transactions found",
		Data:    transaction,
	})
}

func (tc *TransactionController) Create(c echo.Context) error {
	var transactionInput models.TransactionInput

	if err := c.Bind(&transactionInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})

	}

	err := transactionInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	transaction, err := tc.service.Create(transactionInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to create a transaction",
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Transaction]{
		Status:  "success",
		Message: "book created",
		Data:    transaction,
	})
}

func (tc *TransactionController) Update(c echo.Context) error {
	transactionID := c.Param("id")
	var transactionInput models.TransactionInput

	if err := c.Bind(&transactionInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := transactionInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	transaction, err := tc.service.Update(transactionInput, transactionID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to update transaction",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Transaction]{
		Status:  "success",
		Message: "",
		Data:    transaction,
	})
}

func (tc *TransactionController) Delete(c echo.Context) error {
	var transactionID string = c.Param("id")

	err := tc.service.Delete(transactionID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete transaction",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "transaction deleted",
	})
}

func (tc *TransactionController) Restore(c echo.Context) error {
	var transactionID string = c.Param("id")

	transaction, err := tc.service.Restore(transactionID)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "book not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Transaction]{
		Status:  "success",
		Message: "book restored",
		Data:    transaction,
	})
}
func (tc *TransactionController) ForceDelete(c echo.Context) error {
	var transactionID string = c.Param("id")

	err := tc.service.ForceDelete(transactionID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to delete permanently",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "book deleted permanently",
	})
}
