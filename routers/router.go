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

	//Produk

	ProductRepo := repository.NewProductRepository(db)
	ProductHandler := handlers.NewProductHandler(ProductRepo)

	r.POST("/users", UserHandler.CreateUser)
	r.GET("/listrole", RoleHandler.GetRole)
	r.POST("/login", AuthHandler.Login)
	r.GET("/categories", CategoriesHandler.GetCategories)
	r.POST("/categories/add", CategoriesHandler.AddCategories)
	r.GET("/delete/categories/:id", CategoriesHandler.DeleteCategoriesByID)

	// Produk
	r.GET("/api/v1/products", ProductHandler.GetAllProduct)
	r.GET("/api/v1/products/:id", ProductHandler.GetProductByID)
	r.POST("/api/v1/products", ProductHandler.AddProduct)
	r.PATCH("/api/v1/products/:id", ProductHandler.UpdateProductByID)
	r.DELETE("api/v1/products/:id", ProductHandler.DeleteProductByID)

	return r
}
