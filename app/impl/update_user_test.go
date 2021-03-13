package impl

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

func TestUpdateUser_SetRequest(t *testing.T) {
	want := &UpdateUserRequest{
		Name:           "TEST_NAME",
		UniversityName: "TEST_UNIVERSITY_NAME",
		FacultyName:    "TEST_FACULTY_NAME",
		StudentNo:      "TEST_STUDENT_NO",
		CellPhonNo:     "TEST_CELL_PHON_NO",
	}

	bs := []byte(`{
		"name": "TEST_NAME",
		"university_name": "TEST_UNIVERSITY_NAME",
		"faculty_name": "TEST_FACULTY_NAME",
		"student_no": "TEST_STUDENT_NO",
		"cell_phon_no": "TEST_CELL_PHON_NO"
		}`)
	req, _ := http.NewRequest("GET", "/api/v1/update_user", bytes.NewBuffer(bs))

	var ctx *gin.Context
	ctx = &gin.Context{
		Request: req,
	}

	impl := &UpdateUserRequest{}
	impl.SetRequest(ctx)

	if !reflect.DeepEqual(want, impl) {
		t.Errorf("SetRequest is fatal. impl=%#v, want=%#v", impl, want)
	}
}

func TestUpdateUser_Validate(t *testing.T) {
	impl := &UpdateUserRequest{
		Name:           "TEST_NAME",
		UniversityName: "TEST_UNIVERSITY_NAME",
		FacultyName:    "TEST_FACULTY_NAME",
		StudentNo:      "TEST_STUDENT_NO",
		CellPhonNo:     "TEST_CELL_PHON_NO",
	}
	err := impl.Validate()
	if err != nil {
		t.Fatalf("Validate errord. target=%#v, err=%#v", impl, err)
	}
}

func TestUpdateUser_Execute(t *testing.T) {
	want := &UpdateUserResponce{}

	mock := repository.NewDBMock()
	mock.SetUsers([]*model.User{
		{ID: "TEST_USER_ID"},
	})

	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("TEST_USER_ID", con, nil)

	impl := &UpdateUserRequest{
		Name:           "TEST_NAME",
		UniversityName: "TEST_UNIVERSITY_NAME",
		FacultyName:    "TEST_FACULTY_NAME",
		StudentNo:      "TEST_STUDENT_NO",
		CellPhonNo:     "TEST_CELL_PHON_NO",
	}
	got, err := impl.Execute(implCtx)
	if err != nil {
		t.Fatalf("Is error %#v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Execute is fatal. got=%#v, want=%#v", got, want)
	}
}

func TestUpdateUser_Execute_Fail(t *testing.T) {
	wantErr := fmt.Errorf("user is empty; user_id=TEST_USER_ID")

	mock := repository.NewDBMock()
	repo := repository.NewMockDbRepository(mock)
	con, _ := repo.NewConnection()
	implCtx := NewContext("TEST_USER_ID", con, nil)

	impl := &UpdateUserRequest{
		Name:           "TEST_NAME",
		UniversityName: "TEST_UNIVERSITY_NAME",
		FacultyName:    "TEST_FACULTY_NAME",
		StudentNo:      "TEST_STUDENT_NO",
		CellPhonNo:     "TEST_CELL_PHON_NO",
	}
	_, err := impl.Execute(implCtx)
	if err == nil {
		t.Fatalf("Is error %#v", err)
	}
	if wantErr.Error() != err.Error() {
		t.Errorf("Invalid Error msg. wantErr=%v, err=%v", wantErr, err)
	}
}
