package main

import (
	"bytes"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"

	"github.com/disintegration/imaging"
	"github.com/goadesign/goa"
	"github.com/mostlygeek/thumbnails/app"
	"github.com/pkg/errors"
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
	img, err := imaging.Decode(bytes.NewReader(c.image))
	if err != nil {
		return nil
	}
	max := img.Bounds().Max

	var x, y int

	switch ctx.Type {
	case "small": // 1/8th
		x = max.X / 8
		y = max.Y / 8
	case "medium": // 1/4th
		x = max.X / 4
		y = max.Y / 4
	case "large": // 1/2 size
		x = max.X / 2
		y = max.Y / 2
	}

	dstImage := imaging.Resize(img, x, y, imaging.Linear)
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, dstImage, nil); err != nil {
		return err
	}

	ctx.ResponseData.Header().Set("Content-Type", "image/jpeg")
	ctx.OK(buf.Bytes())
	return nil
}

// Upload runs the upload action.
func (c *ImageController) Upload(ctx *app.UploadImageContext) error {
	reader, err := ctx.MultipartReader()
	if err != nil {
		return goa.ErrBadRequest("failed to load multipart request: %s", err)
	}

	if reader == nil {
		return goa.ErrBadRequest("not a multipart request")
	}

	var buf bytes.Buffer
	for {
		buf.Reset()
		p, err := reader.NextPart()
		if err == io.EOF {
			break
		}

		if err != nil {
			return goa.ErrBadRequest("failed to load part: %s", err)
		}

		if _, err = buf.ReadFrom(p); err != nil {
			return errors.Wrap(err, "could not read from part")
		}

		if err = c.setImage(buf.Bytes()); err != nil {
			return errors.Wrap(err, "could not set image")
		}
	}

	ctx.ResponseData.Header().Set("Location", "/")
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
