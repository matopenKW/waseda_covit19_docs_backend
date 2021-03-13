package impl

import (
	"context"
	"encoding/json"
	"fmt"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// CreateUserService is create user service
type CreateUserService struct{}

// New is create user service new
func (s *CreateUserService) New() RequestImpl {
	return &CreateUserRequest{}
}

// CreateUserRequest is create user request
type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SetRequest is request set receiver
func (r *CreateUserRequest) SetRequest(ctx *gin.Context) {
	buf, _ := ctx.GetRawData()
	var m map[string]string
	_ = json.Unmarshal(buf, &m)

	r.Email = m["email"]
	r.Password = m["password"]
}

// Validate is validate receiver
func (r *CreateUserRequest) Validate() error {
	if r.Email == "" {
		return fmt.Errorf("Invalid email. email is blank")
	} else if r.Password == "" {
		return fmt.Errorf("Invalid password. password is blank")
	}

	return nil
}

// Execute is api execute receiver
func (r *CreateUserRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return createUser(r, ctx)
}

// CreateUserResponce is create user responce
type CreateUserResponce struct{}

// GetResponce is responce get receiver
func (r *CreateUserResponce) GetResponce() {
}

func createUser(req *CreateUserRequest, ctx *Context) (ResponceImpl, error) {
	fb, err := repository.OpenAuthJSON()
	if err != nil {
		return nil, err
	}

	utc := (&auth.UserToCreate{}).Email(req.Email).EmailVerified(true).Password(req.Password)
	ur, err := fb.CreateUser(context.Background(), utc)
	if err != nil {
		return nil, err
	}

	con := ctx.GetConnection()
	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.CreateUser(&model.User{
			ID:    model.UserID(ur.UID),
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

	return &CreateUserResponce{}, nil
}
