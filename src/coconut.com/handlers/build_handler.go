package handlers

import (
	"net/http"
	"fmt"
	"coconut.com/config"
	"coconut.com/worker"
	"github.com/gorilla/mux"
	"encoding/json"
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

var BuildConfigHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)["key"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if key == "list" {
		// return all config
		cfg, err := json.Marshal(config.BuildConfigs)
		if err == nil {
			w.Write(cfg)
		}
		return
	}
	if key == "projects" {
		// return list project
		projects := make([]string, len(config.BuildConfigs))
		for i, c := range config.BuildConfigs {
			projects[i] = c.Project
		}
		l, err := json.Marshal(projects)
		if err == nil {
			w.Write(l)
		}
	} else {
		// return list target
		for _, c := range config.BuildConfigs {
			if c.Project == key {
				l, err := json.Marshal(c.Target)
				if err == nil {
					w.Write(l)
				}
				return
			}
		}
	}
})

