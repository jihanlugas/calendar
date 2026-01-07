package propertyprice

import (
	"net/http"
	"strings"

	"github.com/jihanlugas/calendar/jwt"
	"github.com/jihanlugas/calendar/request"
	"github.com/jihanlugas/calendar/response"
	"github.com/jihanlugas/calendar/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	usecase Usecase
}

func NewHandler(usecase Usecase) Handler {
	return Handler{
		usecase: usecase,
	}
}

// GetById
// @Tags Propertyprice
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Propertyprice ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /propertyprice/{id} [get]
func (h Handler) GetById(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetUserInfo, err, nil).SendJSON(c)
	}

	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetParam, err, nil).SendJSON(c)
	}

	preloadSlice := []string{}
	preloads := c.QueryParam("preloads")
	if preloads != "" {
		preloadSlice = strings.Split(preloads, ",")
	}

	vPropertyprice, err := h.usecase.GetById(loginUser, id, preloadSlice...)
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error(), err, nil).SendJSON(c)
	}

	return response.Success(http.StatusOK, "Successfully retrieved propertyprice", vPropertyprice).SendJSON(c)
}

// Create
// @Tags Propertyprice
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param req body request.CreatePropertyprice true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /propertyprice [post]
func (h Handler) Create(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetUserInfo, err, nil).SendJSON(c)
	}

	req := new(request.CreatePropertyprice)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerBind, err, nil).SendJSON(c)
	}

	utils.TrimWhitespace(req)

	err = c.Validate(req)
	if err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerFailedValidation, err, response.ValidationError(err)).SendJSON(c)
	}

	if jwt.IsSaveCompanyIDOR(loginUser, req.CompanyID) {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerIDOR, err, nil).SendJSON(c)
	}

	err = h.usecase.Create(loginUser, *req)
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error(), err, nil).SendJSON(c)
	}

	return response.Success(http.StatusOK, "Successfully created propertyprice", nil).SendJSON(c)
}

// Update
// @Tags Propertyprice
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Propertyprice ID"
// @Param req body request.UpdatePropertyprice true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /propertyprice/{id} [put]
func (h Handler) Update(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetUserInfo, err, nil).SendJSON(c)
	}

	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetParam, err, nil).SendJSON(c)
	}

	req := new(request.UpdatePropertyprice)
	if err = c.Bind(req); err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerBind, err, nil).SendJSON(c)
	}

	utils.TrimWhitespace(req)

	err = c.Validate(req)
	if err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerFailedValidation, err, response.ValidationError(err)).SendJSON(c)
	}

	err = h.usecase.Update(loginUser, id, *req)
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error(), err, nil).SendJSON(c)
	}

	return response.Success(http.StatusOK, "Successfully updated propertyprice", nil).SendJSON(c)
}

// Delete
// @Tags Propertyprice
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Propertyprice ID"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /propertyprice/{id} [delete]
func (h Handler) Delete(c echo.Context) error {
	var err error

	loginUser, err := jwt.GetUserLoginInfo(c)
	if err != nil {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetUserInfo, err, nil).SendJSON(c)
	}

	id := strings.TrimSpace(c.Param("id"))
	if id == "" {
		return response.Error(http.StatusBadRequest, response.ErrorHandlerGetParam, err, nil).SendJSON(c)
	}

	err = h.usecase.Delete(loginUser, id)
	if err != nil {
		return response.Error(http.StatusInternalServerError, err.Error(), err, nil).SendJSON(c)
	}

	return response.Success(http.StatusOK, "Successfully deleted propertyprice", nil).SendJSON(c)
}
