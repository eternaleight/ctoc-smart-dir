package stores

import (
	"github.com/eternaleight/go-backend/models"
	"gorm.io/gorm"
)

type UserStore struct {
	DB *gorm.DB
}

// 新しいUserStoreを生成
func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{DB: db}
}

// ユーザーをデータベースに保存
func (s *UserStore) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

// IDに基づいてユーザー情報を取得
func (s *UserStore) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// メールアドレスに基づいてユーザー情報を取得
func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
