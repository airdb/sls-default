package po

import "github.com/airdb/sailor/dbutils"

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

	dbutils.InitDB(dbNames)
}
