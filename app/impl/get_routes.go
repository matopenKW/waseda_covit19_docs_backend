package impl

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

// GetRoutesService is get routes service
type GetRoutesService struct{}

// New is get routes service new
func (s *GetRoutesService) New() RequestImpl {
	return &GetRoutesRequest{}
}

type GetRoutesRequest struct {
	SeqNo int `json:"seq_no"`
}

func (r *GetRoutesRequest) SetRequest(ctx *gin.Context) {
	r.SeqNo, _ = strconv.Atoi(ctx.Param("seq_no"))
}

func (r *GetRoutesRequest) Validate() error {
	if r.SeqNo == 0 {
		return fmt.Errorf("Invalid seq no")
	}
	return nil
}

func (r *GetRoutesRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return GetRoutes(r, ctx)
}

type GetRoutesResponce struct {
	Routes *model.Route
}

func (r *GetRoutesResponce) GetResponce() {
}

func GetRoutes(req *GetRoutesRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	rs, err := con.FindRoute(ctx.userID, model.RouteSeqNo(req.SeqNo))
	if err != nil {
		return nil, err
	}

	if rs == nil {
		return &GetRoutesResponce{
			Routes: nil,
		}, nil
	}

	return &GetRoutesResponce{
		Routes: rs,
	}, nil
}
