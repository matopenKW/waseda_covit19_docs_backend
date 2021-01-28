package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type GetRoutesRequest struct {
	UserID string
}

func (r *GetRoutesRequest) SetRequest(form url.Values) {
	r.UserID = "user_id"
}

func (r *GetRoutesRequest) Validate() error {
	return nil
}

func (r *GetRoutesRequest) Execute(con repository.Connection) (ResponceImpl, error) {
	return GetRoutes(r, con)
}

type GetRoutesResponce struct {
	Routes []*model.Route
}

func (r *GetRoutesResponce) GetResponce() {
}

func GetRoutes(req *GetRoutesRequest, con repository.Connection) (ResponceImpl, error) {
	rs, err := con.FindRoutesByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	return &GetRoutesResponce{
		Routes: rs,
	}, nil
}
