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
	selectBuildSQL = `SELECT title, manifest_url, note FROM builds WHERE target=$1 ORDER BY title;`

	addBuildSQL = `INSERT INTO builds(title, target, manifest_url, path, note) VALUES($1, $2, $3, $4, $5) RETURNING id;`

	selectProjectSQL = `SELECT * FROM projects;`

	removeBuildSQL = `DELETE FROM builds WHERE manifest_url=$1;`

	findBuildPathSQL = `SELECT path FROM builds WHERE manifest_url=$1;`

	findBuildReleaseNoteSQL = `SELECT note FROM builds WHERE manifest_url=$1;`
)

func Setup(connPool *pgx.ConnPool) {
	pool = connPool
}

func PrepareStmt(conn *pgx.Conn) {
	utils.MustPrepare(conn, "selectBuildSQL", selectBuildSQL)
	utils.MustPrepare(conn, "addBuildSQL", addBuildSQL)
	utils.MustPrepare(conn, "selectProjectSQL", selectProjectSQL)
	utils.MustPrepare(conn, "removeBuildSQL", removeBuildSQL)
	utils.MustPrepare(conn, "findBuildPathSQL", findBuildPathSQL)
	utils.MustPrepare(conn, "findBuildReleaseNoteSQL", findBuildReleaseNoteSQL)
}

func InsertNewBuild(title string, target string, manifestUrl string, path string, note string) error {
	err := pool.QueryRow(addBuildSQL, title, target, manifestUrl, path, note).Scan(nil)
	if err != nil {
		log.Fatalf("inser new build error: %v - %v\n", title, err)
	}
	return err
}

func FindBuild(manifestUrl string) (string, error) {
	var path string
	err := pool.QueryRow(findBuildPathSQL, manifestUrl).Scan(&path)
	return path, err
}

func RemoveBuild(manifestUrl string) error {
	_, err := pool.Exec(removeBuildSQL, manifestUrl)
	return err
}

func GetBuildReleaseNote(manifestUrl string) (string, error) {
	var note string
	err := pool.QueryRow(findBuildReleaseNoteSQL, manifestUrl).Scan(&note)
	return note, err
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
		err = rows.Scan(&p.Title, &p.ManifestUrl, &p.Note)
		if err == nil {
			payloads = append(payloads, p)
		}
	}
	return payloads, nil
}
