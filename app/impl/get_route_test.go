package impl

import (
	"reflect"
	"testing"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestGetRoute_Execute(t *testing.T) {
	want := &GetRoutesResponce{
		Routes: []*model.Route{
			{
				UserID: "test_user_id",
				SeqNo:  1,
			},
			{
				UserID: "test_user_id",
				SeqNo:  2,
			},
			{
				UserID: "test_user_id",
				SeqNo:  3,
			},
		},
	}

	mock := repository.NewDBMock()
	mock.SetRoutes([]*model.Route{
		{
			UserID:      "test_user_id",
			SeqNo:       1,
			Name:        "test_route_name1",
			OutwardTrip: "test_outward_trip1",
			ReturnTrip:  "test_return_trip1",
		},
		{
			UserID:      "test_user_id",
			SeqNo:       2,
			Name:        "test_route_name2",
			OutwardTrip: "test_outward_trip2",
			ReturnTrip:  "test_return_trip2",
		},
		{
			UserID:      "test_user_id",
			SeqNo:       3,
			Name:        "test_route_name3",
			OutwardTrip: "test_outward_trip3",
			ReturnTrip:  "test_return_trip3",
		},
	})
	repo := repository.NewMockDbRepository(mock)

	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &GetRoutesRequest{}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v", got, want)
	}
}
