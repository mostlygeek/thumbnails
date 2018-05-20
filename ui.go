package main

import (
	"github.com/goadesign/goa"
	"github.com/mostlygeek/thumbnails/app"
)

// UIController implements the ui resource.
type UIController struct {
	*goa.Controller
}

// NewUIController creates a ui controller.
func NewUIController(service *goa.Service) *UIController {
	return &UIController{Controller: service.NewController("UIController")}
}

// Show runs the show action.
func (c *UIController) Show(ctx *app.ShowUIContext) error {
	ctx.OK([]byte(`<html><body>
	<h2>Current Image</h2>
	<img src="/image"/>

	<h2>Thumbnails</h2>
	small, medium, large, box ...

	</body></html`))
	return nil
}
