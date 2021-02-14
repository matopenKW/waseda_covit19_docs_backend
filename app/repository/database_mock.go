package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type dbMockConnection struct {
	db *gorm.DB
}

type dbMockTransaction struct {
	db *gorm.DB
}

func (r *dbRepository) NewMockConnection() (Connection, error) {
	return &dbConnection{
		db: r.db,
	}, nil
}

func (c *dbMockConnection) RunTransaction(f func(Transaction) error) error {
	tx := c.db.Begin()

	err := f(&dbTransaction{
		db: tx,
	})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (c *dbMockConnection) GetPosts() ([]*model.Post, error) {
	return []*model.Post{
		{
			ID:      1,
			Content: "test_content",
			Author:  "author",
		},
		{},
	}, nil
}

func (c *dbMockConnection) CreatePost(p *model.Post) (*model.Post, error) {
	result := c.db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (c *dbMockConnection) SavePost(p *model.Post) (*model.Post, error) {
	err := c.db.Save(p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *dbMockConnection) FindRoutesByUserID(UserID string) ([]*model.Route, error) {
	db := c.db.Where("user_id = ?", UserID)

	var ps []*model.Route
	err := db.Find(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *dbMockConnection) FindActivityProgramsByUserID(UserID string) ([]*model.ActivityProgram, error) {
	db := c.db.Where("user_id = ?", UserID)

	var ps []*model.ActivityProgram
	err := db.Find(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *dbMockConnection) CreateActivityProgram(p *model.ActivityProgram) (*model.ActivityProgram, error) {
	result := c.db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (c *dbMockConnection) SaveRoute(r *model.Route) (*model.Route, error) {
	err := c.db.Save(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}
