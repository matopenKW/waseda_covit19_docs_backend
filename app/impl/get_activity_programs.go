package impl

import (
	"github.com/gin-gonic/gin"

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

type ResponceActivityPrograms struct {
	ActivityProgramMap map[string][]*ResponceActivityProgram
}

type GetActivityProgramsResponce struct {
	ActivityPrograms *ResponceActivityPrograms
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

	result := make(map[string][]*ResponceActivityProgram)
	for _, ap := range aps {
		if _, exsits := result[ap.Datetime]; !exsits {
			result[ap.Datetime] = make([]*ResponceActivityProgram, 0)
		}
		result[ap.Datetime] = append(result[ap.Datetime], PresenterActivityProgram(ap))
	}

	return &GetActivityProgramsResponce{
		ActivityPrograms: &ResponceActivityPrograms{
			ActivityProgramMap: result,
		},
	}, nil
}
