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
	RouteID     int    `json:"route_id"`
	Name        string `json:"name"`
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
	Route *model.Route
}

// GetResponce is responce get receiver
func (r *PutRouteResponce) GetResponce() {
}

func putRoute(req *PutRouteRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()

	fmt.Println(req)

	var route *model.Route
	var err error
	err = con.RunTransaction(func(tx repository.Transaction) error {
		var rID model.RouteID
		if req.RouteID == 0 {
			maxID, err := con.FindMaxRouteID()
			if err != nil {
				return err
			}
			rID = maxID + 1
		} else {
			rID = model.RouteID(req.RouteID)
		}
		route, err = tx.SaveRoute(&model.Route{
			ID:          rID,
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
