package repository

import (
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

var mock dbMock

type dbMock struct {
	routes []*model.Route
}

func NewDBMock() *dbMock {
	return &dbMock{}
}

func (m *dbMock) SetRoutes(rs []*model.Route) {
	m.routes = rs
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

func (c *dbMockConnection) FindRoute(id model.RouteID) (*model.Route, error) {
	for _, v := range mock.routes {
		if v.ID == id {
			return v, nil
		}
	}
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

func (t *dbMockTransaction) UpdateRoute(r *model.Route) (*model.Route, error) {
	for _, v := range mock.routes {
		if v.ID == r.ID {
			v.ID = r.ID
			v.UserID = r.UserID
			v.Name = r.Name
			v.OutwardTrip = r.OutwardTrip
			v.ReturnTrip = r.ReturnTrip
		}
		return r, nil
	}
	return r, nil
}

func (t *dbMockTransaction) CreateRoute(r *model.Route) (*model.Route, error) {
	id := model.RouteID(0)
	for _, v := range mock.routes {
		if id <= v.ID {
			id = v.ID
		}
	}
	r.ID = id + 1
	mock.routes = append(mock.routes, r)
	return r, nil
}

func (t *dbMockTransaction) DeleteRoute(id model.RouteID) error {
	rs := make([]*model.Route, 0, 0)
	for _, v := range mock.routes {
		if v.ID == id {
			continue
		}
		rs = append(rs, v)
	}

	mock.routes = rs
	return nil
}
