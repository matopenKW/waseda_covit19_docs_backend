package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type GetRoutesRequest struct {
}

func (r *GetRoutesRequest) SetRequest(form url.Values) {
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
