package stores

import (
	"github.com/eternaleight/go-backend/models"
	"gorm.io/gorm"
)

type PostStore struct {
	DB *gorm.DB
}

// 新しいPostStoreを生成
func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{DB: db}
}

// 投稿をデータベースに保存
func (s *PostStore) CreatePost(post *models.Post) error {
	return s.DB.Create(post).Error
}

// 最新の投稿を取得
func (s *PostStore) GetLatestPosts() ([]models.Post, error) {
	var posts []models.Post
	// 投稿日時の降順に10件の投稿を取得し、それらの投稿者も同時に取得
	if err := s.DB.Order("created_at desc").Limit(10).Preload("Author").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
