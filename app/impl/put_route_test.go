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

func TestPutRoute_SetRequest(t *testing.T) {
	want := &PutRouteRequest{
		RouteID:     1,
		Name:        "test_name",
		OutwardTrip: "test_outward_trip",
		ReturnTrip:  "test_return_trip",
	}

	bs := []byte(`{
		"route_id": 1,
		"name": "test_name", 
		"outward_trip": "test_outward_trip",
		"return_trip": "test_return_trip"
		}`)
	req, _ := http.NewRequest("GET", "/api/v1/put_put", bytes.NewBuffer(bs))

	var ctx *gin.Context
	ctx = &gin.Context{
		Request: req,
	}

	impl := &PutRouteRequest{}
	impl.SetRequest(ctx)

	if !reflect.DeepEqual(want, impl) {
		t.Errorf("SetRequest is fatal. impl=%#v, want=%#v", impl, want)
	}
}

func TestPutRoute_Validate(t *testing.T) {
	impl := &PutRouteRequest{
		Name: "test_name",
	}
	err := impl.Validate()
	if err != nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestPutRoute_Validate_Fail(t *testing.T) {
	impl := &PutRouteRequest{}

	err := impl.Validate()
	if err == nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestPutRoute_Execute(t *testing.T) {
	want := &PutRouteResponce{
		Route: &model.Route{
			ID:          1,
			UserID:      "test_user_id",
			Name:        "test_route_name",
			OutwardTrip: "test_outward_trip",
			ReturnTrip:  "test_return_trip",
		}}

	mock := repository.NewDBMock()
	repo := repository.NewMockDbRepository(mock)

	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &PutRouteRequest{
		RouteID:     1,
		Name:        "test_route_name",
		OutwardTrip: "test_outward_trip",
		ReturnTrip:  "test_return_trip",
	}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v", got, want)
	}
}

func TestPutRoute_RegisterRoute(t *testing.T) {
	want := &model.Route{
		ID:          1,
		UserID:      "test_user_id",
		Name:        "test_route_name",
		OutwardTrip: "test_outward_trip",
		ReturnTrip:  "test_return_trip",
	}

	testRouteID := model.RouteID(1)
	mock := repository.NewDBMock()
	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &PutRouteRequest{
		RouteID:     1,
		Name:        "test_route_name",
		OutwardTrip: "test_outward_trip",
		ReturnTrip:  "test_return_trip",
	}
	_, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	got, err := con.FindRoute(testRouteID)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Find Route is error. want=%#v, got=%#v", want, got)
	}
}

func TestPutRoute_UpdateRoute(t *testing.T) {
	want := &model.Route{
		ID:          1,
		UserID:      "test_user_id",
		Name:        "test_update_route_name",
		OutwardTrip: "test_update_outward_trip",
		ReturnTrip:  "test_update_return_trip",
	}

	testRouteID := model.RouteID(1)
	mock := repository.NewDBMock()
	mock.SetRoutes([]*model.Route{
		{
			ID:          testRouteID,
			Name:        "test_name",
			OutwardTrip: "test_outward_trip",
			ReturnTrip:  "test_return_trip",
		},
	})

	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &PutRouteRequest{
		RouteID:     1,
		Name:        "test_update_route_name",
		OutwardTrip: "test_update_outward_trip",
		ReturnTrip:  "test_update_return_trip",
	}
	_, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	got, err := con.FindRoute(testRouteID)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Find Route is error. want=%#v, got=%#v", want, got)
	}
}
