package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// PutActivityProgramService is put activity program service
type PutActivityProgramService struct{}

// New is activity program service new
func (s *PutActivityProgramService) New() RequestImpl {
	return &PutActivityProgramRequest{}
}

// PutActivityProgramRequest is put activity program request
type PutActivityProgramRequest struct {
	Datetime           string `json:"datetime"`
	StartTime          string `json:"start_time"`
	EndTime            string `json:"end_time"`
	PracticeSectionID  uint   `json:"practice_section_id"`
	PracticeContentsID uint   `json:"practice_contents_id"`
	OutwardTrip        string `json:"outward_trip"`
	ReturnTrip         string `json:"return_trip"`
	ContactPerson1     int    `json:"contact_person1"`
	ContactAbstract1   string `json:"contact_abstract1"`
	ContactPerson2     int    `json:"contact_person2"`
	ContactAbstract2   string `json:"contact_abstract2"`
}

// SetRequest is request set receiver
func (r *PutActivityProgramRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

// Validate is validate receiver
func (r *PutActivityProgramRequest) Validate() error {
	return nil
}

// Execute is api execute receiver
func (r *PutActivityProgramRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return activityProgram(r, ctx)
}

// PutActivityProgramResponce is put activity program responce
type PutActivityProgramResponce struct {
	ActivityProgram *model.ActivityProgram
}

// GetResponce is responce get receiver
func (r *PutActivityProgramResponce) GetResponce() {
}

func activityProgram(req *PutActivityProgramRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	maxSeq, err := con.FindActivityProgramMaxSeqNo(ctx.userID)
	if err != nil {
		return nil, err
	}

	var result *model.ActivityProgram
	err = con.RunTransaction(func(tx repository.Transaction) error {
		result, err = tx.CreateActivityProgram(&model.ActivityProgram{
			UserID:             ctx.userID,
			SeqNo:              maxSeq + 1,
			Datetime:           req.Datetime,
			StartTime:          req.StartTime,
			EndTime:            req.EndTime,
			PracticeSectionID:  req.PracticeSectionID,
			PracticeContentsID: req.PracticeContentsID,
			OutwardTrip:        req.OutwardTrip,
			ReturnTrip:         req.ReturnTrip,
			ContactPerson1:     req.ContactPerson1,
			ContactAbstract1:   req.ContactAbstract1,
			ContactPerson2:     req.ContactPerson2,
			ContactAbstract2:   req.ContactAbstract2,
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &PutActivityProgramResponce{
		ActivityProgram: result,
	}, nil
}
