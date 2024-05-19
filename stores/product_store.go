package stores

import (
	"github.com/eternaleight/go-backend/models"
	"gorm.io/gorm"
)

// 商品に関するデータベース操作を管理
type ProductStore struct {
	db *gorm.DB
}

// 新しいProductStoreを初期化して返す
func NewProductStore(db *gorm.DB) *ProductStore {
	return &ProductStore{db: db}
}

// 新しい商品をデータベースに保存
func (ps *ProductStore) CreateProduct(product *models.Product) error {
	return ps.db.Create(product).Error
}

// 全商品をデータベースから取得
func (ps *ProductStore) ListProducts() ([]models.Product, error) {
	var products []models.Product
	err := ps.db.Find(&products).Error
	return products, err
}

// 指定されたIDの商品をデータベースから取得
func (ps *ProductStore) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := ps.db.First(&product, id).Error
	return &product, err
}

// 指定されたIDの商品情報を更新
func (ps *ProductStore) UpdateProduct(id uint, product *models.Product) error {
	return ps.db.Model(&models.Product{ID: id}).Updates(product).Error
}

// 指定されたIDの商品をデータベースから削除
func (ps *ProductStore) DeleteProduct(id uint) error {
	return ps.db.Delete(&models.Product{}, id).Error
}
