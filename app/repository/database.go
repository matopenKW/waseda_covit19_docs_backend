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

// NewDbRepository is db repository creater
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

func (c *dbConnection) FindRoute(id model.RouteID) (*model.Route, error) {
	r := &model.Route{
		ID: id,
	}
	err := c.db.Find(r).Error
	if err != nil {
		return nil, err
	}
	return r, nil
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

func (t *dbTransaction) UpdateRoute(r *model.Route) (*model.Route, error) {
	err := t.db.Update(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (t *dbTransaction) CreateRoute(r *model.Route) (*model.Route, error) {
	err := t.db.Create(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (t *dbTransaction) DeleteRoute(id model.RouteID) error {
	err := t.db.Delete(&model.Route{
		ID: id,
	}).Error
	if err != nil {
		return err
	}

	return nil
}
