package impl

import (
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
}

func (r *GetRoutesRequest) SetRequest(ctx *gin.Context) {
}

func (r *GetRoutesRequest) Validate() error {
	return nil
}

func (r *GetRoutesRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return GetRoutes(r, ctx)
}

type GetRoutesResponce struct {
	Routes []*model.Route
}

func (r *GetRoutesResponce) GetResponce() {
}

func GetRoutes(req *GetRoutesRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	rs, err := con.FindRoutesByUserID(ctx.GetUserID())
	if err != nil {
		return nil, err
	}

	return &GetRoutesResponce{
		Routes: rs,
	}, nil
}
