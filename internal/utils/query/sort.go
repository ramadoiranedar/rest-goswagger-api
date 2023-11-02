package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func ApplySorting(db *gorm.DB, sort, sortBy *string) {
	if sort != nil && sortBy != nil {
		orderBy := fmt.Sprintf("%s %s", *sortBy, *sort)
		db = db.Order(orderBy)
	}
}
