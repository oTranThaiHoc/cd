package handlers

import (
	"net/http"
	"fmt"
	"coconut.com/config"
	"coconut.com/worker"
)

var BuildHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	project := r.PostFormValue("project")
	target := r.PostFormValue("target")
	title := r.PostFormValue("title")

	for _, cfg := range config.BuildConfigs {
		if cfg.Project == project {
			fmt.Printf("Run build on project %v, target %v, title %v\n", project, target, title)
			j := worker.Job{
				Cfg: cfg,
				Target: target,
				Title: title,
			}
			worker.JobQueue <- j
			break
		}
	}
	w.WriteHeader(http.StatusOK)
})

