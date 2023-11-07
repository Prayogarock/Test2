package handler

import (
	"net/http"
	"strings"
	"technopartner/features/users"
	"technopartner/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	var userInput LoginRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	var response = LoginResponse{
		ID:    dataLogin.ID,
		Name:  dataLogin.Name,
		Token: token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {

	var userInput UserRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	userCore := UserRequestToCore(userInput)
	dataCreate, token, err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	var response = CreateResponse{
		ID:    dataCreate.ID,
		Name:  dataCreate.Name,
		Token: token,
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", response))
}
