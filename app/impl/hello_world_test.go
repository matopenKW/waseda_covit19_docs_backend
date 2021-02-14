package impl

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestHelloWorld_SetRequest(t *testing.T) {
	want := HelloWorldRequest{
		Message: "test_message",
	}
	wantJSON, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("MarshalJSON errord. err=%#v", err)
	}

	req, _ := http.NewRequest("GET", "/api/v1/hello_world", bytes.NewBuffer(wantJSON))

	var ctx *gin.Context
	ctx = &gin.Context{
		Request: req,
	}

	impl := &HelloWorldRequest{}
	impl.SetRequest(ctx)

	if reflect.DeepEqual(want, impl) {
		t.Errorf("SetRequest is fatal. impl=%#v, want=%#v", impl, want)
	}
}

func TestHelloWorld_Validate(t *testing.T) {
	impl := &HelloWorldRequest{
		Message: "test_Message",
	}

	err := impl.Validate()
	if err != nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestHelloWorld_helloWorld(t *testing.T) {
	want := &HelloWorldResponce{
		Message: "hello world test_message",
	}

	dbMock := repository.NewMockDbRepository()
	con, _ := dbMock.NewConnection()
	implCtx := NewContext("test_user_id", con)

	impl := &HelloWorldRequest{
		Message: "test_message",
	}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v", got, want)
	}
}
