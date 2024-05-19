package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eternaleight/go-backend/models"
	"github.com/eternaleight/go-backend/stores"
)

// 新しい投稿を作成する
func (h *Handler) CreatePost(c *gin.Context) {
	var input struct {
		Content string `json:"content"`
	}

	// リクエストからJSONデータをバインド
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 投稿内容が空の場合のエラーチェック
	if input.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "投稿内容がありません"})
		return
	}

	// isAuthenticatedミドルウェアで設定されたuserIDを取得
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDが見つかりません"})
		return
	}

	// userIDの型確認
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDの型が正しくありません"})
		return
	}

	// PostStoreのインスタンスを生成
	postStore := stores.NewPostStore(h.DB)
	post := models.Post{
		Content:  input.Content,
		AuthorID: userID,
	}
	// 投稿データをデータベースに保存
	if err := postStore.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

// 最新の投稿を取得
func (h *Handler) GetLatestPosts(c *gin.Context) {
	var posts []models.Post

	// PostStoreのインスタンスを生成
	postStore := stores.NewPostStore(h.DB)

	// 最新の投稿をデータベースから取得
	posts, err := postStore.GetLatestPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
