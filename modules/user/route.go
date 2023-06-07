package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterUser struct {
	UserRequestHandler RequestHandlerUser
}

func NewRouter(
	dbCrud *gorm.DB) RouterUser {
	return RouterUser{UserRequestHandler: NewUserRequestHandler(dbCrud)}
}

func (r RouterUser) Handle(router *gin.Engine) {
	basepath := "/user"
	user := router.Group(basepath)

	user.POST("/",
		r.UserRequestHandler.CreateUser,
	)

	user.GET("/:id",
		r.UserRequestHandler.GetUserById,
	)
	user.PUT("/:id",
		r.UserRequestHandler.UpdateUser,
	)
	user.DELETE("/:email",
		r.UserRequestHandler.DeleteUser,
	)
}
