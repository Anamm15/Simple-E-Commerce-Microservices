package helper

import (
	"strings"

	"order-service/internal/helper/enum"

	"gorm.io/gorm"
)

func ApplyQueryOptions(
	db *gorm.DB,
	limit *int32,
	offset *int32,
	sort *string,
	filter *enum.OrderStatus,
) *gorm.DB {
	if filter != nil && *filter != "" {
		db = db.Where("status = ?", *filter)
	}

	if sort != nil && *sort != "" {
		order := strings.ToLower(*sort)
		if order == "asc" || order == "desc" {
			db = db.Order("created_at " + order)
		}
	} else {
		db = db.Order("created_at desc")
	}

	if limit != nil && *limit > 0 {
		db = db.Limit(int(*limit))
	}

	if offset != nil && *offset >= 0 {
		db = db.Offset(int(*offset))
	}

	return db
}
