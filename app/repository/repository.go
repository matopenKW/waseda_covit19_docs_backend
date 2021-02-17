package repository

import "github.com/matopenKW/waseda_covit19_docs_backend/app/model"

type Repository interface {
	NewConnection() (Connection, error)
}

type Connection interface {
	RunTransaction(f func(Transaction) error) error

	FindRoute(model.RouteID) (*model.Route, error)
	FindRoutesByUserID(string) ([]*model.Route, error)
	FindActivityProgramsByUserID(string) ([]*model.ActivityProgram, error)
}

type Transaction interface {
	CreateActivityProgram(*model.ActivityProgram) (*model.ActivityProgram, error)
	UpdateRoute(*model.Route) (*model.Route, error)
	CreateRoute(*model.Route) (*model.Route, error)
	DeleteRoute(model.RouteID) error
}
