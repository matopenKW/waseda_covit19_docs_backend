package impl

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

func TestActivityProgramService_SetRequest(t *testing.T) {
	want := &PutActivityProgramRequest{
		Datetime:         "20200101",
		StartTime:        "0900",
		EndTime:          "1800",
		PracticeSection:  "test_practice_section",
		PracticeContents: "test_practice_contents",
		OutwardTrip:      "test_outward_trip",
		ReturnTrip:       "test_return_trip",
		ContactPerson1:   1,
		ContactAbstract1: "test_contact_abstract1",
		ContactPerson2:   1,
		ContactAbstract2: "test_contact_abstract2",
	}

	bs := []byte(`{
			"datetime": "20200101",
			"start_time": "0900",
			"end_time": "1800",
			"practice_section": "test_practice_section",
			"practice_contents": "test_practice_contents",
			"outward_trip": "test_outward_trip",
			"return_trip": "test_return_trip",
			"contact_person1": 1,
			"contact_abstract1": "test_contact_abstract1",
			"contact_person2": 1,
			"contact_abstract2": "test_contact_abstract2"
		}`)
	req, _ := http.NewRequest("GET", "/api/v1/put_activity_program", bytes.NewBuffer(bs))

	var ctx *gin.Context
	ctx = &gin.Context{
		Request: req,
	}

	impl := &PutActivityProgramRequest{}
	impl.SetRequest(ctx)

	if !reflect.DeepEqual(want, impl) {
		t.Errorf("SetRequest is fatal. impl=%#v, want=%#v; diff=%s", impl, want, cmp.Diff(impl, want))
	}
}
