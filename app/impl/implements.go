package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type Master struct {
	practices  []*model.Practice
	activities []*model.Activity
	parts      []*model.Part
}

func NewMaster(practices []*model.Practice, activities []*model.Activity, parts []*model.Part) *Master {
	return &Master{
		practices:  practices,
		activities: activities,
		parts:      parts,
	}
}

type Context struct {
	userID     model.UserID
	connection repository.Connection
	master     *Master
}

func NewContext(userID model.UserID, con repository.Connection, master *Master) *Context {
	return &Context{userID, con, master}
}

func (c *Context) GetUserID() model.UserID {
	return c.userID
}

func (c *Context) GetConnection() repository.Connection {
	return c.connection
}

type ServiceImpl interface {
	New() RequestImpl
}

type RequestImpl interface {
	SetRequest(*gin.Context)
	Validate() error
	Execute(*Context) (ResponceImpl, error)
}

type ResponceImpl interface {
	GetResponce()
}
