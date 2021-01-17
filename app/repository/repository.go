package repository

import "github.com/matopenKW/waseda_covit19_docs_backend/app/model"

type Repository interface {
	NewConnection() (Connection, error)
}

type Connection interface {
	RunTransaction(f func(Transaction) error) error

	GetPosts() ([]*model.Post, error)
	CreatePost(*model.Post) (*model.Post, error)
	SavePost(*model.Post) (*model.Post, error)
}

type Transaction interface {
}
