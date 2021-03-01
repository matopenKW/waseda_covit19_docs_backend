package impl

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestPutUser_SetRequest(t *testing.T) {
	want := &PutUserRequest{
		UID:   "TEST_UID",
		Email: "test@example.com",
	}

	bs := []byte(`{
		"uid": "TEST_UID",
		"email": "test@example.com" 
		}`)
	req, _ := http.NewRequest("GET", "/api/v1/put_user", bytes.NewBuffer(bs))

	var ctx *gin.Context
	ctx = &gin.Context{
		Request: req,
	}

	impl := &PutUserRequest{}
	impl.SetRequest(ctx)

	if !reflect.DeepEqual(want, impl) {
		t.Errorf("SetRequest is fatal. impl=%#v, want=%#v", impl, want)
	}
}

func TestPutUser_Validate(t *testing.T) {
	impl := &PutUserRequest{
		UID:   "TEST_UID",
		Email: "test@example.com",
	}
	err := impl.Validate()
	if err != nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestPutUser_Validate_Fail(t *testing.T) {
	wantErr := fmt.Errorf("Invalid uid. uid is blank")
	impl := &PutUserRequest{}

	err := impl.Validate()
	if err == nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
	if wantErr.Error() != err.Error() {
		t.Errorf("Invalid Error msg. wantErr=%v, err=%v", wantErr, err)
	}

	impl = &PutUserRequest{
		UID: "TEST_UID",
	}
	wantErr = fmt.Errorf("Invalid email. email is blank")

	err = impl.Validate()
	if err == nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
	if wantErr.Error() != err.Error() {
		t.Errorf("Invalid Error msg. wantErr=%v, err=%v", wantErr, err)
	}

}

func TestPutUser_Execute(t *testing.T) {
	want := &PutUserResponce{}

	mock := repository.NewDBMock()
	repo := repository.NewMockDbRepository(mock)

	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &PutUserRequest{
		UID:   "TEST_UID",
		Email: "test@example.com",
	}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v", got, want)
	}
}

func TestPutUser_CreateUser(t *testing.T) {
	want := &model.User{
		ID:    "TEST_USER_ID",
		Email: "test@example.com",
	}

	mock := repository.NewDBMock()
	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &PutUserRequest{
		UID:   "TEST_USER_ID",
		Email: "test@example.com",
	}
	_, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	us, err := con.ListUser()
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	got := us[0]
	if !reflect.DeepEqual(want, got) {
		t.Errorf("Find User is error. want=%#v, got=%#v, diff=%s", got, want, cmp.Diff(got, want))
	}
}

func TestPutUser_CreateUser_Fail(t *testing.T) {
	mock := repository.NewDBMock()
	mock.SetUsers([]*model.User{
		{ID: model.UserID("TEST_USER_ID"), Email: "test@example.com"},
	})
	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("test_user_id", con, nil)

	impl := &PutUserRequest{
		UID:   "TEST_USER_ID",
		Email: "test@example.com",
	}
	_, err := impl.Execute(implCtx)
	if err == nil {
		t.Fatalf("Is Not error")
	}

}
