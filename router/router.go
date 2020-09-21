package router

import (
	"../pkg/modules"
	"../pkg/modules/images"
	"../pkg/modules/migrate"
	"../pkg/scraping"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// APIRouter func
func APIRouter() {

	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	// api group
	api := router.Group("/api")
	{
		// 全レコード取得 cities
		api.GET("/city", modules.GetCities)
		//　１件取得 city
		api.GET("/city/:name", modules.GetCity)
		// 全レコード取得 countries
		api.GET("/country", modules.GetCountries)
		//　１件取得 country
		api.GET("/country/:name", modules.GetCountry)
		// 大陸一覧
		api.GET("/continent", modules.GetContinentsList)
		// スクレイピング
		api.GET("/scraping", scraping.GetScrape)

		preImg := images.PreImageController{}

		// Preuploadされた一覧を取得
		api.GET("/image/pre_upload", preImg.GetAll)
		// Preuploadする
		api.POST("/image/pre_upload", preImg.Upload)

		//　Preuploadのでデータを一枚削除
		api.DELETE("/image/pre_upload/delete", preImg.Delete)
		//　本アップロード
		api.PUT("/image/upload", preImg.ComitUpload)

		upedImg := images.ImageController{} // 画像一枚だけ取得して、Base64で返す。
		api.GET("/image/upload/:id", upedImg.GetFile)
		api.GET("/image/upload", upedImg.GetAll)

		// Delete All images
		api.DELETE("/image/delete/:flag", images.DeleteAllImages)

	}

	router.GET("/migrate", migrate.Migrate)

	router.Run(":8000")
}
