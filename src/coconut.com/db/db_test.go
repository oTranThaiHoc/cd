package db

import (
	"github.com/jackc/pgx"
	"coconut.com/config/pgconf"
	"log"
	"testing"
	"github.com/magiconair/properties/assert"
)

func newPgPool() (pg *pgx.ConnPool, err error) {
	cfg := pgconf.Config(nil)
	cfg.AfterConnect = func(conn *pgx.Conn) error {
		PrepareStmt(conn)
		return nil
	}
	pg, err = pgx.NewConnPool(cfg)
	if err != nil {
		return nil, err
	}
	return
}

func init() {
	c, err := newPgPool()
	if err != nil {
		log.Fatal(err)
		return
	}
	Setup(c)
}

func TestInsertNewBuild(t *testing.T) {
	title := "TEST.Example.1234"
	target := "NARUTO"
	manifestUrl := "itms-services://?action=download-manifest&url=https://ipa.haipq.com/payloads/NARUTO/1525918397/app.plist"
	path := ""

	err := InsertNewBuild(title, target, manifestUrl, path)
	assert.Equal(t, err, nil)
}