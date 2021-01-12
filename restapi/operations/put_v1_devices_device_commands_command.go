// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/momper14/edgex-restapp/models"
)

// PutV1DevicesDeviceCommandsCommandHandlerFunc turns a function with the right signature into a put v1 devices device commands command handler
type PutV1DevicesDeviceCommandsCommandHandlerFunc func(PutV1DevicesDeviceCommandsCommandParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn PutV1DevicesDeviceCommandsCommandHandlerFunc) Handle(params PutV1DevicesDeviceCommandsCommandParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// PutV1DevicesDeviceCommandsCommandHandler interface for that can handle valid put v1 devices device commands command params
type PutV1DevicesDeviceCommandsCommandHandler interface {
	Handle(PutV1DevicesDeviceCommandsCommandParams, *models.User) middleware.Responder
}

// NewPutV1DevicesDeviceCommandsCommand creates a new http.Handler for the put v1 devices device commands command operation
func NewPutV1DevicesDeviceCommandsCommand(ctx *middleware.Context, handler PutV1DevicesDeviceCommandsCommandHandler) *PutV1DevicesDeviceCommandsCommand {
	return &PutV1DevicesDeviceCommandsCommand{Context: ctx, Handler: handler}
}

/*PutV1DevicesDeviceCommandsCommand swagger:route PUT /v1/devices/{device}/commands/{command} putV1DevicesDeviceCommandsCommand

Issue the PUT command referenced by the command name to the device/sensor.

*/
type PutV1DevicesDeviceCommandsCommand struct {
	Context *middleware.Context
	Handler PutV1DevicesDeviceCommandsCommandHandler
}

func (o *PutV1DevicesDeviceCommandsCommand) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutV1DevicesDeviceCommandsCommandParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}