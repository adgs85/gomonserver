package statdatabase

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

var dataSource = func() *sql.DB {

	db, err := sql.Open("postgres", psqlConnectionStr)
	CheckError(err)

	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return db
}()

func GetConnWithContext() (*sql.Conn, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	con, err := dataSource.Conn(ctx)
	CheckError(err)
	return con, ctx, cancel
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
