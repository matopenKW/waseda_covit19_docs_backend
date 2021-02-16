package repository

import (
	"github.com/jinzhu/gorm"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type dbMockRepository struct {
}

type dbMockConnection struct {
	db *gorm.DB
}

type dbMockTransaction struct {
	db *gorm.DB
}

// NewMockDbRepository is mock db repository creater
func NewMockDbRepository() Repository {
	return &dbMockRepository{}
}

func (r *dbMockRepository) NewConnection() (Connection, error) {
	return &dbMockConnection{}, nil
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
	return nil, nil
}

func (c *dbMockConnection) SavePost(p *model.Post) (*model.Post, error) {
	return nil, nil
}

func (c *dbMockConnection) FindRoutesByUserID(UserID string) ([]*model.Route, error) {
	return nil, nil
}

func (c *dbMockConnection) FindActivityProgramsByUserID(UserID string) ([]*model.ActivityProgram, error) {
	return nil, nil
}

func (c *dbMockConnection) CreateActivityProgram(p *model.ActivityProgram) (*model.ActivityProgram, error) {
	return nil, nil
}

func (c *dbMockTransaction) SaveRoute(r *model.Route) (*model.Route, error) {
	return nil, nil
}

func (c *dbMockTransaction) DeleteRoute(id model.RouteID) error {
	return nil
}
