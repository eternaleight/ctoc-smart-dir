package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/eternaleight/go-backend/api/handlers"
	"github.com/eternaleight/go-backend/api/middlewares"
	"github.com/eternaleight/go-backend/stores"
)

func setupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// トレーリングスラッシュへのリダイレクトを無効にする
	r.RedirectTrailingSlash = false

	config := cors.DefaultConfig()
	config.AllowCredentials = true

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS") // 環境変数から読み取る
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000" // デフォルト値
	}
	config.AllowOrigins = []string{allowedOrigins} // フロントエンドのオリジンに合わせて変更

	r.Use(cors.New(config))

	// 'Authorization'ヘッダーを許可するためにヘッダーを追加
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	// 新しいハンドラのインスタンスを作成し、データベースを渡す
	productStore := stores.NewProductStore(db)
	purchaseStore := stores.NewPurchaseStore(db)

	handler := handlers.NewHandler(db)
	productHandler := handlers.NewProductHandler(productStore)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseStore)

	// auth
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handler.Register)
		auth.POST("/login", handler.Login)
	}

	// posts
	posts := r.Group("/api/posts").Use(middlewares.IsAuthenticated())
	{
		posts.POST("", handler.CreatePost)
		posts.GET("", handler.GetLatestPosts)
	}

	// user
	user := r.Group("/api/user").Use(middlewares.IsAuthenticated())
	user.GET("", handler.GetUser)

	// products
	products := r.Group("/api/products").Use(middlewares.IsAuthenticated())
	{
		products.POST("", productHandler.CreateProduct)
		products.GET("", productHandler.ListProducts)
		products.GET("/:id", productHandler.GetProductByID)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}

	// purchase
	purchase := r.Group("/api/purchase").Use(middlewares.IsAuthenticated())
	purchase.POST("", purchaseHandler.CreatePurchase)

	return r
}
