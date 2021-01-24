package impl

import "github.com/matopenKW/waseda_covit19_docs_backend/app/model"

type ResponceImpl interface {
	GetResponce()
}

type HelloWorldResponce struct {
	Message string
}

func (r *HelloWorldResponce) GetResponce() {
}

type GetPostsResponce struct {
	Posts []*model.Post
}

func (r *GetPostsResponce) GetResponce() {
}

type PutPostResponce struct {
	Post *model.Post
}

func (r *PutPostResponce) GetResponce() {
}

type UpdatePostResponce struct {
	Post *model.Post
}

func (r *UpdatePostResponce) GetResponce() {
}
