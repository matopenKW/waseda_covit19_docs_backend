package impl

import (
	"net/url"
	"time"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type PutActivityProgramRequest struct {
}

func (r *PutActivityProgramRequest) SetRequest(form url.Values) {

}

func (r *PutActivityProgramRequest) Validate() error {
	return nil
}

func (r *PutActivityProgramRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return putActivityProgram(r, ctx)
}

type PutActivityProgramResponce struct {
	Post *model.ActivityProgram
}

func (r *PutActivityProgramResponce) GetResponce() {
}

func putActivityProgram(req *PutActivityProgramRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	result, err := con.CreateActivityProgram(&model.ActivityProgram{
		ID:               2,
		UserId:           "userid",
		Datetime:         "Datetime",
		StartTime:        "StartTime",
		EndTime:          "EndTime",
		PracticeSection:  "PracticeSection",
		PracticeContents: "PracticeContents",
		VenueID:          2,
		Outbound:         "Outbound",
		ReturnTrip:       "ReturnTrip",
		ContactPerson1:   "ContactPerson1",
		ContactAbstract1: "ContactAbstract1",
		ContactPerson2:   "ContactPerson2",
		ContactAbstract2: "ContactAbstract2",
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &PutActivityProgramResponce{
		Post: result,
	}, nil
}
