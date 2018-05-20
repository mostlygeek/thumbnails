//go:generate goagen bootstrap -d github.com/mostlygeek/thumbnails/design

package main

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/mostlygeek/thumbnails/app"
)

func main() {
	// Create service
	service := goa.New("thumbnail")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "image" controller
	c := NewImageController(service)
	if err := c.LoadImage("./tart.jpg"); err != nil {
		fmt.Println("Could not load default image")
		return
	}

	app.MountImageController(service, c)
	// Mount "ui" controller
	c2 := NewUIController(service)
	app.MountUIController(service, c2)

	// Start service
	if err := service.ListenAndServe("localhost:8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
