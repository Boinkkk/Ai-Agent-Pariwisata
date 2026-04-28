package routers

import (
	"tutorial/handlers"
	"tutorial/repository"
	"tutorial/service"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type crudHandler interface {
	Create(c *gin.Context)
	FindByID(c *gin.Context)
	FindAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

func registerCRUDRoutes(r gin.IRouter, path string, handler crudHandler) {
	r.GET(path, handler.FindAll)
	r.GET(path+"/:id", handler.FindByID)
	r.POST(path, handler.Create)
	r.PATCH(path+"/:id", handler.Update)
	r.DELETE(path+"/:id", handler.Delete)
}

func SetupRouter(db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	roleRepo := repository.NewRoleRepository(db)
	roleService := service.NewRoleService(roleRepo)
	roleHandler := handlers.NewRoleHandler(roleService)

	categoriesRepo := repository.NewCategoriesRepository(db)
	categoriesService := service.NewCategoriesService(categoriesRepo)
	categoriesHandler := handlers.NewCategoriesHandler(categoriesService)

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	addressesHandler := handlers.NewAddressesHandler(service.NewAddressesService(repository.NewAddressesRepository(db)))
	articleHandler := handlers.NewArticleHandler(service.NewArticleService(repository.NewArticleRepository(db)))
	cartHandler := handlers.NewCartHandler(service.NewCartService(repository.NewCartRepository(db)))
	cartItemHandler := handlers.NewCartItemHandler(service.NewCartItemService(repository.NewCartItemRepository(db)))
	chatSessionHandler := handlers.NewChatSessionHandler(service.NewChatSessionService(repository.NewChatSessionRepository(db)))
	chatMessageHandler := handlers.NewChatMessageHandler(service.NewChatMessageService(repository.NewChatMessageRepository(db)))
	courierHandler := handlers.NewCourierHandler(service.NewCourierService(repository.NewCourierRepository(db)))
	knowledgeBaseHandler := handlers.NewKnowledgeBaseHandler(service.NewKnowledgeBaseService(repository.NewKnowledgeBaseRepository(db)))
	orderHandler := handlers.NewOrderHandler(service.NewOrderService(repository.NewOrderRepository(db)))
	orderItemHandler := handlers.NewOrderItemHandler(service.NewOrderItemService(repository.NewOrderItemRepository(db)))
	orderStatusHistoryHandler := handlers.NewOrderStatusHistoryHandler(service.NewOrderStatusHistoryService(repository.NewOrderStatusHistoryRepository(db)))
	paymentHandler := handlers.NewPaymentHandler(service.NewPaymentService(repository.NewPaymentRepository(db)))
	productReviewHandler := handlers.NewProductReviewHandler(service.NewProductReviewService(repository.NewProductReviewRepository(db)))

	r.POST("/users", userHandler.CreateUser)
	r.GET("/listrole", roleHandler.GetRole)
	r.POST("/login", authHandler.Login)
	r.GET("/categories", categoriesHandler.GetCategories)
	r.POST("/categories/add", categoriesHandler.AddCategories)
	r.GET("/delete/categories/:id", categoriesHandler.DeleteCategoriesByID)
	r.GET("/api/v1/products/slug/:slug", productHandler.GetProductBySlug)

	api := r.Group("/api/v1")
	registerCRUDRoutes(api, "/users", userHandler)
	registerCRUDRoutes(api, "/roles", roleHandler)
	registerCRUDRoutes(api, "/categories", categoriesHandler)
	registerCRUDRoutes(api, "/products", productHandler)
	registerCRUDRoutes(api, "/addresses", addressesHandler)
	registerCRUDRoutes(api, "/articles", articleHandler)
	registerCRUDRoutes(api, "/carts", cartHandler)
	registerCRUDRoutes(api, "/cart-items", cartItemHandler)
	registerCRUDRoutes(api, "/chat-sessions", chatSessionHandler)
	registerCRUDRoutes(api, "/chat-messages", chatMessageHandler)
	registerCRUDRoutes(api, "/couriers", courierHandler)
	registerCRUDRoutes(api, "/knowledge-bases", knowledgeBaseHandler)
	registerCRUDRoutes(api, "/orders", orderHandler)
	registerCRUDRoutes(api, "/order-items", orderItemHandler)
	registerCRUDRoutes(api, "/order-status-histories", orderStatusHistoryHandler)
	registerCRUDRoutes(api, "/payments", paymentHandler)
	registerCRUDRoutes(api, "/product-reviews", productReviewHandler)

	return r
}
