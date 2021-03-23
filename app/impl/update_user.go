package impl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// UpdateUserService is put route service
type UpdateUserService struct{}

// New is put route service new
func (s *UpdateUserService) New() RequestImpl {
	return &UpdateUserRequest{}
}

// UpdateUserRequest is put route request
type UpdateUserRequest struct {
	Name           string `json:"name"`
	UniversityName string `json:"university_name"`
	FacultyName    string `json:"faculty_name"`
	StudentNo      string `json:"student_no"`
	CellPhonNo     string `json:"cell_phon_no"`
	Ki             int    `json:"ki"`
	PartID         int    `json:"part_id"`
}

// SetRequest is request set receiver
func (r *UpdateUserRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

// Validate is validate receiver
func (r *UpdateUserRequest) Validate() error {
	return nil
}

// Execute is api execute receiver
func (r *UpdateUserRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return UpdateUser(r, ctx)
}

// UpdateUserResponce is put route responce
type UpdateUserResponce struct{}

// GetResponce is responce get receiver
func (r *UpdateUserResponce) GetResponce() {
}

func UpdateUser(req *UpdateUserRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	u, err := con.FindUser(ctx.userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, fmt.Errorf("user is empty; user_id=%s", ctx.userID)
	}

	err = con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.UpdateUser(&model.User{
			ID:             ctx.userID,
			Email:          u.Email,
			Name:           req.Name,
			UniversityName: req.UniversityName,
			FacultyName:    req.FacultyName,
			StudentNo:      req.StudentNo,
			CellPhonNo:     req.CellPhonNo,
			Ki:             req.Ki,
			PartID:         req.PartID,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &UpdateUserResponce{}, nil
}
