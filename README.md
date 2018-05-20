# thumbnails
a Goa (golang) thumbnail generator API


This is just a learning app to learn Goa's DSL and actually build something with it.  
This app does this:

* Only keeps one image at a time in memory
* Provides an HTML page with all the thumbnails under `GET /`
* Update the image with `POST|PUT /image` with a JPEG image
* Show the image at `GET /image`
* Show exif metadata about the image at `GET /image/metadata`.  Returns a JSON blob of the metadata
* Generates thumbnails of the image under `GET /image/thumbnail/:type`.  Type can be:
  - small, create a 1/8 size version of the image
  - medium, creates a 1/4 size version of the image
  - large, creates a 1/2 size version of the image
  - square, creates a 125x125 image thumbnail
  - box(x/y), takes X/Y size params and attempts to fit image inside it



