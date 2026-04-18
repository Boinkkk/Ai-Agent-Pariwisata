package routers

import (
	"tutorial/controllers"
	"tutorial/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	UserHandler := controllers.NewUserHandle(userRepo)

	roleRepo := repository.NewRoleRepository(db)
	RoleHandler := controllers.NewRoleHandler(roleRepo)

	r.POST("/users", UserHandler.CreateUser)
	r.GET("/listrole", RoleHandler.GetRole)

	return r
}
