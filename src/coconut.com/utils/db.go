package utils

import (
	"github.com/jackc/pgx"
	"fmt"
)

// MustPrepare returns a prepared pg statement
func MustPrepare(db *pgx.Conn, name, query string) *pgx.PreparedStatement {
	stmt, err := db.Prepare(name, query)
	if err != nil {
		panic(fmt.Sprintf("query=%v, preparing failed: %v", query, err))
	}
	return stmt
}
