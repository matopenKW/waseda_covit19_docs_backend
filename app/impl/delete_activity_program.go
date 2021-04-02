package impl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// DeleteActivityProgramService is delete activity program service
type DeleteActivityProgramService struct{}

// New is delete activity program service new
func (s *DeleteActivityProgramService) New() RequestImpl {
	return &DeleteActivityProgramRequest{}
}

// DeleteActivityProgramRequest is delete activity program request
type DeleteActivityProgramRequest struct {
	SeqNo *int `json:"seq_no, omitempty"`
}

// SetRequest is request set receiver
func (r *DeleteActivityProgramRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

// Validate is validate receiver
func (r *DeleteActivityProgramRequest) Validate() error {
	if r.SeqNo == nil {
		return fmt.Errorf("invalid seq no")
	}
	return nil
}

// Execute is api execute receiver
func (r *DeleteActivityProgramRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return DeleteActivityProgram(r, ctx)
}

// DeleteActivityProgramResponce is put route responce
type DeleteActivityProgramResponce struct {
}

// GetResponce is responce get receiver
func (r *DeleteActivityProgramResponce) GetResponce() {
}

// DeleteActivityProgram is api execute function of private
func DeleteActivityProgram(req *DeleteActivityProgramRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	err := con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.DeleteActivityProgram(ctx.userID, model.ActivityProgramSeqNo(*req.SeqNo))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &DeleteActivityProgramResponce{}, nil
}
