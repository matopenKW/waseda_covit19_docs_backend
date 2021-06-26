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

type ResponceRoute struct {
	SeqNo       int
	Name        string
	OutwardTrip string
	ReturnTrip  string
}

type GetRoutesResponce struct {
	Routes []*ResponceRoute
	Places []*model.Place
}

func (r *GetRoutesResponce) GetResponce() {
}

func GetRoutes(req *GetRoutesRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	rs, err := con.FindRoutesByUserID(ctx.GetUserID())
	if err != nil {
		return nil, err
	}

	routes := make([]*ResponceRoute, 0)
	for _, r := range rs {
		routes = append(routes, PresenterRoute(r))
	}
	return &GetRoutesResponce{
		Routes: routes,
		Places: ctx.master.places,
	}, nil
}

func PresenterRoute(r *model.Route) *ResponceRoute {
	return &ResponceRoute{
		SeqNo:       int(r.SeqNo),
		Name:        r.Name,
		OutwardTrip: r.OutwardTrip,
		ReturnTrip:  r.ReturnTrip,
	}
}
