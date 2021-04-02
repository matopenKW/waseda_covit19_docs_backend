package impl

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

// GetActivityProgramService is get ActivityProgram service
type GetActivityProgramService struct {
}

// New is get ActivityProgram service new
func (s *GetActivityProgramService) New() RequestImpl {
	return &GetActivityProgramRequest{}
}

// GetActivityProgramRequest is get activity program request
type GetActivityProgramRequest struct {
	SeqNo int `json:"seq_no"`
}

// SetRequest is set request
func (r *GetActivityProgramRequest) SetRequest(ctx *gin.Context) {
	r.SeqNo, _ = strconv.Atoi(ctx.Param("seq_no"))
}

// Validate is validate
func (r *GetActivityProgramRequest) Validate() error {
	if r.SeqNo == 0 {
		return fmt.Errorf("invalid seq no")
	}
	return nil
}

// Execute is execute
func (r *GetActivityProgramRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return getActivityProgram(r, ctx)
}

// GetActivityProgramResponce is responce
type GetActivityProgramResponce struct {
	ActivityProgram *model.ActivityProgram
	PracticeName    string
	ActivityName    string
}

// GetResponce is get responce
func (r *GetActivityProgramResponce) GetResponce() {
}

func getActivityProgram(req *GetActivityProgramRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	ap, err := con.FindActivityProgram(ctx.userID, model.ActivityProgramSeqNo(req.SeqNo))
	if err != nil {
		return nil, err
	}

	if ap == nil {
		return &GetActivityProgramResponce{
			ActivityProgram: nil,
		}, nil
	}

	var practiceName string
	for _, p := range ctx.master.practices {
		if p.ID == ap.PracticeSectionID {
			practiceName = p.Name
		}
	}

	var activityName string
	for _, a := range ctx.master.activities {
		if a.ID == ap.PracticeContentsID {
			activityName = a.Name
		}
	}
	return &GetActivityProgramResponce{
		ActivityProgram: ap,
		PracticeName:    practiceName,
		ActivityName:    activityName,
	}, nil
}
