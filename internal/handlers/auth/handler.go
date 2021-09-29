package handlerAuth

import (
	"errors"
	"net/http"
	util "pelatihan-be/helpers/utils"
	Auth "pelatihan-be/internal/controllers/auth"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service Auth.Service
}

func NewHandler(service Auth.Service) *handler {
	return &handler{service: service}
}

func (h *handler) LoginHandler(ctx *gin.Context) {

	var payload Auth.LoginRequest
	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		response := util.APIResponseFailed(err, http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// validate
	validate := payload.ValidatorLogin()
	if validate != "" {
		response := util.APIResponseFailed(errors.New(validate), http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.service.Login(&payload)
	if err != nil {
		response := util.APIResponseFailed(err, http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	response := util.APIResponse("Successfuly loggedin", http.StatusOK, true, data)
	ctx.JSON(http.StatusOK, response)
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	var payload Auth.RegisterRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		response := util.APIResponseFailed(err, http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	validate := payload.ValidatorRegister()
	if validate != "" {
		response := util.APIResponseFailed(errors.New(validate), http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.service.Register(&payload)
	if err != nil {
		response := util.APIResponseFailed(err, http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := util.APIResponse("Successfuly register", http.StatusOK, true, data)
	ctx.JSON(http.StatusOK, response)
}
