package impl

import (
	"net/url"

	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
)

type PutRouteRequest struct {
}

func (r *PutRouteRequest) SetRequest(form url.Values) {
}

func (r *PutRouteRequest) Validate() error {
	return nil
}

func (r *PutRouteRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return PutRoute(r, ctx)
}

type PutRouteResponce struct {
	Route *model.Route
}

func (r *PutRouteResponce) GetResponce() {
}

func PutRoute(req *PutRouteRequest, ctx *Context) (ResponceImpl, error) {
	return &PutRouteResponce{
		Route: &model.Route{
			ID:          1,
			UserID:      "user_id",
			OutwardTrip: "町田→新宿→（中央線）→東京",
			ReturnTrip:  "東京→新宿→町田",
		},
	}, nil
}
