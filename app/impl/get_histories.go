package impl

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

// GetHistoriesService is get histories service
type GetHistoriesService struct{}

// New is get histories service new
func (s *GetHistoriesService) New() RequestImpl {
	return &GetHistoriesRequest{}
}

// GetHistoriesRequest is get histories request
type GetHistoriesRequest struct {
}

// SetRequest is request set receiver
func (r *GetHistoriesRequest) SetRequest(ctx *gin.Context) {
}

// Validate is validate receiver
func (r *GetHistoriesRequest) Validate() error {
	return nil
}

// Execute is api execute receiver
func (r *GetHistoriesRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return getHistories(r, ctx)
}

type History struct {
	Month            string
	ActivityPrograms []*model.ActivityProgram
}

// GetHistoriesResponce is put histories responce
type GetHistoriesResponce struct {
	Histories []*History
}

// GetResponce is responce get receiver
func (r *GetHistoriesResponce) GetResponce() {
}

func getHistories(req *GetHistoriesRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	aps, err := con.ListActivityProgramsByUserID(ctx.userID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	hm := make(map[string][]*model.ActivityProgram)
	for _, v := range aps {
		month := v.Datetime[4:6]
		hm[month] = append(hm[month], v)
	}

	hs := make([]*History, 0)
	for k, v := range hm {
		h := &History{
			Month:            k,
			ActivityPrograms: v,
		}
		hs = append(hs, h)
	}

	return &GetHistoriesResponce{
		Histories: hs,
	}, nil
}
