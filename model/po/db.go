package po

import "airdb.io/airdb/sailor/dbutil"

const (
	dbMinaAPIRead  = "AIRDB_DB_MINA_API_READ"
	dbMinaAPIWrite = "AIRDB_DB_MINA_API_WRITE"
)

func InitDB() {
	dbNames := []string{}

	dbNames = append(dbNames,
		dbMinaAPIRead,
		// dbMinaAPIWirte,
	)

	dbutil.InitDB(dbNames)
}
