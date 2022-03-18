package statdatabase

import (
	"context"
	"database/sql"
	"time"

	"github.com/adgs85/gomonserver/monserver"
	_ "github.com/lib/pq"
)

var dataSource = func() *sql.DB {

	db, err := sql.Open("postgres", psqlConnectionStr)
	monserver.PanicOnError(err)

	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return db
}()

func GetConnWithContext() (*sql.Conn, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	con, err := dataSource.Conn(ctx)
	monserver.PanicOnError(err)
	return con, ctx, cancel
}
