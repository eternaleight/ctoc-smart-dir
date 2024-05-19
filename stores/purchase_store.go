package stores

import (
	"github.com/eternaleight/go-backend/models"
	"gorm.io/gorm"
)

// 購入に関連する操作を管理
type PurchaseStore struct {
	db *gorm.DB
}

// 新しいPurchaseStoreのインスタンスを初期化
func NewPurchaseStore(db *gorm.DB) *PurchaseStore {
	return &PurchaseStore{db: db}
}

// 新しい購入をデータベースに保存
func (ps *PurchaseStore) CreatePurchase(purchase *models.Purchase) error {
	return ps.db.Create(purchase).Error
}

// 指定されたIDの購入情報を取得
func (ps *PurchaseStore) GetPurchaseByID(id uint) (*models.Purchase, error) {
	var purchase models.Purchase
	err := ps.db.First(&purchase, id).Error
	return &purchase, err
}
