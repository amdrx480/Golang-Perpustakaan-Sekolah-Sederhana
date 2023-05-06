package routes

import (
	"os"
	"perpustakaan/controllers"
	"perpustakaan/middlewares"

	echojwt "github.com/labstack/echo-jwt/v4"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {

	//implement logger middleware
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       os.Getenv("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	authMiddlewareConfig := jwtConfig.Init()

	//admin
	adminsController := controllers.IniAdminController(&jwtConfig)

	admins := e.Group("api/v1/admins")

	admins.POST("/register", adminsController.Register)
	admins.POST("/login", adminsController.Login)

	//books
	bookController := controllers.InitBookController()

	bookRoutes := e.Group("api/v1", echojwt.WithConfig(authMiddlewareConfig))
	bookRoutes.Use(middlewares.VerifyTooken)

	bookRoutes.GET("/books", bookController.GetAll)
	//id menggunakan param("id")
	bookRoutes.GET("/books/:id", bookController.GetByID)
	bookRoutes.POST("/books", bookController.Create)
	bookRoutes.PUT("/books/:id", bookController.Update)
	bookRoutes.DELETE("/books/:id", bookController.Delete)
	bookRoutes.POST("/books/:id", bookController.Restore)
	bookRoutes.DELETE("/books/:id/force", bookController.ForceDelete)

	//members
	memberController := controllers.InitMemberkController()

	memberRoutes := e.Group("api/v1", echojwt.WithConfig(authMiddlewareConfig))
	memberRoutes.Use(middlewares.VerifyTooken)

	memberRoutes.GET("/members", memberController.GetAll)
	//id menggunakan param("id")
	memberRoutes.GET("/members/:id", memberController.GetByID)
	memberRoutes.POST("/members", memberController.Create)
	memberRoutes.PUT("/members/:id", memberController.Update)
	memberRoutes.DELETE("/members/:id", memberController.Delete)
	memberRoutes.POST("/members/:id", memberController.Restore)
	memberRoutes.DELETE("/members/:id/force", memberController.ForceDelete)

	//transactions
	transactionController := controllers.InitTransactionController()

	transactionRoutes := e.Group("api/v1", echojwt.WithConfig(authMiddlewareConfig))
	transactionRoutes.Use(middlewares.VerifyTooken)

	memberRoutes.GET("/transactions", transactionController.GetAll)
	//id menggunakan param("id")
	transactionRoutes.GET("/transactions/:id", transactionController.GetByID)
	transactionRoutes.POST("/transactions", transactionController.Create)
	transactionRoutes.PUT("/transactions/:id", transactionController.Update)
	transactionRoutes.DELETE("/transactions/:id", transactionController.Delete)
	transactionRoutes.POST("/transactions/:id", transactionController.Restore)
	transactionRoutes.DELETE("/transactions/:id/force", transactionController.ForceDelete)
}
