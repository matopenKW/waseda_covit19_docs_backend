package impl

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/model"
	"github.com/matopenKW/waseda_covit19_docs_backend/app/repository"
)

// DeleteRouteService is delete route service
type DeleteRouteService struct{}

// New is delete route service new
func (s *DeleteRouteService) New() RequestImpl {
	return &DeleteRouteRequest{}
}

// DeleteRouteRequest is delete route request
type DeleteRouteRequest struct {
	RouteID *int `json:"route_id,omitempty"`
}

// SetRequest is request set receiver
func (r *DeleteRouteRequest) SetRequest(ctx *gin.Context) {
	_ = ctx.ShouldBindJSON(&r)
}

// Validate is validate receiver
func (r *DeleteRouteRequest) Validate() error {
	log.Println(r)
	if r.RouteID == nil {
		return fmt.Errorf("Invalid route id")
	}
	return nil
}

// Execute is api execute receiver
func (r *DeleteRouteRequest) Execute(ctx *Context) (ResponceImpl, error) {
	return deleteRoute(r, ctx)
}

// DeleteRouteResponce is put route responce
type DeleteRouteResponce struct {
}

// GetResponce is responce get receiver
func (r *DeleteRouteResponce) GetResponce() {
}

// deleteRoute is api execute function of private
func deleteRoute(req *DeleteRouteRequest, ctx *Context) (ResponceImpl, error) {
	con := ctx.GetConnection()
	err := con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.DeleteRoute(model.RouteID(*req.RouteID))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &DeleteRouteResponce{}, nil
}
