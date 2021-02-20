package repository

import (
	"errors"

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

func (c *dbConnection) FindActivityProgram(ap *model.ActivityProgram) (*model.ActivityProgram, error) {
	err := c.db.Find(ap).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return ap, nil
}

func (c *dbConnection) FindActivityProgramMaxSeqNo(userID string) (model.ActivityProgramSeqNo, error) {
	ap := &model.ActivityProgram{}
	err := c.db.Limit(1).Order("seq_no DESC").Where("user_id = ?", userID).Find(ap).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return ap.SeqNo, nil
}

func (c *dbConnection) ListActivityPrograms(userID string) ([]*model.ActivityProgram, error) {
	aps := []*model.ActivityProgram{}
	err := c.db.Find(&aps).Where("user_id = ?", userID).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return aps, nil
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

func (c *dbConnection) FindMaxRouteID() (model.RouteID, error) {
	r := &model.Route{}
	err := c.db.Last(r).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return r.ID, nil
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

func (t *dbTransaction) CreateActivityProgram(p *model.ActivityProgram) (*model.ActivityProgram, error) {
	result := t.db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (t *dbTransaction) SaveRoute(r *model.Route) (*model.Route, error) {
	err := t.db.Save(r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
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
