package resources

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"revised-server/models"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetResourcesIDHandlerFunc turns a function with the right signature into a get resources ID handler
type GetResourcesIDHandlerFunc func(GetResourcesIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResourcesIDHandlerFunc) Handle(params GetResourcesIDParams) middleware.Responder {
	return fn(params)
}

// GetResourcesIDHandler interface for that can handle valid get resources ID params
type GetResourcesIDHandler interface {
	Handle(GetResourcesIDParams) middleware.Responder
}

// NewGetResourcesID creates a new http.Handler for the get resources ID operation
func NewGetResourcesID(ctx *middleware.Context, handler GetResourcesIDHandler) *GetResourcesID {
	return &GetResourcesID{Context: ctx, Handler: handler}
}

/*GetResourcesID swagger:route GET /resources/{id} Resources getResourcesId

Resources

Returns the single resource with the matching ID.

*/
type GetResourcesID struct {
	Context *middleware.Context
	Handler GetResourcesIDHandler
}

func (o *GetResourcesID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetResourcesIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetResourcesIDOKBody get resources ID o k body

swagger:model GetResourcesIDOKBody
*/
type GetResourcesIDOKBody struct {

	/* resource

	Required: true
	*/
	Resource *models.Resource `json:"resource"`
}
