package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"uptime/backend/controllers"
	"uptime/backend/database"
	"uptime/backend/middleware"
	"uptime/backend/models"
	"uptime/backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

//go:embed frontend_dist/*
var frontendDist embed.FS

func main() {
	// 在开发环境中加载 .env 文件，如果文件不存在则忽略错误
	// 在生产环境中，变量将通过环境变量直接提供
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	database.ConnectDatabase()
	if os.Getenv("REDIS_URI") != "" {
		database.ConnectRedis()
	}

	createAdminUser()

	go services.StartMonitoring()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	r.Use(cors.New(config))

	api := r.Group("/api")
	{
		api.POST("/login", controllers.Login)

		authorized := api.Group("/")
		authorized.Use(middleware.RequireAuth)
		{
			authorized.POST("/monitors", controllers.CreateMonitor)
			authorized.GET("/monitors", controllers.GetMonitors)
			authorized.GET("/monitors/:id", controllers.GetMonitor)
			authorized.PUT("/monitors/:id", controllers.UpdateMonitor)
			authorized.DELETE("/monitors/:id", controllers.DeleteMonitor)
			authorized.POST("/monitors/:id/check", controllers.ManualCheck)
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Serve frontend static files
	dist, err := fs.Sub(frontendDist, "frontend_dist")
	if err != nil {
		log.Fatal("failed to get subtree of embedded files: ", err)
	}
	http.FS(dist)
	r.StaticFS("/_next", http.Dir("frontend_dist/_next"))
	r.StaticFileFS("/", "index.html", http.FS(dist))
	r.NoRoute(func(c *gin.Context) {
		p := c.Request.URL.Path
		// if is a file, return 404
		if path.Ext(p) != "" {
			c.AbortWithStatus(404)
			return
		}
		c.FileFromFS("/", http.FS(dist))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	r.Run(":" + port)
}

func createAdminUser() {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	if username == "" || password == "" {
		log.Println("ADMIN_USERNAME or ADMIN_PASSWORD not set, skipping admin user creation.")
		return
	}

	var user models.User
	database.DB.First(&user, "username = ?", username)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}

	if user.ID == 0 {
		// Create new admin user
		admin := models.User{Username: username, Password: string(hashedPassword)}
		if err := database.DB.Create(&admin).Error; err != nil {
			log.Fatalf("Failed to create admin user: %v", err)
		}
		log.Println("Admin user created successfully.")
	} else {
		// Update existing admin user's password
		if err := database.DB.Model(&user).Update("password", string(hashedPassword)).Error; err != nil {
			log.Fatalf("Failed to update admin user: %v", err)
		}
		log.Println("Admin user password updated successfully.")
	}
}