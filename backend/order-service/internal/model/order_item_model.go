package model

import (
	"time"

	"github.com/google/uuid"
)

type OrderItem struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	OrderID uuid.UUID `gorm:"type:uuid;not null;index" json:"order_id"`

	ProductID uuid.UUID `gorm:"type:varchar(100);not null" json:"product_id"`

	ProductNameSnapshot      string `gorm:"type:varchar(255);not null" json:"product_name_snapshot"`
	ProductPriceSnapshot     int64  `gorm:"not null" json:"product_price_snapshot"`
	ProductThumbnailSnapshot string `gorm:"type:varchar(255);not null" json:"product_thumbnail_snapshot"`
	ProductWeightSnapshot    int32  `gorm:"not null" json:"product_weight_snapshot"`

	Quantity int32 `gorm:"not null" json:"quantity"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
