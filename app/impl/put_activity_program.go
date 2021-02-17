package impl

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// PutActivityProgramRequest is put activity program request
type PutActivityProgramRequest struct {
	Datetime         *time.Time `json:"datetime"`
	StartTime        string     `json:"start_time"`
	EndTime          string     `json:"end_time"`
	PracticeSection  string     `json:"practice_section"`
	PracticeContents string     `json:"practice_contents"`
	VenueID          int        `json:"venue_id"`
	OutwardTrip      string     `json:"outward_trip"`
	ReturnTrip       string     `json:"return_trip"`
	ContactPerson1   bool       `json:"contact_person_1"`
	ContactAbstract1 string     `json:"contact_abstract_1"`
	ContactPerson2   bool       `json:"contact_person_2"`
	ContactAbstract2 string     `json:"contact_abstract_2"`
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
	return putActivityProgram(r, ctx)
}

// PutActivityProgramResponce is put activity program responce
type PutActivityProgramResponce struct {
	Post *model.ActivityProgram
}

// GetResponce is responce get receiver
func (r *PutActivityProgramResponce) GetResponce() {
}

func putActivityProgram(req *PutActivityProgramRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	var result *model.ActivityProgram
	var err error
	err = con.RunTransaction(func(tx repository.Transaction) error {
		result, err = tx.CreateActivityProgram(&model.ActivityProgram{
			UserID:           ctx.userID,
			Datetime:         req.Datetime,
			StartTime:        req.StartTime,
			EndTime:          req.EndTime,
			PracticeSection:  req.PracticeSection,
			PracticeContents: req.PracticeContents,
			VenueID:          req.VenueID,
			OutwardTrip:      req.OutwardTrip,
			ReturnTrip:       req.ReturnTrip,
			ContactPerson1:   req.ContactPerson1,
			ContactAbstract1: req.ContactAbstract1,
			ContactPerson2:   req.ContactPerson2,
			ContactAbstract2: req.ContactAbstract2,
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
		Post: result,
	}, nil
}
