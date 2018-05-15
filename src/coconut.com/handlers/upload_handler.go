package handlers

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"time"
	"coconut.com/utils"
	"io/ioutil"
	"coconut.com/config"
	"coconut.com/db"
	"log"
)

var UploadHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	bundleId := r.PostFormValue("bundleid")
	targetName := r.PostFormValue("target")
	title := r.PostFormValue("title")

	now := time.Now().Unix()
	d := fmt.Sprintf("./payloads/%v/%v/", targetName, now)
	utils.CreateDirIfNotExist(d)

	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile(d + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// create app.plist
	manifest := fmt.Sprintf(config.ManifestFormat, fmt.Sprintf("%v/payloads/%v/%v/%v", config.HttpEndPoint, targetName, now, handler.Filename), bundleId, title, title)
	err = ioutil.WriteFile(d + "app.plist", []byte(manifest), 0666)
	if err != nil {
		log.Fatal(err)
		return
	}

	manifestUrlFormat := "itms-services://?action=download-manifest&url=%v/payloads/%v/%v/app.plist"
	manifestUrl := fmt.Sprintf(manifestUrlFormat, config.HttpEndPoint, targetName, now)

	// insert to db
	// title, manifestUrl
	err = db.InsertNewBuild(title, targetName, manifestUrl)
	if err != nil {
		log.Fatal(err)
	}
})
