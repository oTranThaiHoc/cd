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
		log.Printf("parse uploadfile failed: %v\n", err)
		return
	}
	defer file.Close()

	bundleId := r.PostFormValue("bundleid")
	targetName := r.PostFormValue("target")
	title := r.PostFormValue("title")
	note := r.PostFormValue("note")

	now := time.Now().Unix()
	d := fmt.Sprintf("./payloads/%v/%v/", targetName, now)
	utils.CreateDirIfNotExist(d)

	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile(d + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("write file failed: %v\n", err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// create app.plist
	manifest := fmt.Sprintf(config.ManifestFormat, fmt.Sprintf("%v/payloads/%v/%v/%v", config.HttpEndPoint, targetName, now, handler.Filename), bundleId, title, title)
	err = ioutil.WriteFile(d + "app.plist", []byte(manifest), 0666)
	if err != nil {
		log.Printf("write app.plist failed: %v\n", err)
		return
	}

	manifestUrlFormat := "itms-services://?action=download-manifest&url=%v/payloads/%v/%v/app.plist"
	manifestUrl := fmt.Sprintf(manifestUrlFormat, config.HttpEndPoint, targetName, now)

	// insert to db
	// title, manifestUrl
	err = db.InsertNewBuild(title, targetName, manifestUrl, d, note)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("new build added: %v\n", title)
	}
})

var UploadPublicHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		log.Printf("parse uploadfile failed: %v\n", err)
		return
	}
	defer file.Close()

	if len(handler.Filename) <= 0 {
		log.Println("empty file name")
		return
	}

	d := "./static/"
	utils.CreateDirIfNotExist(d)

	fmt.Fprintf(w, "%v", handler.Header)
	outFileName := d + handler.Filename
	_ = os.Remove(outFileName)
	f, err := os.OpenFile(outFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Printf("write file failed: %v\n", err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
})

