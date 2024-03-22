package repository

import (
	"context"
	"dimasfadilah/gin-gorm-clean-arch/model"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

type CategoryRepository interface {
	// transaction related
	BeginTx(ctx context.Context) (*gorm.DB, error)
	CommitTx(ctx context.Context, tx *gorm.DB) error
	RollbackTx(ctx context.Context, tx *gorm.DB) error

	// functional related
	CreateCategory(ctx context.Context, tx *gorm.DB, category model.Category) (model.Category, error)
	FindOneByID(ctx context.Context, tx *gorm.DB, id uint64) (model.Category, error)
	FindAll(ctx context.Context, tx *gorm.DB) ([]model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}

func (c *CategoryRepositoryImpl) BeginTx(ctx context.Context) (*gorm.DB, error) {
	tx := c.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return tx, nil
}

func (c *CategoryRepositoryImpl) CommitTx(ctx context.Context, tx *gorm.DB) error {
	err := tx.WithContext(ctx).Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepositoryImpl) RollbackTx(ctx context.Context, tx *gorm.DB) error {
	err := tx.WithContext(ctx).Rollback().Error
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryRepositoryImpl) CreateCategory(ctx context.Context, tx *gorm.DB, category model.Category) (model.Category, error) {
	var err error
	if tx == nil {
		tx = c.db.WithContext(ctx).Debug().Create(&category)
		err = tx.Error
	} else {
		err = tx.WithContext(ctx).Debug().Create(&category).Error
	}

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (c *CategoryRepositoryImpl) FindOneByID(ctx context.Context, tx *gorm.DB, id uint64) (model.Category, error) {
	var category model.Category
	var err error
	if tx == nil {
		err = c.db.WithContext(ctx).Where("id = ?", id).First(&category).Error
	} else {
		err = tx.WithContext(ctx).Where("id = ?", id).First(&category).Error
	}

	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) ([]model.Category, error) {
	var categories []model.Category
	var err error
	if tx == nil {
		err = c.db.WithContext(ctx).Find(&categories).Error
	} else {
		err = tx.WithContext(ctx).Find(&categories).Error
	}

	if err != nil {
		return []model.Category{}, err
	}

	return categories, nil
}
