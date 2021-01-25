package impl

import (
	"fmt"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"net/url"
	"strconv"
	"time"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type PutActivityProgramRequest struct {
	UserID           string
	Datetime         *time.Time
	StartTime        *time.Time
	EndTime          *time.Time
	PracticeSection  string
	PracticeContents string
	VenueID          int
	RouteID          int
	ContactPerson1   string
	ContactAbstract1 string
	ContactPerson2   string
	ContactAbstract2 string
}

func (r *PutActivityProgramRequest) SetRequest(form url.Values) {
	r.PracticeSection = form.Get("practice_section")
	r.PracticeContents = form.Get("practice_section")
	r.VenueID, _ = strconv.Atoi(form.Get("venue_id"))
	r.RouteID, _ = strconv.Atoi(form.Get("route_id"))
	r.ContactPerson1 = form.Get("contact_person1")
	r.ContactAbstract1 = form.Get("contact_abstract1")
	r.ContactPerson2 = form.Get("contact_person2")
	r.ContactAbstract2 = form.Get("contact_abstract2")
}

func (r *PutActivityProgramRequest) Validate() error {
	if r == nil {
		return fmt.Errorf("PutActivityProgramRequest is Nil")
	}

	if r.Datetime == nil {
		return fmt.Errorf("Datetime id not nil value")
	}

	return nil
}

func (r *PutActivityProgramRequest) Execute(con repository.Connection) (ResponceImpl, error) {
	return putActivityProgram(r, con)
}

type PutActivityProgramResponce struct {
	Message string
}

func (r *PutActivityProgramResponce) GetResponce() {
}

func putActivityProgram(req *PutActivityProgramRequest, con repository.Connection) (ResponceImpl, error) {
	err := con.RunTransaction(func(tx repository.Transaction) error {
		tx.CreateActivityProgram(&model.ActivityProgram{})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &PutActivityProgramResponce{
		Message: "hello world",
	}, nil
}
