package repository

import "github.com/matopenKW/waseda_covit19_docs_backend/app/model"

type Repository interface {
	NewConnection() (Connection, error)
}

type Connection interface {
	RunTransaction(f func(Transaction) error) error

	FindActivityProgram(userID string, seqNo model.ActivityProgramSeqNo) (*model.ActivityProgram, error)
	FindActivityProgramMaxSeqNo(string) (model.ActivityProgramSeqNo, error)
	ListActivityPrograms(string) ([]*model.ActivityProgram, error)
	FindRoute(model.RouteID) (*model.Route, error)
	FindMaxRouteID() (model.RouteID, error)
	FindRoutesByUserID(string) ([]*model.Route, error)
	FindActivityProgramsByUserID(string) ([]*model.ActivityProgram, error)
}

type Transaction interface {
	CreateActivityProgram(*model.ActivityProgram) (*model.ActivityProgram, error)
	SaveRoute(*model.Route) (*model.Route, error)
	UpdateRoute(*model.Route) (*model.Route, error)
	CreateRoute(*model.Route) (*model.Route, error)
	DeleteRoute(model.RouteID) error
}
