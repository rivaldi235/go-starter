package router

import (
	"database/sql"
	"service-code/src/user/userDelivery"
	"service-code/src/user/userRepository"
	"service-code/src/user/userUsecase"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	userRepo := userRepository.NewUserRepository(db)
	userUC := userUsecase.NewUserUsecase(userRepo)
	userDelivery.NewUserDelivery(v1Group, userUC)
}
