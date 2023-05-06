package controllers

import (
	"net/http"
	"perpustakaan/models"
	"perpustakaan/services"

	"github.com/labstack/echo/v4"
)

type MemberController struct {
	service services.MemberService
}

func InitMemberkController() MemberController {
	return MemberController{
		service: services.InitMemberService(),
	}
}

func (mc *MemberController) GetAll(c echo.Context) error {
	member, err := mc.service.GetAll()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: "failed to fetch members data",
		})
	}

	return c.JSON(http.StatusOK, models.Response[[]models.Member]{
		Status:  "success",
		Message: "all Members",
		Data:    member,
	})
}

func (mc *MemberController) GetByID(c echo.Context) error {
	var memberID string = c.Param("id")

	member, err := mc.service.GetByID(memberID)

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "member not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Member]{
		Status:  "success",
		Message: "member found",
		Data:    member,
	})
}

func (mc *MemberController) Create(c echo.Context) error {
	var memberInput models.MemberInput

	if err := c.Bind(&memberInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})

	}

	err := memberInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	member, err := mc.service.Create(memberInput)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, models.Response[models.Member]{
		Status:  "success",
		Message: "member created",
		Data:    member,
	})
}
func (mc *MemberController) Update(c echo.Context) error {
	var memberID string = c.Param("id")

	var memberInput models.MemberInput

	if err := c.Bind(&memberInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	err := memberInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	member, err := mc.service.Update(memberInput, memberID)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response[models.Member]{
		Status:  "success",
		Message: "member updated",
		Data:    member,
	})

}
func (mc *MemberController) Delete(c echo.Context) error {
	var memberID string = c.Param("id")

	err := mc.service.Delete(memberID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "member deleted",
	})
}

func (mr *MemberController) Restore(c echo.Context) error {
	var memberID string = c.Param("id")

	member, err := mr.service.Restore(memberID)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response[string]{
			Status:  "failed",
			Message: "member not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response[models.Member]{
		Status:  "success",
		Message: "member restored",
		Data:    member,
	})
}

func (mr *MemberController) ForceDelete(c echo.Context) error {
	var memberID string = c.Param("id")

	err := mr.service.ForceDelete(memberID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response[string]{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Response[string]{
		Status:  "success",
		Message: "member deleted permanently",
	})
}
