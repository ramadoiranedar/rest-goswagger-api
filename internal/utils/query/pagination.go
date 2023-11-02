package utils

import (
	"math"

	"github.com/ramadoiranedar/rest-goswagger-api/gen/models"
	"gorm.io/gorm"
)

func ApplyPagination(db *gorm.DB, currentPage, perPage int64, model interface{}) (*models.Pagination, error) {
	var count int64
	if err := db.Model(model).Count(&count).Error; err != nil {
		return nil, err
	}

	offset := (currentPage - 1) * perPage
	db = db.Offset(int(offset)).Limit(int(perPage))
	pagination := createPagination(int(currentPage), int(perPage), count)

	return pagination, nil
}

func createPagination(currentPage, perPage int, totalData int64) *models.Pagination {
	totalPage := int64(math.Ceil(float64(totalData) / float64(perPage)))

	return &models.Pagination{
		CurrentPage: int64(currentPage),
		PerPage:     int64(perPage),
		TotalData:   totalData,
		TotalPage:   totalPage,
	}
}
