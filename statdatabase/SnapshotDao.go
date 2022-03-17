package statdatabase

import (
	"context"
	"database/sql"
	"time"

	"github.com/adgs85/gomonmarshalling/monmarshalling"
)

const insertSnapshot = `
INSERT INTO stats_snapshot
	(host_name, stat_type, instance_name, collected_ts, last_updated, polling_rate_ms, payload) 
VALUES 
	($1, $2, $3, $4, $5, $6, $7)
on CONFLICT (host_name, stat_type) DO UPDATE
	SET host_name=$1, stat_type=$2, instance_name=$3, collected_ts=$4, last_updated=$5, polling_rate_ms=$6, payload=$7

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

	_, err2 := stmt.ExecContext(ctx,
		meta.HostName,
		meta.StatType,
		meta.InstanceName,
		time.UnixMilli(meta.AgentTimestampUnixMs).UTC(),
		time.Now().UTC(),
		meta.PollRateMs,
		stat.Payload)

	CheckError(err2)
}
