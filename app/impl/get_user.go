package impl

import (
	"github.com/gin-gonic/gin"
)

// GetUserService is get user service
type GetUserService struct{}

// New is get user service new
func (s *GetUserService) New() RequestImpl {
	return &GetUserRequest{}
}

type GetUserRequest struct {
}

func (r *GetUserRequest) SetRequest(ctx *gin.Context) {
}

func (r *GetUserRequest) Validate() error {
	return nil
}

func (r *GetUserRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return GetUser(r, ctx)
}

type ResponceUser struct {
	Name           string
	UniversityName string
	FacultyName    string
	StudentNo      string
	CellPhonNo     string
}

type GetUserResponce struct {
	User *ResponceUser
}

func (r *GetUserResponce) GetResponce() {
}

func GetUser(req *GetUserRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	u, err := con.FindUser(ctx.userID)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return &GetUserResponce{
			User: &ResponceUser{},
		}, nil
	}

	return &GetUserResponce{
		User: &ResponceUser{
			Name:           u.Name,
			UniversityName: u.UniversityName,
			FacultyName:    u.FacultyName,
			StudentNo:      u.StudentNo,
			CellPhonNo:     u.CellPhonNo,
		},
	}, nil
}
