package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	OrderID uuid.UUID `gorm:"type:uuid;index;not null" json:"order_id"`

	Amount        int64   `gorm:"not null" json:"amount"`
	PaymentMethod *string `gorm:"type:varchar(50)" json:"payment_method"`
	Status        string  `gorm:"type:varchar(30);index;not null; default:'pending'" json:"status"`

	MidtransTransactionID *string `gorm:"type:varchar(100)" json:"midtrans_transaction_id,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}
