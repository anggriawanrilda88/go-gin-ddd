package routes

import (
	"pos/infrastructure/auth"
	"pos/infrastructure/persistence"
	"pos/interfaces/middleware"
	handler "pos/interfaces/rest/handler/website"

	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine, tk *auth.Token, services *persistence.Repositories, redisService *auth.RedisService) {
	WebsiteRoute(r, tk, services, redisService)
}

// SellerCenterRouter a completely separate router for sinbad seller center routes
func WebsiteRoute(r *gin.Engine, tk *auth.Token, services *persistence.Repositories, redisService *auth.RedisService) {
	r.Use(middleware.CORSMiddleware()) //For CORS

	//set repository
	usersHandler := handler.NewUsers(services.User, redisService.Auth, tk)

	// users routes
	r.GET("/website/api/v1/users", usersHandler.GetUsers)
}

//// example route gin
// users := interfaces.NewUsers(services.User, redisService.Auth, tk)
// foods := interfaces.NewFood(services.Food, services.User, fd, redisService.Auth, tk)
// authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

// r.Use(middleware.CORSMiddleware()) //For CORS

// //user routes
// r.POST("/users", users.SaveUser)
// r.GET("/users", users.GetUsers)
// r.GET("/users/:user_id", users.GetUser)

// //website user routes
// r.GET("/website/api/v1/users", customer.GetUsers)

// //post routes
// r.POST("/food", middleware.AuthMiddleware(), middleware.MaxSizeAllowed(8192000), foods.SaveFood)
// r.PUT("/food/:food_id", middleware.AuthMiddleware(), middleware.MaxSizeAllowed(8192000), foods.UpdateFood)
// r.GET("/food/:food_id", foods.GetFoodAndCreator)
// r.DELETE("/food/:food_id", middleware.AuthMiddleware(), foods.DeleteFood)
// r.GET("/food", foods.GetAllFood)

// //authentication routes
// r.POST("/login", authenticate.Login)
// r.POST("/logout", authenticate.Logout)
// r.POST("/refresh", authenticate.Refresh)

//Starting the application
