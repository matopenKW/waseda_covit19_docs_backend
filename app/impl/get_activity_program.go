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

type ResponceActivityProgram struct {
	SeqNo              model.ActivityProgramSeqNo
	Datetime           string
	StartTime          string
	EndTime            string
	PracticeSectionID  uint
	PracticeContentsID uint
	OutwardTrip        string
	ReturnTrip         string
	ContactPerson1     int
	ContactAbstract1   string
	ContactPerson2     int
	ContactAbstract2   string
	PlaceID            int
}

// GetActivityProgramResponce is responce
type GetActivityProgramResponce struct {
	ActivityProgram *ResponceActivityProgram
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
		ActivityProgram: PresenterActivityProgram(ap),
		PracticeName:    practiceName,
		ActivityName:    activityName,
	}, nil
}

func PresenterActivityProgram(ap *model.ActivityProgram) *ResponceActivityProgram {
	return &ResponceActivityProgram{
		SeqNo:              ap.SeqNo,
		Datetime:           ap.Datetime,
		StartTime:          ap.StartTime,
		EndTime:            ap.EndTime,
		PracticeSectionID:  ap.PracticeSectionID,
		PracticeContentsID: ap.PracticeContentsID,
		OutwardTrip:        ap.OutwardTrip,
		ReturnTrip:         ap.ReturnTrip,
		ContactPerson1:     ap.ContactPerson1,
		ContactAbstract1:   ap.ContactAbstract1,
		ContactPerson2:     ap.ContactPerson2,
		ContactAbstract2:   ap.ContactAbstract2,
		PlaceID:            ap.PlaceID,
	}
}
