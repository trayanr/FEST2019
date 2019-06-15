package controllers

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

var R *render.Engine
var assetsBox = packr.NewBox("../templates")

func init() {
	R = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "layouts/main.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{},
	})

}
