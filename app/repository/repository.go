package repository

import (
	"io"

	"google.golang.org/api/drive/v3"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type Repository interface {
	NewConnection() (Connection, error)
}

type Connection interface {
	RunTransaction(f func(Transaction) error) error

	FindUser(model.UserID) (*model.User, error)
	FindActivityProgram(model.UserID, model.ActivityProgramSeqNo) (*model.ActivityProgram, error)
	FindActivityProgramMaxSeqNo(model.UserID) (model.ActivityProgramSeqNo, error)
	ListActivityPrograms() ([]*model.ActivityProgram, error)
	FindRoute(model.UserID, model.RouteSeqNo) (*model.Route, error)
	FindRouteMaxSeqNo(model.UserID) (model.RouteSeqNo, error)
	FindRoutesByUserID(model.UserID) ([]*model.Route, error)
	ListActivityProgramsByUserID(model.UserID) ([]*model.ActivityProgram, error)
	LatestLastUpload() (*model.LastUpload, error)
	ListUser() ([]*model.User, error)
}

type Transaction interface {
	CreateActivityProgram(*model.ActivityProgram) (*model.ActivityProgram, error)
	SaveRoute(*model.Route) (*model.Route, error)
	UpdateRoute(*model.Route) (*model.Route, error)
	CreateRoute(*model.Route) (*model.Route, error)
	DeleteRoute(model.UserID, model.RouteSeqNo) error
	UpdateLastUpload(*model.LastUpload) error
	CreateUser(*model.User) error
	UpdateUser(*model.User) error
	UpdateActivityProgram(*model.ActivityProgram) (*model.ActivityProgram, error)
	DeleteActivityProgram(model.UserID, model.ActivityProgramSeqNo) error
}

type GoogleDriveRepository interface {
	GetClient() (GoogleDriveClient, error)
}

type GoogleDriveClient interface {
	GetService() (GoogleDriveService, error)
}

type GoogleDriveService interface {
	Create(io.Reader, *drive.File) (*drive.File, error)
	Delete(string) error
}
