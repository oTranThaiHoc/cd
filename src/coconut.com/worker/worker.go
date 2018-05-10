package worker

import (
	"coconut.com/config"
	"os/exec"
	"log"
	"time"
)

type Job struct {
	Cfg config.BuildConfig
	Target string
	Title string
}

var (
	JobQueue = make(chan Job)
)

func init() {
	go func() {
		for {
			select {
			case job := <- JobQueue: {
				cmd := exec.Command("/bin/sh", "/Users/nguyen.van.hung/Workspaces/ci/scripts/remote_deploy.sh", job.Cfg.Path, job.Target, job.Title)
				_, err := cmd.Output()
				if err != nil {
					log.Fatal(err)
				}
			}
			default:
				break
			}

			time.Sleep(5 * time.Second)
		}
	}()
}
