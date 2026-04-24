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

	CategoriesRepo := repository.NewCategoriesRepository(db)
	CategoriesHandler := handlers.NewCategoriesHandler(CategoriesRepo)

	AuthRepo := repository.NewAuthRepository(db)
	AuthService := service.NewAuthService(AuthRepo)
	AuthHandler := handlers.NewAuthHandler(AuthService)

	r.POST("/users", UserHandler.CreateUser)
	r.GET("/listrole", RoleHandler.GetRole)
	r.POST("/login", AuthHandler.Login)
	r.GET("/categories", CategoriesHandler.GetCategories)

	return r
}
