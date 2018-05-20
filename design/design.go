package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("thumbnail", func() {
	Description("Generate thumbnails to learn goa's DSL")
	BasePath("/")
})

var _ = Resource("ui", func() {
	Action("show", func() {
		Routing(GET("/"))
		Description("Serve the HTML UI")
		Response(OK, "text/html")
	})
})

var _ = Resource("image", func() {
	BasePath("image")
	Action("upload", func() {
		Routing(POST("/"), PUT("/"))
		Description("Update the image")
		Response(OK)
		Response(BadRequest)
	})

	Action("show", func() {
		Routing(GET("/"))
		Description("GET the current image being handled")
		Response(OK, "image/jpeg")
	})

	Action("metadata", func() {
		Routing(GET("/metadata"))
		Description("Extract EXIF metadata")
		Response(OK, "application/json")
	})

	Action("thumbnail", func() {
		Routing(GET("/thumbnail/:type"))
		Params(func() {
			Param("type", String, "Size of thumbnail to generate", func() {
				Enum("small", "medium", "large", "square", "box")
			})
			Param("x", Integer, "Width of bounding box")
			Param("y", Integer, "Height of bounding box")
		})
		Response(OK)
		Response(BadRequest) // if missing X or Y when box = height
	})
})
