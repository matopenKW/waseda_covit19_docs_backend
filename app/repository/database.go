package repository

import (
	"errors"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

// NewDbConnection is db connection
func NewDbConnection() (*gorm.DB, error) {
	return gorm.Open("postgres", os.Getenv("DATABASE_URL"))
}

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

func (c *dbConnection) FindActivityProgram(userID string, seqNo model.ActivityProgramSeqNo) (*model.ActivityProgram, error) {
	ap := &model.ActivityProgram{}
	err := c.db.Where("user_id = ? and seq_no = ?", userID, seqNo).Find(&ap).Error
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

func (c *dbConnection) FindRoute(userID string, seqNo model.RouteSeqNo) (*model.Route, error) {
	r := &model.Route{}
	err := c.db.Where("user_id = ? and seq_no = ?", userID, seqNo).Find(&r).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return r, nil
}

func (c *dbConnection) FindRouteMaxSeqNo(userID string) (model.RouteSeqNo, error) {
	r := &model.Route{}
	err := c.db.Order("seq_no DESC").Where("user_id = ?", userID).Last(r).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return 0, err
	}
	return r.SeqNo, nil
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

func (c *dbConnection) LatestLastUpload() (*model.LastUpload, error) {
	result := &model.LastUpload{}
	err := c.db.Last(&result).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return result, nil
}

func (c *dbConnection) ListUser() ([]*model.User, error) {
	us := []*model.User{}
	err := c.db.Find(us).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return us, nil
}

func (t *dbTransaction) CreateActivityProgram(p *model.ActivityProgram) (*model.ActivityProgram, error) {
	result := t.db.Create(p)
	if result.Error != nil {
		return nil, result.Error
	}

	return p, nil
}

func (t *dbTransaction) SaveRoute(r *model.Route) (*model.Route, error) {
	err := t.db.Save(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (t *dbTransaction) UpdateRoute(r *model.Route) (*model.Route, error) {
	err := t.db.Model(&model.Route{}).Where("user_id = ? and seq_no = ?", r.UserID, r.SeqNo).Update(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (t *dbTransaction) CreateRoute(r *model.Route) (*model.Route, error) {
	err := t.db.Create(&r).Error
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (t *dbTransaction) DeleteRoute(userID string, seqNo model.RouteSeqNo) error {
	err := t.db.Where("user_id = ? and seq_no = ?", userID, seqNo).Delete(&model.Route{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *dbTransaction) UpdateLastUpload(m *model.LastUpload) error {
	err := t.db.Model(&model.LastUpload{}).Update("drive_id", m.DriveID).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *dbTransaction) CreateUser(u *model.User) error {
	err := t.db.Create(u).Error
	if err != nil {
		return err
	}

	return nil
}
