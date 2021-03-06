package impl

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestGetActivityPrograms_SetRequest(t *testing.T) {
	want := &GetActivityProgramsRequest{}
	req, _ := http.NewRequest("GET", "/api/v1/delete_put", nil)

	var ctx *gin.Context
	ctx = &gin.Context{
		Request: req,
	}

	impl := &GetActivityProgramsRequest{}
	impl.SetRequest(ctx)

	if !reflect.DeepEqual(want, impl) {
		t.Errorf("SetRequest is fatal. impl=%#v, want=%#v", impl, want)
	}
}

func TestGetActivityPrograms_Validate(t *testing.T) {
	impl := &GetActivityProgramsRequest{}
	err := impl.Validate()
	if err != nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestGetActivityPrograms_Execute(t *testing.T) {
	want := &GetActivityProgramsResponce{
		ActivityPrograms: map[string][]*HistoryForEachUser{
			"20200101": {
				{
					ActivityProgram: &ResponceActivityProgram{
						SeqNo:              1,
						Datetime:           "20200101",
						StartTime:          "test_start_time",
						EndTime:            "end_start_time",
						PracticeSectionID:  1,
						PracticeContentsID: 1,
						OutwardTrip:        "outward_trip",
						ReturnTrip:         "return_trip",
						ContactPerson1:     1,
						ContactAbstract1:   "contact_abstract_1",
						ContactPerson2:     1,
						ContactAbstract2:   "contact_abstract_2",
					},
					User: &ResponceUser{Name: "test_user_name_1"},
				},
				{
					ActivityProgram: &ResponceActivityProgram{
						SeqNo:              1,
						Datetime:           "20200101",
						StartTime:          "test_start_time",
						EndTime:            "end_start_time",
						PracticeSectionID:  1,
						PracticeContentsID: 1,
						OutwardTrip:        "outward_trip",
						ReturnTrip:         "return_trip",
						ContactPerson1:     1,
						ContactAbstract1:   "contact_abstract_1",
						ContactPerson2:     1,
						ContactAbstract2:   "contact_abstract_2",
					},
					User: &ResponceUser{Name: "test_user_name_2"},
				},
				{
					ActivityProgram: &ResponceActivityProgram{
						SeqNo:              1,
						Datetime:           "20200101",
						StartTime:          "test_start_time",
						EndTime:            "end_start_time",
						PracticeSectionID:  1,
						PracticeContentsID: 1,
						OutwardTrip:        "outward_trip",
						ReturnTrip:         "return_trip",
						ContactPerson1:     1,
						ContactAbstract1:   "contact_abstract_1",
						ContactPerson2:     1,
						ContactAbstract2:   "contact_abstract_2",
					},
					User: &ResponceUser{Name: "test_user_name_3"},
				},
			},
			"20200102": {
				{
					ActivityProgram: &ResponceActivityProgram{
						SeqNo:              1,
						Datetime:           "20200102",
						StartTime:          "test_start_time",
						EndTime:            "end_start_time",
						PracticeSectionID:  1,
						PracticeContentsID: 1,
						OutwardTrip:        "outward_trip",
						ReturnTrip:         "return_trip",
						ContactPerson1:     1,
						ContactAbstract1:   "contact_abstract_1",
						ContactPerson2:     1,
						ContactAbstract2:   "contact_abstract_2",
					},
					User: &ResponceUser{Name: "test_user_name_1"},
				},
				{
					ActivityProgram: &ResponceActivityProgram{
						SeqNo:              1,
						Datetime:           "20200102",
						StartTime:          "test_start_time",
						EndTime:            "end_start_time",
						PracticeSectionID:  1,
						PracticeContentsID: 1,
						OutwardTrip:        "outward_trip",
						ReturnTrip:         "return_trip",
						ContactPerson1:     1,
						ContactAbstract1:   "contact_abstract_1",
						ContactPerson2:     1,
						ContactAbstract2:   "contact_abstract_2",
					},
					User: &ResponceUser{Name: "test_user_name_2"},
				},
			},
			"20200103": {
				{
					ActivityProgram: &ResponceActivityProgram{
						SeqNo:              1,
						Datetime:           "20200103",
						StartTime:          "test_start_time",
						EndTime:            "end_start_time",
						PracticeSectionID:  1,
						PracticeContentsID: 1,
						OutwardTrip:        "outward_trip",
						ReturnTrip:         "return_trip",
						ContactPerson1:     1,
						ContactAbstract1:   "contact_abstract_1",
						ContactPerson2:     1,
						ContactAbstract2:   "contact_abstract_2",
					},
					User: &ResponceUser{Name: "test_user_name_1"},
				},
			},
		},
	}

	mock := repository.NewDBMock()
	mock.SetUsers([]*model.User{
		{ID: "test_user_id_1", Name: "test_user_name_1"},
		{ID: "test_user_id_2", Name: "test_user_name_2"},
		{ID: "test_user_id_3", Name: "test_user_name_3"},
	})
	mock.SetActivityPrograms([]*model.ActivityProgram{
		{
			UserID:             "test_user_id_1",
			SeqNo:              1,
			Datetime:           "20200101",
			StartTime:          "test_start_time",
			EndTime:            "end_start_time",
			PracticeSectionID:  1,
			PracticeContentsID: 1,
			OutwardTrip:        "outward_trip",
			ReturnTrip:         "return_trip",
			ContactPerson1:     1,
			ContactAbstract1:   "contact_abstract_1",
			ContactPerson2:     1,
			ContactAbstract2:   "contact_abstract_2",
		},
		{
			UserID:             "test_user_id_2",
			SeqNo:              1,
			Datetime:           "20200101",
			StartTime:          "test_start_time",
			EndTime:            "end_start_time",
			PracticeSectionID:  1,
			PracticeContentsID: 1,
			OutwardTrip:        "outward_trip",
			ReturnTrip:         "return_trip",
			ContactPerson1:     1,
			ContactAbstract1:   "contact_abstract_1",
			ContactPerson2:     1,
			ContactAbstract2:   "contact_abstract_2",
		},
		{
			UserID:             "test_user_id_3",
			SeqNo:              1,
			Datetime:           "20200101",
			StartTime:          "test_start_time",
			EndTime:            "end_start_time",
			PracticeSectionID:  1,
			PracticeContentsID: 1,
			OutwardTrip:        "outward_trip",
			ReturnTrip:         "return_trip",
			ContactPerson1:     1,
			ContactAbstract1:   "contact_abstract_1",
			ContactPerson2:     1,
			ContactAbstract2:   "contact_abstract_2",
		},
		{
			UserID:             "test_user_id_1",
			SeqNo:              1,
			Datetime:           "20200102",
			StartTime:          "test_start_time",
			EndTime:            "end_start_time",
			PracticeSectionID:  1,
			PracticeContentsID: 1,
			OutwardTrip:        "outward_trip",
			ReturnTrip:         "return_trip",
			ContactPerson1:     1,
			ContactAbstract1:   "contact_abstract_1",
			ContactPerson2:     1,
			ContactAbstract2:   "contact_abstract_2",
		},
		{
			UserID:             "test_user_id_2",
			SeqNo:              1,
			Datetime:           "20200102",
			StartTime:          "test_start_time",
			EndTime:            "end_start_time",
			PracticeSectionID:  1,
			PracticeContentsID: 1,
			OutwardTrip:        "outward_trip",
			ReturnTrip:         "return_trip",
			ContactPerson1:     1,
			ContactAbstract1:   "contact_abstract_1",
			ContactPerson2:     1,
			ContactAbstract2:   "contact_abstract_2",
		},
		{
			UserID:             "test_user_id_1",
			SeqNo:              1,
			Datetime:           "20200103",
			StartTime:          "test_start_time",
			EndTime:            "end_start_time",
			PracticeSectionID:  1,
			PracticeContentsID: 1,
			OutwardTrip:        "outward_trip",
			ReturnTrip:         "return_trip",
			ContactPerson1:     1,
			ContactAbstract1:   "contact_abstract_1",
			ContactPerson2:     1,
			ContactAbstract2:   "contact_abstract_2",
		},
	})
	repo := repository.NewMockDbRepository(mock)

	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &GetActivityProgramsRequest{}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v; diff=%s", got, want, cmp.Diff(got, want))
	}
}

// func TestGetActivityPrograms_GetActivityPrograms(t *testing.T) {
// 	want

// 	testRouteUserID := model.UserID("test_user_id")
// 	testRouteSeqNo := model.RouteSeqNo(1)
// 	mock := repository.NewDBMock()
// 	mock.SetActivityPrograms([]*model.ActivityProgram{
// 		{

// 		},
// 	})

// 	repo := repository.NewMockDbRepository(mock)
// 	con, _ := repo.NewConnection()
// 	implCtx := NewContext("test_user_id", con, nil)

// 	impl := &GetActivityProgramsRequest{}
// 	_, err := impl.Execute(implCtx)
// 	if err != nil {
// 		t.Fatalf("Is error %#v", err)
// 	}

// 	got, err := con.FindRoute(testRouteUserID, testRouteSeqNo)
// 	if err != nil {
// 		t.Fatalf("Is error %#v", err)
// 	}
// 	if got != nil {
// 		t.Errorf("Find Route is not nil. got=%#v", got)
// 	}
// }
