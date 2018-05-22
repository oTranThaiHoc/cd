package renderer

import (
	"net/http"
	"github.com/thedevsaddam/renderer"
	"coconut.com/config"
)

var rnd *renderer.Render

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

var HomePageRenderHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	h := map[string]interface{}{
		"HttpEndPoint": config.HttpEndPoint,
		"WsEndPoint":   config.WsEndPoint,
	}
	rnd.HTML(w, http.StatusOK, "home", h)
})

var AboutPageRenderHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "about", nil)
})
