package main

import (
	"fmt"
	"os"
	"time"

	"github.com/airdb/scf-airdb/api"
	"github.com/airdb/scf-airdb/internal/version"
	"github.com/airdb/scf-airdb/model/po"
)

func main() {
	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	version.CreatedAt = time.Now().In(loc)
	fmt.Println(os.Getenv("AIRDB_DB_MINA_API_READ"))
	po.InitDB()

	api.NewRouter()
}
