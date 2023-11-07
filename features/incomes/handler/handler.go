package handler

import (
	"net/http"
	"strconv"
	"strings"
	"technopartner/app/middlewares"
	"technopartner/features/incomes"
	"technopartner/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type IncomeHandler struct {
	incomeService incomes.IncomeServiceInterface
}

func New(service incomes.IncomeServiceInterface) *IncomeHandler {
	return &IncomeHandler{
		incomeService: service,
	}
}

func (handler *IncomeHandler) CreateIncome(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	var incomeInput IncomeRequest
	if err := c.Bind(&incomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error binding data. Data is not valid.", nil))
	}

	validate := validator.New()
	if err := validate.Struct(incomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	incomeCore := IncomeRequestToCore(incomeInput)
	err := handler.incomeService.CreateIncome(UserID, incomeCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error inserting data", nil))
		}
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success inserting data", nil))
}

func (handler *IncomeHandler) GetAllIncomes(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	result, err := handler.incomeService.GetAllIncomes(UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error retrieving incomes", nil))
	}
	var incomeResponse []IncomeResponseAll
	for _, value := range result {
		incomeResponse = append(incomeResponse, IncomeResponseAll{
			Name:   value.Name,
			Jumlah: value.Jumlah,
		})
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success retrieving incomes", incomeResponse))
}

func (handler *IncomeHandler) DeleteIncome(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	IncomeID := c.Param("id")
	if IncomeID == "" {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "income ID is required", nil))
	}

	IncomeIDUint, err := strconv.Atoi(IncomeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid income ID", nil))
	}

	err = handler.incomeService.DeleteIncome(UserID, uint(IncomeIDUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error deleting income", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success deleting income", nil))
}

func (handler *IncomeHandler) UpdateIncome(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	IncomeID := c.Param("id")
	if IncomeID == "" {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "income ID is required", nil))
	}

	IncomeIDUint, err := strconv.Atoi(IncomeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid income ID", nil))
	}

	var incomeInput IncomeRequest
	if err := c.Bind(&incomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error binding data. Data is not valid.", nil))
	}

	validate := validator.New()
	if err := validate.Struct(incomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	incomeCore := IncomeRequestToCore(incomeInput)
	err = handler.incomeService.UpdateIncome(UserID, uint(IncomeIDUint), incomeCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating income", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success updating income", nil))
}
