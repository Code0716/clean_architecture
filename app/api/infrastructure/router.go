package infrastructure

import (
	"time"

	"github.com/Code0716/clean_architecture/app/api/interfaces/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	router.Use(cors.New(config))

	// api group
	api := router.Group("/api/v1")
	{

		userController := controllers.NewUserController(ConnectMySQL())
		api.GET("/users", func(c *gin.Context) { userController.Index(c) })
		api.GET("/users/:id", func(c *gin.Context) { userController.Show(c) })
		uuid, _ := GetUuid()
		api.POST("/users", func(c *gin.Context) { userController.Create(c, uuid, time.Now()) })

		preImagesController := controllers.NewPreImagesController(ConnectMySQL())
		api.GET("/image/pre_upload", func(c *gin.Context) { preImagesController.GetAll(c) }) // Preuploadされた一覧を取得
		imagesController := controllers.NewImagesController(ConnectMySQL())
		api.GET("/image/upload", func(c *gin.Context) { imagesController.GetAll(c) }) // uploadされた一覧を取得

	}

	//router.GET("/migrate", migrate.Migrate)

	router.Run(":8000")
}
