// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTodosHandlerFunc turns a function with the right signature into a get todos handler
type GetTodosHandlerFunc func(GetTodosParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTodosHandlerFunc) Handle(params GetTodosParams) middleware.Responder {
	return fn(params)
}

// GetTodosHandler interface for that can handle valid get todos params
type GetTodosHandler interface {
	Handle(GetTodosParams) middleware.Responder
}

// NewGetTodos creates a new http.Handler for the get todos operation
func NewGetTodos(ctx *middleware.Context, handler GetTodosHandler) *GetTodos {
	return &GetTodos{Context: ctx, Handler: handler}
}

/*
	GetTodos swagger:route GET /todos getTodos

Get all ToDo items
*/
type GetTodos struct {
	Context *middleware.Context
	Handler GetTodosHandler
}

func (o *GetTodos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTodosParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
