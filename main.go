package main

import (
	"fmt"
	"os"
	"time"

	"github.com/airdb/scf-noah/api"
	"github.com/airdb/scf-noah/internal/version"
	"github.com/airdb/scf-noah/model/po"
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
