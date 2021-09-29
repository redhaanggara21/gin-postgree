package handlerToken

import (
	"errors"
	"net/http"
	util "pelatihan-be/helpers/utils"
	controllerToken "pelatihan-be/internal/controllers/token"

	"github.com/gin-gonic/gin"
)

type handler struct {
	service controllerToken.Service
}

func NewHandler(service controllerToken.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateToken(ctx *gin.Context) {

	var payload controllerToken.CreateRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		response := util.APIResponseFailed(err, http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	validate := payload.ValidatorCreate()
	if validate != "" {
		response := util.APIResponseFailed(errors.New(validate), http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.service.Create(&payload)
	if err != nil {
		response := util.APIResponseFailed(err, http.StatusUnprocessableEntity, false, nil)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := util.APIResponse("Berhasil generate token", http.StatusOK, true, data)
	ctx.JSON(http.StatusOK, response)
}
