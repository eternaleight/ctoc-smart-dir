package stores

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/eternaleight/go-backend/models"
)

type AuthStore struct {
	DB *gorm.DB
}

// 新しいAuthStoreを生成
func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{DB: db}
}

// ユーザーを登録
func (s *AuthStore) RegisterUser(username, email, password string) (*models.User, error) {
	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	// ユーザー情報をデータベースに保存
	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	result := s.DB.Create(user)
	return user, result.Error
}

// メールアドレスに基づいてユーザー情報を取得
func (s *AuthStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

// ハッシュ化されたパスワードと平文のパスワードを比較
func (s *AuthStore) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
