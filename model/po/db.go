package po

import "airdb.io/airdb/sailor/dbutil"

const (
	dbMasterRead  = "AIRDB_MASTER_DB_READ"
	dbMasterWrite = "AIRDB_MASTER_DB_WRITE"
)

func InitDB() {
	dbNames := []string{}

	dbNames = append(dbNames,
		dbMasterWrite,
		dbMasterRead,
	)

	dbutil.InitDB(dbNames)
}
