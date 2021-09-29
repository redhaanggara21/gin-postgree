package handlerAuth

import (
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

		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := util.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.service.Login(&payload)

	if err != nil {
		//fmt.Println("handler 2")
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := util.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := util.APIResponse("Successfuly loggedin", http.StatusOK, "success", data)

	ctx.JSON(http.StatusOK, response)
}

func (h *handler) RegisterHandler(ctx *gin.Context) {

	var payload Auth.RegisterRequest
	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := util.APIResponse("Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data, err := h.service.Register(&payload)
	if err != nil {
		errors := util.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := util.APIResponse("Register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := util.APIResponse("Successfuly register", http.StatusOK, "success", data)

	ctx.JSON(http.StatusOK, response)

}
