package impl

import (
	"bytes"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestDeleteRoute_SetRequest(t *testing.T) {
	seq := 1
	want := &DeleteRouteRequest{
		SeqNo: &seq,
	}

	bs := []byte(`{"seq_no":1}`)
	req, _ := http.NewRequest("GET", "/api/v1/delete_put", bytes.NewBuffer(bs))

	var ctx *gin.Context
	ctx = &gin.Context{
		Request: req,
	}

	impl := &DeleteRouteRequest{}
	impl.SetRequest(ctx)

	if !reflect.DeepEqual(want, impl) {
		t.Errorf("SetRequest is fatal. impl=%#v, want=%#v", impl, want)
	}
}

func TestDeleteRoute_Validate(t *testing.T) {
	rID := 1
	impl := &DeleteRouteRequest{
		SeqNo: &rID,
	}
	err := impl.Validate()
	if err != nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestDeleteRoute_Validate_Fail(t *testing.T) {
	impl := &DeleteRouteRequest{}
	err := impl.Validate()
	if err == nil {
		t.Fatalf("Validate not errord. target=%#v, err=%#v", impl, err)
	}
}

func TestDeleteRoute_Execute(t *testing.T) {
	want := &DeleteRouteResponce{}

	mock := repository.NewDBMock()
	mock.SetRoutes([]*model.Route{
		{
			UserID: "test_user_id",
			Name:   "test_route_name",
		},
	})
	repo := repository.NewMockDbRepository(mock)

	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	rID := 1
	impl := &DeleteRouteRequest{
		SeqNo: &rID,
	}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v", got, want)
	}
}

func TestDeleteRoute_DeleteRoute(t *testing.T) {
	testRouteUserID := "test_user_id"
	testRouteSeqNo := model.RouteSeqNo(1)
	mock := repository.NewDBMock()
	mock.SetRoutes([]*model.Route{
		{
			SeqNo:  testRouteSeqNo,
			UserID: "test_user_id",
			Name:   "test_route_name",
		},
	})

	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	rID := 1
	impl := &DeleteRouteRequest{
		SeqNo: &rID,
	}
	_, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	got, err := con.FindRoute(testRouteUserID, testRouteSeqNo)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}
	if got != nil {
		t.Errorf("Find Route is not nil. got=%#v", got)
	}
}
