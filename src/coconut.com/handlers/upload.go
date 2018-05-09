package handlers

import (
	"net/http"
	"fmt"
	"os"
	"io"
	"time"
	"coconut.com/utils"
	"io/ioutil"
	"coconut.com/payload"
)

var UploadHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bundleId := r.PostFormValue("bundleid")
	projectName := r.PostFormValue("project")
	title := r.PostFormValue("title")

	now := time.Now().Unix()
	d := fmt.Sprintf("./payloads/%v/%v/", projectName, now)
	utils.CreateDirIfNotExist(d)

	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile(d + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// create app.plist
	plistFormat := `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>items</key>
	<array>
		<dict>
			<key>assets</key>
			<array>
				<dict>
					<key>kind</key>
					<string>software-package</string>
					<key>url</key>
					<string>%v</string>
				</dict>
			</array>
			<key>metadata</key>
			<dict>
				<key>bundle-identifier</key>
				<string>%v</string>
				<key>bundle-version</key>
				<string>1.0</string>
				<key>kind</key>
				<string>software</string>
				<key>subtitle</key>
				<string>%v</string>
				<key>title</key>
				<string>%v</string>
			</dict>
		</dict>
	</array>
</dict>
</plist>
`
	manifest := fmt.Sprintf(plistFormat, fmt.Sprintf("%v/payloads/%v/%v/%v", utils.Domain, projectName, now, handler.Filename), bundleId, title, title)
	err = ioutil.WriteFile(d + "app.plist", []byte(manifest), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	urlFormat := "itms-services://?action=download-manifest&url=%v/payloads/%v/%v/app.plist"
	url := fmt.Sprintf(urlFormat, utils.Domain, projectName, now)
	jsonFormat := `{
	"title":"%v",
	"url":"%v"
}`
	json := fmt.Sprintf(jsonFormat, title, url)
	err = ioutil.WriteFile(d + "app.json", []byte(json), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := payload.Payload {
		Title: title,
		Url: url,
	}

	payload.Payloads[0] = p
})