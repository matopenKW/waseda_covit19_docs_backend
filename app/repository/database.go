package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type dbRepository struct {
	db *gorm.DB
}

type dbConnection struct {
	db *gorm.DB
}

type dbTransaction struct {
	db *gorm.DB
}

func NewDbRepository(db *gorm.DB) Repository {
	return &dbRepository{
		db: db,
	}
}

func (r *dbRepository) NewConnection() (Connection, error) {
	return &dbConnection{
		db: r.db,
	}, nil
}

func (c *dbConnection) RunTransaction(f func(Transaction) error) error {
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

func (c *dbConnection) GetPosts() ([]*model.Post, error) {
	var ps []*model.Post
	err := c.db.Find(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *dbConnection) CreatePost(p *model.Post) (*model.Post, error) {
	result := c.db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (c *dbConnection) SavePost(p *model.Post) (*model.Post, error) {
	err := c.db.Save(p).Error

	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *dbConnection) FindRoutesByUserID(UserID string) ([]*model.Route, error) {
	db := c.db.Where("user_id = ?", UserID)

	var ps []*model.Route
	err := db.Find(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *dbConnection) FindActivityProgramsByUserID(UserID string) ([]*model.ActivityProgram, error) {
	db := c.db.Where("user_id = ?", UserID)

	var ps []*model.ActivityProgram
	err := db.Find(&ps).Error
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (c *dbConnection) CreateActivityProgram(p *model.ActivityProgram) (*model.ActivityProgram, error) {
	result := c.db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (c *dbTransaction) SaveRoute(r *model.Route) (*model.Route, error) {
	err := c.db.Save(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}
