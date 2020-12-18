package po

import "airdb.io/airdb/sailor/dbutil"

const (
	dbMinaAPIRead  = "AIRDB_DB_MINA_API_READ"
	dbMinaAPIWirte = "AIRDB_DB_MINA_API_WRITE"
)

func InitDB() {
	dbNames := []string{}

	dbNames = append(dbNames,
		// dbBbhjBBSRead,
		dbMinaAPIRead,
		// dbMinaAPIWirte,
	)

	dbutil.InitDB(dbNames)
}
