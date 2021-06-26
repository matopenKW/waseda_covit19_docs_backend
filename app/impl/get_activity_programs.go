package impl

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// GetActivityProgramsService is get activity programs service
type GetActivityProgramsService struct {
	Datetime string `json:"datetime"`
}

// New is get activity programs service new
func (s *GetActivityProgramsService) New() RequestImpl {
	return &GetActivityProgramsRequest{}
}

type GetActivityProgramsRequest struct {
}

func (r *GetActivityProgramsRequest) SetRequest(ctx *gin.Context) {
}

func (r *GetActivityProgramsRequest) Validate() error {
	return nil
}

func (r *GetActivityProgramsRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return GetActivityPrograms(r, ctx)
}

type HistoryForEachUser struct {
	ActivityProgram *ResponceActivityProgram
	User            *ResponceUser
}

type GetActivityProgramsResponce struct {
	ActivityPrograms map[string][]*HistoryForEachUser
}

func (r *GetActivityProgramsResponce) GetResponce() {
}

func GetActivityPrograms(req *GetActivityProgramsRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	aps, err := con.ListActivityPrograms(repository.ActivityProgramFilter{
		OrderBy: repository.ActivityProgramOrderByDatetimeAsc,
	})
	if err != nil {
		return nil, err
	}

	us, err := con.ListUser()
	if err != nil {
		return nil, err
	}

	uMap := make(map[model.UserID]*model.User)
	for _, u := range us {
		uMap[u.ID] = u
	}

	duplicate := make(map[string]int)
	result := make(map[string][]*HistoryForEachUser)
	for _, ap := range aps {
		if _, ok := duplicate[fmt.Sprintf("%s-%s", ap.UserID, ap.Datetime)]; ok {
			continue
		}

		if _, exsits := result[ap.Datetime]; !exsits {
			result[ap.Datetime] = make([]*HistoryForEachUser, 0)
		}

		hfeu := &HistoryForEachUser{
			ActivityProgram: PresenterActivityProgram(ap),
		}

		if u, ok := uMap[ap.UserID]; ok {
			hfeu.User = PresenterUser(u)
		}
		result[ap.Datetime] = append(result[ap.Datetime], hfeu)
		duplicate[fmt.Sprintf("%s-%s", ap.UserID, ap.Datetime)] = 0
	}

	return &GetActivityProgramsResponce{
		ActivityPrograms: result,
	}, nil
}
