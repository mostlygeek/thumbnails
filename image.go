package main

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/goadesign/goa"
	"github.com/mostlygeek/thumbnails/app"
	"github.com/xor-gate/goexif2/exif"
)

// ImageController implements the image resource.
type ImageController struct {
	*goa.Controller
	image    []byte
	metadata *exif.Exif
}

func (c *ImageController) LoadImage(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return c.setImage(data)
}

func (c *ImageController) setImage(data []byte) error {
	var err error

	c.image = data

	// load the exif2 data
	buf := bytes.NewReader(data)
	c.metadata, err = exif.Decode(buf)
	if err != nil {
		return err
	}

	return nil
}

// NewImageController creates a image controller.
func NewImageController(service *goa.Service) *ImageController {
	return &ImageController{Controller: service.NewController("ImageController")}
}

// Show runs the show action.
func (c *ImageController) Show(ctx *app.ShowImageContext) error {
	ctx.OK(c.image)
	return nil
}

// Thumbnail runs the thumbnail action.
func (c *ImageController) Thumbnail(ctx *app.ThumbnailImageContext) error {
	// ImageController_Thumbnail: start_implement

	// Put your logic here

	return nil
	// ImageController_Thumbnail: end_implement
}

// Upload runs the upload action.
func (c *ImageController) Upload(ctx *app.UploadImageContext) error {
	// ImageController_Upload: start_implement

	// Put your logic here

	return nil
	// ImageController_Upload: end_implement
}

func (c *ImageController) Metadata(ctx *app.MetadataImageContext) error {
	j, err := c.metadata.MarshalJSON()
	if err != nil {
		return err
	}
	ctx.OK(j)
	return nil
}
