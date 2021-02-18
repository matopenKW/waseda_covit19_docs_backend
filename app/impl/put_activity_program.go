package impl

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// ActivityProgramService is activity program service
type ActivityProgramService struct{}

// New is activity program service new
func (s *ActivityProgramService) New() RequestImpl {
	return &ActivityProgramRequest{}
}

// ActivityProgramRequest is put activity program request
type ActivityProgramRequest struct {
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
func (r *ActivityProgramRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

// Validate is validate receiver
func (r *ActivityProgramRequest) Validate() error {
	return nil
}

// Execute is api execute receiver
func (r *ActivityProgramRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return activityProgram(r, ctx)
}

// ActivityProgramResponce is put activity program responce
type ActivityProgramResponce struct {
	Post *model.ActivityProgram
}

// GetResponce is responce get receiver
func (r *ActivityProgramResponce) GetResponce() {
}

func activityProgram(req *ActivityProgramRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	log.Println(req)

	maxID, err := con.FindMaxActivityProgramID()
	if err != nil {
		return nil, err
	}

	var result *model.ActivityProgram
	err = con.RunTransaction(func(tx repository.Transaction) error {
		result, err = tx.CreateActivityProgram(&model.ActivityProgram{
			ID:               maxID + 1,
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

	return &ActivityProgramResponce{
		Post: result,
	}, nil
}

// FIXME test data
// `
// {
// 	"datetime": null,
// 	"start_time": "0900",
// 	"end_time": "1800",
// 	"practice_section": "aaa",
// 	"practice_contents": "ssss",
// 	"venue_id": 1,
// 	"route_id": 3,
// 	"outward_trip": "",
// 	"return_trip": "",
// 	"contact_person1": false,
// 	"contact_abstract1": "",
// 	"contact_person2": false,
// 	"contact_abstract2": ""
// }
// `
