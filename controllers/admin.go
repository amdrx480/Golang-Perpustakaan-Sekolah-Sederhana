package controllers

import (
	"net/http"
	"perpustakaan/middlewares"
	"perpustakaan/models"
	"perpustakaan/services"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	service services.AdminService
}

func IniAdminController(jwtAuth *middlewares.JWTConfig) AdminController {
	return AdminController{
		service: services.InitAdminService(jwtAuth),
	}
}

func (ac *AdminController) Register(c echo.Context) error {
	var adminInput models.AdminInput

	if err := c.Bind(&adminInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := adminInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	admin, err := ac.service.Register(adminInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response[models.Admin]{
		Status:  "success",
		Message: "admin registered",
		Data:    admin,
	})
}

func (uc *AdminController) Login(c echo.Context) error {
	var adminInput models.AdminInput

	if err := c.Bind(&adminInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := adminInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	token, err := uc.service.Login(adminInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}
