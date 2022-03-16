package statdatabase

import (
	"context"
	"database/sql"
	"time"

	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

const insertSnapshot = `INSERT INTO stats_snapshot
(host_name, stat_type, collected_ts, last_updated, polling_rate_ms, payload) VALUES ($1, $2, $3, $4, $5, $6);
`

func InsertStat(stat *monmarshalling.Stat) {
	conn, ctx, cancel := GetConnWithContext()
	defer conn.Close()
	defer cancel()
	insertStatSnapshot(ctx, *conn, stat)
}

func insertStatSnapshot(ctx context.Context, c sql.Conn, stat *monmarshalling.Stat) {

	stmt, err := c.PrepareContext(ctx, insertSnapshot)
	CheckError(err)
	meta := stat.MetaData

	r, err := stmt.ExecContext(ctx,
		meta.HostName,
		meta.StatType,
		time.UnixMilli(meta.AgentTimestampUnixMs).UTC(),
		time.Now().UTC(),
		meta.PollRateMs,
		stat.Payload)

	r.RowsAffected()
	CheckError(err)
}
