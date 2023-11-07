package handler

import (
	"net/http"
	"strconv"
	"strings"
	"technopartner/app/middlewares"
	"technopartner/features/outcomes"
	"technopartner/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type OutcomeHandler struct {
	outcomeService outcomes.OutcomeServiceInterface
}

func New(service outcomes.OutcomeServiceInterface) *OutcomeHandler {
	return &OutcomeHandler{
		outcomeService: service,
	}
}

func (handler *OutcomeHandler) CreateOutcome(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	var outcomeInput OutcomeRequest
	if err := c.Bind(&outcomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error binding data. Data is not valid.", nil))
	}

	validate := validator.New()
	if err := validate.Struct(outcomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	outcomeCore := OutcomeRequestToCore(outcomeInput)
	err := handler.outcomeService.CreateOutcome(UserID, outcomeCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error inserting data", nil))
		}
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success inserting data", nil))
}

func (handler *OutcomeHandler) GetAllOutcomes(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	result, err := handler.outcomeService.GetAllOutcomes(UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error retrieving outcomes", nil))
	}
	var outcomeResponse []OutcomeResponseAll
	for _, value := range result {
		outcomeResponse = append(outcomeResponse, OutcomeResponseAll{
			Name:   value.Name,
			Jumlah: value.Jumlah,
		})

	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success retrieving outcomes", outcomeResponse))
}

func (handler *OutcomeHandler) DeleteOutcome(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	OutcomeID := c.Param("id")
	if OutcomeID == "" {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "outcome ID is required", nil))
	}

	OutcomeIDUint, err := strconv.Atoi(OutcomeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid outcome ID", nil))
	}

	err = handler.outcomeService.DeleteOutcome(UserID, uint(OutcomeIDUint))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error deleting outcome", nil))
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success deleting outcome", nil))
}

func (handler *OutcomeHandler) UpdateOutcome(c echo.Context) error {
	UserID, er := middlewares.ExtractTokenUser(c)
	if er != nil {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, er.Error(), nil))
	}

	OutcomeID := c.Param("id")
	if OutcomeID == "" {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "outcome ID is required", nil))
	}

	OutcomeIDUint, err := strconv.Atoi(OutcomeID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "invalid outcome ID", nil))
	}

	var outcomeInput OutcomeRequest
	if err := c.Bind(&outcomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error binding data. Data is not valid.", nil))
	}

	validate := validator.New()
	if err := validate.Struct(outcomeInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	outcomeCore := OutcomeRequestToCore(outcomeInput)
	err = handler.outcomeService.UpdateOutcome(UserID, uint(OutcomeIDUint), outcomeCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating outcome", nil))
		}
	}

	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success updating outcome", nil))
}
