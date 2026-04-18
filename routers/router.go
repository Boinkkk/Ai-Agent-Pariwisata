package routers

import (
	"tutorial/handlers"
	"tutorial/repository"
	"tutorial/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	UserRepo := repository.NewUserRepository(db)
	UserService := service.NewUserService(UserRepo)
	UserHandler := handlers.NewUserHandler(UserService)

	roleRepo := repository.NewRoleRepository(db)
	RoleHandler := handlers.NewRoleHandler(roleRepo)

	r.POST("/users", UserHandler.CreateUser)
	r.GET("/listrole", RoleHandler.GetRole)

	return r
}
