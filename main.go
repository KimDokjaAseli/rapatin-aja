package main

import (
	"os"
	"rapatln_backend/config"
	"rapatln_backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORS configuration
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Admin-ID"}
	r.Use(cors.New(configCors))

	// Connect to database
	config.ConnectDatabase()

	// Routes
	api := r.Group("/api")
	{
		// Auth
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Meetings
		api.GET("/meetings", controllers.GetMeetings)
		api.GET("/meetings/:id", controllers.GetMeeting)
		api.POST("/meetings", controllers.CreateMeeting)
		api.PATCH("/meetings/:id", controllers.UpdateMeeting)
		api.DELETE("/meetings/:id", controllers.DeleteMeeting)

		// User
		api.GET("/user/:id", controllers.GetUser)
		api.PUT("/user/:id", controllers.UpdateUser)

		// Admin
		api.GET("/admin/stats", controllers.GetAdminStats)
		api.GET("/admin/users", controllers.AdminGetAllUsers)
		api.PUT("/admin/users/:id/role", controllers.AdminUpdateUserRole)
		api.GET("/admin/meetings", controllers.AdminGetAllMeetings)
		api.DELETE("/admin/meetings/:id", controllers.AdminDeleteMeeting)
	}

	// Serve Frontend Static Files
	r.StaticFile("/", "./fe/index.html")
	r.StaticFile("/app.js", "./fe/app.js")
	r.StaticFile("/style.css", "./fe/style.css")

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
