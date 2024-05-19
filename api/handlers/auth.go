package handlers

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/eternaleight/go-backend/stores"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DBのインスタンスを持つ構造体
type Handler struct {
	DB *gorm.DB
}

// 新しいHandlerのインスタンスを初期化
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{DB: db}
}

func getGravatarURL(email string, size int) string {
	emailHash := fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(strings.TrimSpace(email)))))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s?s=%d&d=identicon", emailHash, size)
}

// 新しいユーザーを登録
func (h *Handler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// リクエストからJSONデータをバインド
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// AuthStoreのインスタンスを生成
	authStore := stores.NewAuthStore(h.DB)

	// emailを使ってemailMd5Hashの作成
	emailMd5Hash := fmt.Sprintf("%x", md5.Sum([]byte(input.Email)))

	user, err := authStore.RegisterUser(input.Username, input.Email, input.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	gravatarURL := getGravatarURL(input.Email, 800)
	// JWTトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}
	// JWTトークンをHTTP-only Cookieとして設定（90日の有効期限で設定）
	c.SetCookie("authToken", tokenString, 60*60*24*90, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"user": user, "emailMd5Hash": emailMd5Hash, "gravatarURL": gravatarURL})
}

// ユーザーのログインを処理
func (h *Handler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// リクエストからJSONデータをバインド
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// AuthStoreのインスタンスを生成
	authStore := stores.NewAuthStore(h.DB)

	// メールアドレスを使ってユーザーを取得
	user, err := authStore.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "メールアドレスが存在しません"})
		return
	}

	// ユーザーのパスワードを検証
	err = authStore.ComparePassword(user.Password, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "パスワードが間違っています"})
		return
	}

	// JWTトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "トークンの生成に失敗しました"})
		return
	}
	gravatarURL := getGravatarURL(input.Email, 800)
	// JWTトークンをHTTP-only Cookieとして設定（90日の有効期限で設定）
	c.SetCookie("authToken", tokenString, 60*60*24*90, "/", "", false, true)
	// この設定はセキュアなHTTPS接続では動作しません。実運用では変更してください。6番目の引数（secureフィールド）をfalseにしています。これは、開発環境や非セキュアなHTTP接続で動作する場合の設定です。HTTPS上でアプリケーションを実際に運用する場合、この値をtrueに変更して、Cookieがセキュアな接続でのみ送信されるようにする必要があります。

	c.JSON(http.StatusOK, gin.H{"gravatarURL": gravatarURL})
}
