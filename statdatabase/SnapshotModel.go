package statdatabase

import "time"

type SnapshotModel struct {
	HostName      string    `db:"host_name"`
	StatType      string    `db:"stat_type"`
	InstanceName  string    `db:"instance_name"`
	CollectedTs   time.Time `db:"collected_ts"`
	LastUpdated   time.Time `db:"last_updated"`
	PollingRateMs int       `db:"polling_rate_ms"`
	Payload       string    `db:"payload"`
}
