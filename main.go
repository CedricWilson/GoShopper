package main

import (
	config "main/config"
	controller "main/controllers"
	"main/di"
	"main/repository"
	"main/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.InitDB()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	// bookRepository repository.BookRepository = repository.NewBookRepository(db)
	// jwtService     service.JWTservice        = service.NewJWTservice()
	// userService    service.UserService       = service.NewUserService(userRepository)
	// bookService    service.BookService       = service.NewBookService(bookRepository)
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService)
	// bookController controller.BookController = controller.NewBookController(bookService, jwtService)
	// userController controller.UserController = controller.NewUserController(userService, jwtService)
)

func main() {
	router := gin.Default()
	// di.InitDB()

	di.InitRedis()
	

	//
	router.LoadHTMLFiles("static/index.tmpl.html")
	
	router.Static("/css", "static/css")
	router.Static("/images", "static/images")


	
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"title": "GoShopper",
		})
	})
	authRoutes := router.Group("v2")
	{
		// authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	controller.UserApi(router)

	controller.ProductApi(router)
	controller.CartApi(router)

	//
	router.Run(":8080")
}

