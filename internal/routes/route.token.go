package routes

import (
	controllerToken "pelatihan-be/internal/controllers/token"
	handlertoken "pelatihan-be/internal/handlers/token"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitTokenRoutes(db *gorm.DB, route *gin.Engine) {

	repository := controllerToken.NewRepository(db)
	service := controllerToken.NewService(repository)
	handler := handlertoken.NewHandler(service)

	gRoute := route.Group("/api/v1/token")
	gRoute.POST("/create", handler.CreateToken)

}
