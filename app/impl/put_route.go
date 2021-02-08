package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

type PutRouteRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OutwardTrip string `json:"outward_trip"`
	ReturnTrip  string `json:"return_trip"`
}

func (r *PutRouteRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
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
	con := ctx.GetConnection()

	var err error
	var route *model.Route
	err = con.RunTransaction(func(tx repository.Transaction) error {
		route, err = tx.SaveRoute(&model.Route{
			UserID:      ctx.userID,
			Name:        req.Name,
			OutwardTrip: req.OutwardTrip,
			ReturnTrip:  req.ReturnTrip,
		})

		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &PutRouteResponce{
		Route: route,
	}, nil
}
