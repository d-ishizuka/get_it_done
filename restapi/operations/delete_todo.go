// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteTodoHandlerFunc turns a function with the right signature into a delete todo handler
type DeleteTodoHandlerFunc func(DeleteTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTodoHandlerFunc) Handle(params DeleteTodoParams) middleware.Responder {
	return fn(params)
}

// DeleteTodoHandler interface for that can handle valid delete todo params
type DeleteTodoHandler interface {
	Handle(DeleteTodoParams) middleware.Responder
}

// NewDeleteTodo creates a new http.Handler for the delete todo operation
func NewDeleteTodo(ctx *middleware.Context, handler DeleteTodoHandler) *DeleteTodo {
	return &DeleteTodo{Context: ctx, Handler: handler}
}

/*
	DeleteTodo swagger:route DELETE /todos/{id} deleteTodo

Delete a ToDo item by ID
*/
type DeleteTodo struct {
	Context *middleware.Context
	Handler DeleteTodoHandler
}

func (o *DeleteTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTodoParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
