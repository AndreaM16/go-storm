package service

import (
	"github.com/andream16/go-storm/model/request"
	"database/sql"
	manufacturerService "github.com/andream16/go-storm/api/feature/manufacturer/service"
	itemService "github.com/andream16/go-storm/api/feature/item/service"
	categoryService "github.com/andream16/go-storm/api/feature/category/service"
	reviewService "github.com/andream16/go-storm/api/feature/review/service"
)

func AddAmazonEntry(amazonEntry request.Amazon, db *sql.DB) error {
	var manufacturer = amazonEntry.Manufacturer
	var item = amazonEntry.Item
	var category = amazonEntry.Category
	var review = amazonEntry.Review
	if len(manufacturer.Name) > 0 {
		manufacturerService.AddManufacturer(manufacturer, db)
	}
	if len(item.Item) > 0 {
		itemService.PatchManufacturer(item, db)
	}
	if len(category.Item) > 0 {
		categoryService.AddCategoriesByItem(category, db)
	}
	if len(review.Item) > 0 {
		reviewService.AddReview(review, db)
	}
	return nil
}