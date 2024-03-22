package migration

import (
	"dimasfadilah/gin-gorm-clean-arch/model"
	"errors"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	db.AutoMigrate(&model.Category{})
}

func SeederData(db *gorm.DB) {
	Seeder(db)
}

func Seeder(db *gorm.DB) error {
	if err := categorySeeder(db); err != nil {
		return err
	}

	return nil
}

func categorySeeder(db *gorm.DB) error {
	categories := []model.Category{
		{
			Name: "Laptop",
		},
		{
			Name: "Smartphone",
		},
		{
			Name: "Tablet",
		},
		{
			Name: "Desktop",
		},
	}

	for _, data := range categories {
		var category model.Category
		err := db.Where("name = ?", data.Name).First(&category).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		isData := db.Find(&category, "nama = ?", data.Name).RowsAffected > 0
		if !isData {
			err = db.Create(&data).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}
