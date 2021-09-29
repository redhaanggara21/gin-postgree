package routes

import (
	controllerAuth "pelatihan-be/internal/controllers/auth"
	handlerAuth "pelatihan-be/internal/handlers/auth"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {

	repository := controllerAuth.NewRepository(db)
	service := controllerAuth.NewService(repository)
	handler := handlerAuth.NewHandler(service)

	gRoute := route.Group("/api/v1")
	gRoute.POST("/login", handler.LoginHandler)
	gRoute.POST("/register", handler.RegisterHandler)

}
