package helper

import (
	"strings"

	"gorm.io/gorm"
)

func ApplyQueryOptions(
	db *gorm.DB,
	limit *int32,
	offset *int32,
	sort *string,
	filter *string,
) *gorm.DB {
	if filter != nil && *filter != "" {
		// contoh filter: status
		db = db.Where("status = ?", *filter)
	}

	if sort != nil && *sort != "" {
		// contoh: created_at desc | created_at asc
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
