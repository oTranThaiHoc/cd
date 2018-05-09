package handlers

import (
	"net/http"
	"encoding/json"
	"coconut.com/payload"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"fmt"
)

var PayloadsHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	projects, err := enumerateDirectory("./payloads")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(projects)

	var apps []payload.PayloadList
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
		appPayload := payload.PayloadList{
			Project:project.Name(),
			Apps:appPayloads,
		}

		apps = append(apps, appPayload)
	}

	p, _ := json.Marshal(apps)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(p))
})

func enumerateDirectory(dirPath string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	return files, err
}

func walkDir(root string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			dirs = append(dirs, path)
		}
		return nil
	})
	return dirs, err
}
