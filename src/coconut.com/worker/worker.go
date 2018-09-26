package worker

import (
	"coconut.com/config"
	"os/exec"
	"log"
	"time"
	"fmt"
	"coconut.com/utils"
	"io/ioutil"
	"coconut.com/db"
	"os"
	"io"
	"bufio"
	"strings"
)

type Job struct {
	Cfg config.BuildOption
	Target string
	Title string
	ClientIp string
	BundleId string
	Note string
	Pull string
	Version string
}

var (
	JobQueue = make(chan Job)
	JobDone = make(chan Job)
)

func init() {
	go func() {
		for {
			select {
			case job := <- JobQueue: {
				buildScript := fmt.Sprintf("%v/local_deploy.sh", config.ScriptPath)
				cmd := exec.Command("/bin/sh", buildScript, job.Cfg.Path, job.Target, job.Title, job.Pull, job.Version)
				_, err := cmd.Output()
				if err != nil {
					log.Printf("Build target %v with title %v failed: %v\n", job.Target, job.Title, err)
				} else {
					target := strings.Replace(job.Target, " ", "", -1)
					buildDirPath := fmt.Sprintf("%v/build_%v", config.ScriptPath, target)
					contents, err := enumerateDirectory(buildDirPath)
					if err != nil {
						break
					}
					for _, content := range contents {
						if strings.HasSuffix(content.Name(), "ipa") {
							generateBuild(job, content.Name())
							break
						}
					}
				}
			}
			default:
				break
			}

			time.Sleep(5 * time.Second)
		}
	}()
}

func enumerateDirectory(dirPath string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	return files, err
}

func generateBuild(job Job, buildFileName string) {
	target := strings.Replace(job.Target, " ", "", -1)

	now := time.Now().Unix()
	d := fmt.Sprintf("./payloads/%v/%v/", target, now)
	utils.CreateDirIfNotExist(d)

	// get build
	buildFilePath := fmt.Sprintf("%v/build_%v/%v", config.ScriptPath, target, buildFileName)
	fmt.Printf("build file path: %v\n", buildFilePath)

	// Read all content of src to data
	r, err := os.Open(buildFilePath)
	if err != nil {
		log.Println(err)
		return
	}
	// Write data to dst
	f, err := os.OpenFile(d + buildFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	file := bufio.NewReader(r)
	io.Copy(f, file)

	// create app.plist
	manifest := fmt.Sprintf(config.ManifestFormat, fmt.Sprintf("%v/payloads/%v/%v/%v", config.HttpEndPoint, target, now, buildFileName), job.BundleId, job.Title, job.Title)
	err = ioutil.WriteFile(d + "app.plist", []byte(manifest), 0666)
	if err != nil {
		log.Println(err)
		return
	}

	manifestUrlFormat := "itms-services://?action=download-manifest&url=%v/payloads/%v/%v/app.plist"
	manifestUrl := fmt.Sprintf(manifestUrlFormat, config.HttpEndPoint, target, now)

	// insert to db
	// title, manifestUrl
	err = db.InsertNewBuild(job.Title, job.Target, manifestUrl, d, job.Note)
	if err != nil {
		log.Println(err)
	} else {
		log.Printf("new local build added: %v\n", job.Title)
		JobDone <- job
	}
}