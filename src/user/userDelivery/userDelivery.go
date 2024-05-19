package userDelivery

import (
	"service-code/model/dto/json"
	userEntity "service-code/model/entity/user"
	"service-code/pkg/middleware"
	"service-code/pkg/validation"
	"service-code/src/user"

	"github.com/gin-gonic/gin"
)

type UserDelivery struct {
	userUC user.UserUsecase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUC user.UserUsecase) {
	handler := UserDelivery{
		userUC: userUC,
	}
	userGroup := v1Group.Group("user")
	userGroup.Use(middleware.JWTAuth())
	{
		userGroup.POST("/login", middleware.BasicAuth, handler.LoginHandler)
		userGroup.POST("/create", handler.CreateUserHandler)
		userGroup.GET("", handler.handlerGetAllUsers)
		userGroup.GET("/:id", handler.handleGetUserById)
		userGroup.PUT("/:id", handler.handlerUpdateUser)
		userGroup.DELETE("/:id", handler.handlerDeleteUser)
	}
}

func (u *UserDelivery) LoginHandler(c *gin.Context) {
	var req userEntity.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		json.NewResponseBadRequest(c, validation.GetValidationError(err), "Bad Request", "01", "01")
		return
	}

	token, err := u.userUC.Login(req.Email, req.Password)
	if err != nil {
		json.NewResponseBadRequest(c, validation.GetValidationError(err), "Bad Request", "01", "01")
		return
	}

	json.NewResponseSuccess(c, token, "success", "01", "01")
}

func (c *UserDelivery) handlerGetAllUsers(ctx *gin.Context) {
	version, err := c.userUC.UserList()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, version, "success", "01", "01")
}

func (c *UserDelivery) handleGetUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := c.userUC.UserById(id)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, user, "success", "01", "01")
}

func (c *UserDelivery) handlerUpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user userEntity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		json.NewResponseBadRequest(ctx, validation.GetValidationError(err), "Bad Request", "01", "01")
		return
	}

	if err := c.userUC.UpdateUser(id, user.Fullname, user.Password); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseCreated(ctx, "success", "01", "01")
}

func (c *UserDelivery) CreateUserHandler(ctx *gin.Context) {
	var user userEntity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		json.NewResponseBadRequest(ctx, validation.GetValidationError(err), "Bad Request", "01", "01")
		return
	}

	if err := c.userUC.CreateUser(user.Fullname, user.Email, user.Password); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseCreated(ctx, "success", "01", "01")
}

func (c *UserDelivery) handlerDeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	requestorID := ctx.GetString("userID")

	if err := c.userUC.DeleteUser(id, requestorID); err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseCreated(ctx, "success", "01", "01")
}
