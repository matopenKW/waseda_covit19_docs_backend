package impl

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// PutRouteService is put route service
type PutRouteService struct{}

// New is put route service new
func (s *PutRouteService) New() RequestImpl {
	return &PutRouteRequest{}
}

// PutRouteRequest is put route request
type PutRouteRequest struct {
	Name        string `json:"name"`
	SeqNo       *int   `json:"seq_no, omitempty"`
	OutwardTrip string `json:"outward_trip"`
	ReturnTrip  string `json:"return_trip"`
}

// SetRequest is request set receiver
func (r *PutRouteRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

// Validate is validate receiver
func (r *PutRouteRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("Invalid Name. Name is blank")
	}
	return nil
}

// Execute is api execute receiver
func (r *PutRouteRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return putRoute(r, ctx)
}

// PutRouteResponce is put route responce
type PutRouteResponce struct {
	Route *ResponceRoute
}

// GetResponce is responce get receiver
func (r *PutRouteResponce) GetResponce() {
}

func putRoute(req *PutRouteRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	var err error
	var route *model.Route
	if req.SeqNo == nil {
		maxSeq, err := con.FindRouteMaxSeqNo(ctx.userID)
		if err != nil {
			return nil, err
		}

		route, err = createRoute(ctx, &model.Route{
			UserID:      ctx.userID,
			SeqNo:       maxSeq + 1,
			Name:        req.Name,
			OutwardTrip: req.OutwardTrip,
			ReturnTrip:  req.ReturnTrip,
		})
		if err != nil {
			return nil, err
		}

	} else {
		route, err = updateRoute(ctx, &model.Route{
			UserID:      ctx.userID,
			SeqNo:       model.RouteSeqNo(*req.SeqNo),
			Name:        req.Name,
			OutwardTrip: req.OutwardTrip,
			ReturnTrip:  req.ReturnTrip,
		})
		if err != nil {
			return nil, err
		}

	}

	return &PutRouteResponce{
		Route: PresenterRoute(route),
	}, nil
}

func createRoute(ctx *Context, r *model.Route) (*model.Route, error) {
	con := ctx.GetConnection()

	var err error
	var route *model.Route
	err = con.RunTransaction(func(tx repository.Transaction) error {
		route, err = tx.CreateRoute(r)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return route, nil
}

func updateRoute(ctx *Context, r *model.Route) (*model.Route, error) {
	con := ctx.GetConnection()

	var err error
	var route *model.Route
	err = con.RunTransaction(func(tx repository.Transaction) error {
		route, err = tx.UpdateRoute(r)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return route, nil
}
