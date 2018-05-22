package handlers

import (
	"net/http"
	"fmt"
	"coconut.com/config"
	"github.com/gorilla/mux"
	"encoding/json"
	"log"
	"coconut.com/payload"
	"io/ioutil"
	"os"
	"coconut.com/db"
	"coconut.com/worker"
)

var PayloadsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var payloads []payload.List
	for _, buildOption := range config.BuildOptions {
		for _, target := range buildOption.Targets {
			ps, err := db.LoadPayloadList(target.Name)
			if err == nil {
				p := payload.List{
					Target:target.Name,
					Payloads:ps,
				}
				payloads = append(payloads, p)
			}
		}
	}
	payloadJson, _ := json.Marshal(payloads)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(payloadJson)
})

var PayloadsHandlers = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	projects, err := enumerateDirectory("./payloads")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(projects)

	var apps []payload.List
	for _, project := range projects {
		if !project.IsDir() {
			continue
		}
		builds, err := enumerateDirectory(fmt.Sprintf("./payloads/%v", project.Name()))
		if err != nil {
			continue
		}
		var appPayloads []payload.Payload
		for _, build := range builds {
			if !build.IsDir() {
				continue
			}
			appJsonPath := fmt.Sprintf("./payloads/%v/%v/app.json", project.Name(), build.Name())
			raw, err := ioutil.ReadFile(appJsonPath)
			if err != nil {
				continue
			}
			var app payload.Payload
			err = json.Unmarshal(raw, &app)
			if err == nil {
				appPayloads = append(appPayloads, app)
			}
		}
		appPayload := payload.List {
			Target:project.Name(),
			Payloads:appPayloads,
		}

		apps = append(apps, appPayload)
	}

	p, _ := json.Marshal(apps)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write(p)
})

func enumerateDirectory(dirPath string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	return files, err
}

var BuildHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("build command: %v\n", r)
	project := r.PostFormValue("project")
	target := r.PostFormValue("target")
	title := r.PostFormValue("title")

	bundleId := ""
	for _, p := range config.BuildOptions {
		if p.Project == project {
			for _, t := range p.Targets {
				if t.Name == target {
					bundleId = t.BundleId
					break
				}
			}
		}
		if len(bundleId) > 0 {
			break
		}
	}

	for _, cfg := range config.BuildOptions {
		if cfg.Project == project {
			j := worker.Job{
				Cfg: cfg,
				Target: target,
				Title: title,
				ClientIp: r.RemoteAddr,
				BundleId: bundleId,
			}
			worker.JobQueue <- j
			fmt.Printf("Run build on project %v, target %v, title %v, bundleid: %v, clientIp: %v\n", project, target, title, bundleId, r.RemoteAddr)
			break
		}
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
})

var BuildConfigHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	key, ok := mux.Vars(r)["key"]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if key == "list" {
		// return all config
		cfg, err := json.Marshal(config.BuildOptions)
		if err == nil {
			w.Write(cfg)
		}
		return
	}
	if key == "projects" {
		// return list project
		projects := make([]string, len(config.BuildOptions))
		for i, c := range config.BuildOptions {
			projects[i] = c.Project
		}
		l, err := json.Marshal(projects)
		if err == nil {
			w.Write(l)
		}
	} else {
		// return list target
		for _, c := range config.BuildOptions {
			if c.Project == key {
				l, err := json.Marshal(c.Targets)
				if err == nil {
					w.Write(l)
				}
				return
			}
		}
	}
})

var RemoveBuildHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	manifestUrl := r.PostFormValue("manifest_url")
	path, err := db.FindBuild(manifestUrl)
	if err != nil {
		log.Printf("can not find build with manifest: %v\n", manifestUrl)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = db.RemoveBuild(manifestUrl)
	if err == nil {
		// remove physical directory
		os.RemoveAll(path)
	}
})
