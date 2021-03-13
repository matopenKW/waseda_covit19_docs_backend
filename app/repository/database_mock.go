package repository

import (
	"fmt"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

var mock dbMock

type dbMock struct {
	activityPrograms []*model.ActivityProgram
	routes           []*model.Route
	lastUploads      []*model.LastUpload
	users            []*model.User
}

func NewDBMock() *dbMock {
	return &dbMock{}
}

func (m *dbMock) SetActivityPrograms(aps []*model.ActivityProgram) {
	m.activityPrograms = aps
}

func (m *dbMock) SetRoutes(rs []*model.Route) {
	m.routes = rs
}

func (m *dbMock) SetUsers(us []*model.User) {
	m.users = us
}

type dbMockRepository struct {
}

type dbMockConnection struct {
}

type dbMockTransaction struct {
}

// NewMockDbRepository is mock db repository creater
func NewMockDbRepository(m *dbMock) Repository {
	mock = *m
	return &dbMockRepository{}
}

func (r *dbMockRepository) NewConnection() (Connection, error) {
	return &dbMockConnection{}, nil
}

func (c *dbMockConnection) RunTransaction(f func(Transaction) error) error {
	return f(&dbMockTransaction{})
}

func (c *dbMockConnection) FindActivityProgram(userID string, seqNo model.ActivityProgramSeqNo) (*model.ActivityProgram, error) {
	for _, v := range mock.activityPrograms {
		if v.UserID == userID && v.SeqNo == seqNo {
			return v, nil
		}
	}
	return nil, nil
}

func (c *dbMockConnection) FindActivityProgramMaxSeqNo(userID string) (model.ActivityProgramSeqNo, error) {
	seq := model.ActivityProgramSeqNo(0)
	for _, v := range mock.activityPrograms {
		if v.UserID == userID {
			seq = v.SeqNo
		}
	}
	return seq, nil
}

func (c *dbMockConnection) ListActivityPrograms(userID string) ([]*model.ActivityProgram, error) {
	aps := []*model.ActivityProgram{}
	for _, v := range mock.activityPrograms {
		if v.UserID == userID {
			aps = append(aps, v)
		}
	}
	return aps, nil
}

func (c *dbMockConnection) FindRoute(userID string, seqNo model.RouteSeqNo) (*model.Route, error) {
	for _, v := range mock.routes {
		if v.UserID == userID && v.SeqNo == seqNo {
			return v, nil
		}
	}
	return nil, nil
}

func (c *dbMockConnection) FindRouteMaxSeqNo(userID string) (model.RouteSeqNo, error) {
	seq := model.RouteSeqNo(0)
	for _, v := range mock.routes {
		if v.UserID == userID {
			seq = v.SeqNo
		}
	}
	return seq, nil
}

func (c *dbMockConnection) FindRoutesByUserID(UserID string) ([]*model.Route, error) {
	return nil, nil
}

func (c *dbMockConnection) FindActivityProgramsByUserID(UserID string) ([]*model.ActivityProgram, error) {
	return nil, nil
}

func (c *dbMockConnection) LatestLastUpload() (*model.LastUpload, error) {
	if len(mock.lastUploads) != 0 {
		return mock.lastUploads[len(mock.lastUploads)-1], nil
	}
	return nil, nil
}

func (c *dbMockConnection) ListUser() ([]*model.User, error) {
	return mock.users, nil
}

func (t *dbMockTransaction) CreateActivityProgram(ap *model.ActivityProgram) (*model.ActivityProgram, error) {
	for _, v := range mock.activityPrograms {
		if v.UserID == ap.UserID && v.SeqNo == ap.SeqNo {
			return nil, fmt.Errorf("conflict")
		}
	}
	mock.activityPrograms = append(mock.activityPrograms, ap)
	return ap, nil
}

func (t *dbMockTransaction) SaveRoute(r *model.Route) (*model.Route, error) {
	exists := false
	for _, v := range mock.routes {
		if v.UserID == r.UserID && v.SeqNo == r.SeqNo {
			exists = true
			break
		}
	}
	if exists {
		return t.UpdateRoute(r)
	}
	return t.CreateRoute(r)
}

func (t *dbMockTransaction) UpdateRoute(r *model.Route) (*model.Route, error) {
	for _, v := range mock.routes {
		if v.UserID == r.UserID && v.SeqNo == r.SeqNo {
			v.UserID = r.UserID
			v.SeqNo = r.SeqNo
			v.Name = r.Name
			v.OutwardTrip = r.OutwardTrip
			v.ReturnTrip = r.ReturnTrip
		}
		return r, nil
	}
	return r, nil
}

func (t *dbMockTransaction) CreateRoute(r *model.Route) (*model.Route, error) {
	seq := model.RouteSeqNo(0)
	for _, v := range mock.routes {
		if v.UserID == r.UserID && seq <= v.SeqNo {
			seq = v.SeqNo
		}
	}
	r.SeqNo = seq + 1
	mock.routes = append(mock.routes, r)
	return r, nil
}

func (t *dbMockTransaction) DeleteRoute(userID string, seqNo model.RouteSeqNo) error {
	rs := make([]*model.Route, 0, 0)
	for _, v := range mock.routes {
		if v.UserID == userID && v.SeqNo == seqNo {
			continue
		}
		rs = append(rs, v)
	}

	mock.routes = rs
	return nil
}

func (t *dbMockTransaction) UpdateLastUpload(m *model.LastUpload) error {
	for _, v := range mock.lastUploads {
		v.DriveID = m.DriveID
	}
	return nil
}

func (t *dbMockTransaction) CreateUser(u *model.User) error {
	for _, m := range mock.users {
		if m.ID == u.ID {
			return fmt.Errorf("conflict")
		}
	}

	mock.users = append(mock.users, u)
	return nil
}
