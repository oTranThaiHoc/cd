package renderer

import (
	"net/http"
	"github.com/thedevsaddam/renderer"
	"coconut.com/config"
	"os"
	"fmt"
	"container/list"
)

var rnd *renderer.Render
const (
	root_folder = "/Users/hung/workspaces/cd/static"
)

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./tpl/*.html",
	}

	rnd = renderer.New(opts)
}

// Manages directory listings
type dirlisting struct {
	Server string
	Children_files []string
}

func copyToArray(src *list.List) []string {
	dst := make([]string, src.Len())

	i := 0
	for e := src.Front(); e != nil; e = e.Next() {
		dst[i] = e.Value.(string)
		i = i + 1
	}

	return dst
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

var FilesPageRenderHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "files", nil)
})

var FilesListingHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open(root_folder)
	if err != nil {
		http.Error(w, "404 Not Found : Error while opening the file.", 404)
		return
	}

	defer f.Close()

	// Checking if the opened handle is really a file
	statinfo, err := f.Stat()
	if err != nil {
		http.Error(w, "500 Internal Error : stat() failure.", 500)
		return
	}

	if !statinfo.IsDir() {
		return
	}

	names, _ := f.Readdir(-1)

	children_files_tmp := list.New()

	for _, val := range names {
		if val.Name()[0] == '.' {
			continue
		} // Remove hidden files from listing

		if !val.IsDir() {
			children_files_tmp.PushBack(config.HttpEndPoint + "/static/" + val.Name())
		}
	}

	// And transfer the content to the final array structure
	children_files := copyToArray(children_files_tmp)
	fmt.Printf("files: %v\n", children_files)

	data := dirlisting{
		Server: config.HttpEndPoint,
		Children_files: children_files,
	}

	rnd.HTML(w, http.StatusOK, "files", data)
})
