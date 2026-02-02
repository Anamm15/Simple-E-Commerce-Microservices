package helper

import (
	"strings"

	"order-service/internal/helper/enum"

	"gorm.io/gorm"
)

func QueryValidation(page *int32, limit *int32, sort *string, filter *enum.OrderStatus) (int32, int32, string, enum.OrderStatus) {
	var finalPage, finalLimit int32
	finalSort := "desc"
	var finalFilter enum.OrderStatus = ""

	finalPage = 1
	finalLimit = 10
	if page != nil && *page > 0 {
		finalPage = *page
	}

	if limit != nil && *limit > 0 {
		finalLimit = *limit
	}

	if sort != nil {
		s := *sort
		if s == "asc" || s == "desc" {
			finalSort = s
		}
	}
	if filter != nil {
		f := *filter
		switch f {
		case enum.OrderStatusPending, enum.OrderStatusPaid, enum.OrderStatusFailed:
			finalFilter = f
		default:
			finalFilter = ""
		}
	}

	return finalPage, finalLimit, finalSort, finalFilter
}

func ApplyQueryOptions(
	db *gorm.DB,
	limit int32,
	offset int32,
	sort string,
	filter enum.OrderStatus,
) *gorm.DB {
	if db == nil {
		return nil
	}

	// Filter
	if filter != "" {
		db = db.Where("status = ?", filter)
	}

	// Sorting - Tambahkan validasi agar tidak terjadi SQL Injection
	order := strings.ToLower(sort)
	if order == "asc" || order == "desc" {
		db = db.Order("created_at " + order)
	} else {
		// Default sorting
		db = db.Order("created_at DESC")
	}

	// Limit & Offset
	if limit > 0 {
		db = db.Limit(int(limit))
	}

	if offset > 0 {
		db = db.Offset(int(offset))
	}

	return db
}
