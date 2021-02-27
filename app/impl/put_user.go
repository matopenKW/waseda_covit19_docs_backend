package impl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// PutUserService is put route service
type PutUserService struct{}

// New is put route service new
func (s *PutUserService) New() RequestImpl {
	return &PutUserRequest{}
}

// PutUserRequest is put route request
type PutUserRequest struct {
	UID   string `json:"uid"`
	Email string `json:"email"`
}

// SetRequest is request set receiver
func (r *PutUserRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

// Validate is validate receiver
func (r *PutUserRequest) Validate() error {
	if r.UID == "" {
		return fmt.Errorf("Invalid uid. uid is blank")
	} else if r.Email == "" {
		return fmt.Errorf("Invalid email. email is blank")
	}

	return nil
}

// Execute is api execute receiver
func (r *PutUserRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return putUser(r, ctx)
}

// PutUserResponce is put route responce
type PutUserResponce struct{}

// GetResponce is responce get receiver
func (r *PutUserResponce) GetResponce() {
}

func putUser(req *PutUserRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	err := con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.CreateUser(&model.User{
			ID:    model.UserID(req.UID),
			Email: req.Email,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &PutUserResponce{}, nil
}
