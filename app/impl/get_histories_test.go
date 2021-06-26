package impl

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestGetHistoriesService_Execute(t *testing.T) {
	want := &GetHistoriesResponce{
		Histories: []*History{
			{
				Month: "10",
				ActivityPrograms: []*ResponceActivityProgram{
					{
						SeqNo:    1,
						Datetime: "20201001",
					},
					{
						SeqNo:    2,
						Datetime: "20201002",
					},
					{
						SeqNo:    3,
						Datetime: "20201003",
					},
				},
			},
			{
				Month: "11",
				ActivityPrograms: []*ResponceActivityProgram{
					{
						SeqNo:    4,
						Datetime: "20201101",
					},
				},
			},
		},
	}

	mock := repository.NewDBMock()
	mock.SetActivityPrograms([]*model.ActivityProgram{
		{
			UserID:   "test_user_id",
			SeqNo:    1,
			Datetime: "20201001",
		},
		{
			UserID:   "test_user_id",
			SeqNo:    2,
			Datetime: "20201002",
		},
		{
			UserID:   "test_user_id",
			SeqNo:    3,
			Datetime: "20201003",
		},
		{
			UserID:   "test_user_id",
			SeqNo:    4,
			Datetime: "20201101",
		},
	})
	repo := repository.NewMockDbRepository(mock)

	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &GetHistoriesRequest{}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v, diff=%s", got, want, cmp.Diff(got, want))
	}
}
