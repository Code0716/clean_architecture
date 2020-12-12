package infrastructure

import (
	"github.com/Code0716/clean_architecture/app/api/interfaces/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8000"}
	router.Use(cors.New(config))

	userController := controllers.NewUserController(ConnectMySQL())
	preImagesController := controllers.NewPreImagesController(ConnectMySQL())
	imagesController := controllers.NewImagesController(ConnectMySQL())

	// api group
	api := router.Group("/api/v1")
	{
		// login
		api.POST("/login", func(c *gin.Context) { userController.Login(c, passwordVerify, getNewToken) })

		api.GET("/users", func(c *gin.Context) { validateJWT(c, userController.Index) })
		api.GET("/users/:id", func(c *gin.Context) { validateJWT(c, userController.Show) })
		api.POST("/users", func(c *gin.Context) { userController.Create(c, GetUuid(), time.Now(), passwordHash, getNewToken) })
		api.DELETE("/users/:id", func(c *gin.Context) { userController.LogicalDelete(c, time.Now()) })

		// 画像
		api.GET("/image/pre_upload", func(c *gin.Context) { validateJWT(c, preImagesController.GetAll) }) // Preuploadされた一覧を取得
		api.GET("/image/upload", func(c *gin.Context) { validateJWT(c, imagesController.GetAll) })        // uploadされた一覧を取得

	}

	//router.GET("/migrate", migrate.Migrate)

	router.Run(":8000")
}
