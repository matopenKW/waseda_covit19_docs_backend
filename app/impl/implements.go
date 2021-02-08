package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type Context struct {
	userID     string
	connection repository.Connection
}

func NewContext(userID string, con repository.Connection) *Context {
	return &Context{userID, con}
}

func (c *Context) GetUserID() string {
	return c.userID
}

func (c *Context) GetConnection() repository.Connection {
	return c.connection
}

type RequestImpl interface {
	SetRequest(*gin.Context)
	Validate() error
	Execute(*Context) (ResponceImpl, error)
}

type ResponceImpl interface {
	GetResponce()
}
