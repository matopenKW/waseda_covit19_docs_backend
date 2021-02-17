package impl

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestDeleteRoute_SetRequest(t *testing.T) {
	want := &DeleteRouteRequest{
		RouteID: 1,
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("MarshalJSON errord. err=%#v", err)
	}

	req, _ := http.NewRequest("GET", "/api/v1/delete_put", bytes.NewBuffer(wantJSON))

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
	impl := &DeleteRouteRequest{}
	err := impl.Validate()
	if err != nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestDeleteRoute_Execute(t *testing.T) {
	want := &DeleteRouteResponce{}

	mock := repository.NewDBMock()
	mock.SetRoutes([]*model.Route{
		{
			ID:     1,
			UserID: "test_user_id",
			Name:   "test_route_name",
		},
	})
	repo := repository.NewMockDbRepository(mock)

	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con)

	impl := &DeleteRouteRequest{
		RouteID: 1,
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
	testRouteID := model.RouteID(1)
	mock := repository.NewDBMock()
	mock.SetRoutes([]*model.Route{
		{
			ID:     testRouteID,
			UserID: "test_user_id",
			Name:   "test_route_name",
		},
	})

	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con)

	impl := &DeleteRouteRequest{
		RouteID: 1,
	}
	_, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	got, err := con.FindRoute(testRouteID)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}
	if got != nil {
		t.Errorf("Find Route is not nil. got=%#v", got)
	}
}
