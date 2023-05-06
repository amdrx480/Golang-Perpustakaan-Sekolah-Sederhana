package controllers

import (
	"net/http"
	"perpustakaan/models"
	"perpustakaan/services"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	service services.BookService
}

func InitBookController() BookController {
	return BookController{
		//penyebab 		service: services.InitBookService() adalah ini
		// service: services.BookService{},
		service: services.InitBookService(),
	}
}

func (bc *BookController) GetAll(c echo.Context) error {
	book, err := bc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch books data",
		})
	}
	return c.JSON(http.StatusOK, models.Response[[]models.Book]{
		Status:  "success",
		Message: "all Books",
		Data:    book,
	})
}

func (bc *BookController) GetByID(c echo.Context) error {
	var bookID string = c.Param("id")

	book, err := bc.service.GetByID(bookID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "book not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Book]{
		Status:  "success",
		Message: "book found",
		Data:    book,
	})
}
func (bc *BookController) Create(c echo.Context) error {
	var bookInput models.BookInput

	if err := c.Bind(&bookInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})

	}

	err := bookInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "validation failed",
		})
	}

	book, err := bc.service.Create(bookInput)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, models.Response[models.Book]{
		Status:  "success",
		Message: "book created",
		Data:    book,
	})
}
func (bc *BookController) Update(c echo.Context) error {
	var bookID string = c.Param("id")

	var bookInput models.BookInput

	if err := c.Bind(&bookInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := bookInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	book, err := bc.service.Update(bookInput, bookID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response[models.Book]{
		Status:  "success",
		Message: "book updated",
		Data:    book,
	})

}
func (bc *BookController) Delete(c echo.Context) error {
	var bookID string = c.Param("id")

	err := bc.service.Delete(bookID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "book deleted",
	})
}

func (bc *BookController) Restore(c echo.Context) error {
	var bookID string = c.Param("id")

	book, err := bc.service.Restore(bookID)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "book not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Book]{
		Status:  "success",
		Message: "book restored",
		Data:    book,
	})
}

func (bc *BookController) ForceDelete(c echo.Context) error {
	var bookID string = c.Param("id")

	err := bc.service.ForceDelete(bookID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "book deleted permanently",
	})
}
