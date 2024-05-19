package models

import "time"

type User struct {
	ID           uint          `gorm:"primaryKey" json:"id"`
	Username     string        `gorm:"column:username" json:"username"`
	Email        string        `gorm:"unique;index;column:email" json:"email"`
	EmailMd5Hash string        `gorm:"column:emailMd5Hash;unique;index" json:"emailMd5Hash"`
	Password     string        `gorm:"column:password" json:"password"`
	Posts        []Post        `gorm:"foreignKey:AuthorID" json:"posts"`
	Profile      Profile       `gorm:"foreignKey:UserID" json:"profile"`
	Products     []Product     `gorm:"foreignKey:SellerID" json:"products"`
	Purchases    []Purchase    `gorm:"foreignKey:BuyerID" json:"purchases"`
	Favorites    []Favorite    `gorm:"foreignKey:UserID" json:"favorites"`
	RequestCards []RequestCard `gorm:"foreignKey:RequesterID" json:"requestCards"`
}

func (User) TableName() string {
	return "User"
}

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"column:content" json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;column:createdAt" json:"createdAt"`
	AuthorID  uint      `gorm:"column:authorId;index" json:"authorId"`
}

func (Post) TableName() string {
	return "Post"
}

type Profile struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Bio             string `gorm:"size:1000;column:bio" json:"bio"`
	ProfileImageUrl string `gorm:"column:profileImageUrl" json:"profileImageUrl"`
	UserID          uint   `gorm:"column:userId;index;unique" json:"userId"`
}

func (Profile) TableName() string {
	return "Profile"
}

type Product struct {
	ID          uint       `gorm:"primaryKey;column:id" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Price       int        `gorm:"column:price" json:"price"`
	ImageURL    string     `gorm:"column:imageUrl" json:"imageUrl"`
	ImageSize   uint64     `gorm:"column:imageSize" json:"imageSize"`
	VideoURL    string     `gorm:"column:videoUrl" json:"videoUrl"`
	CreatedAt   time.Time  `gorm:"column:createdAt" json:"createdAt"`
	SellerID    uint       `gorm:"column:sellerId;index" json:"sellerId"`
	Purchases   []Purchase `gorm:"foreignKey:ProductID" json:"purchases"`
}

func (Product) TableName() string {
	return "Product"
}

type Purchase struct {
	ID              uint      `gorm:"primaryKey;column:id" json:"id"`
	Price           int       `gorm:"column:price" json:"price"`
	ProductID       uint      `gorm:"column:productId;index" json:"productId"`
	BuyerID         uint      `gorm:"column:buyerId;index" json:"buyerId"`
	PurchaseDate    time.Time `gorm:"column:purchaseDate" json:"purchaseDate"`
	StripePaymentID string    `gorm:"column:stripePaymentId" json:"stripePaymentId"`
}

func (Purchase) TableName() string {
	return "Purchase"
}

type Image struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	URL       string    `gorm:"column:url" json:"url"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;column:createdAt" json:"createdAt"`
}

func (Image) TableName() string {
	return "Image"
}

type Favorite struct {
	ID        uint `gorm:"primaryKey;column:id" json:"id"`
	UserID    uint `gorm:"column:userId;index" json:"userId" gorm:"foreignKey:UserID"`
	ProductID uint `gorm:"column:productId;index" json:"productId" gorm:"foreignKey:ProductID"`
}

func (Favorite) TableName() string {
	return "Favorite"
}

type RequestCard struct {
	ID          uint   `gorm:"primaryKey;column:id" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	RequesterID uint   `gorm:"column:requesterId;index" json:"requesterId" gorm:"foreignKey:RequesterID"`
}

func (RequestCard) TableName() string {
	return "RequestCard"
}
