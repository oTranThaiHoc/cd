package db

import (
	"github.com/jackc/pgx"
	"log"
	"coconut.com/utils"
	"coconut.com/config"
	"coconut.com/payload"
)

var (
	pool *pgx.ConnPool
)

const (
	selectBuildSQL = `SELECT title, manifest_url FROM builds WHERE target=$1;`

	addBuildSQL = `INSERT INTO builds(title, target, manifest_url) VALUES($1, $2, $3) RETURNING id;`

	selectProjectSQL = `SELECT * FROM projects;`
)

func Setup(connPool *pgx.ConnPool) {
	pool = connPool
}

func PrepareStmt(conn *pgx.Conn) {
	utils.MustPrepare(conn, "selectBuildSQL", selectBuildSQL)
	utils.MustPrepare(conn, "addBuildSQL", addBuildSQL)
	utils.MustPrepare(conn, "selectProjectSQL", selectProjectSQL)
}

func InsertNewBuild(title string, target string, manifestUrl string) error {
	err := pool.QueryRow(addBuildSQL, title, target, manifestUrl).Scan(nil)
	if err != nil {
		log.Fatalf("inser new build error: %v - %v\n", title, err)
	}
	return err
}

func LoadBuildOptions() ([]config.BuildOption, error) {
	rows, err := pool.Query(selectProjectSQL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var projects []config.BuildOption
	for rows.Next() {
		var p config.BuildOption
		err = rows.Scan(&p.Id, &p.Project, &p.Targets, &p.Path)
		if err == nil {
			projects = append(projects, p)
		}
	}
	return projects, nil
}

func LoadPayloadList(target string) ([]payload.Payload, error) {
	rows, err := pool.Query(selectBuildSQL, target)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	var payloads []payload.Payload
	for rows.Next() {
		var p payload.Payload
		err = rows.Scan(&p.Title, &p.ManifestUrl)
		if err == nil {
			payloads = append(payloads, p)
		}
	}
	return payloads, nil
}
