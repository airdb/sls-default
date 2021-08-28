package main

import (
	"time"

	"github.com/airdb/sls-default/internal/app/adapter"
	"github.com/airdb/sls-default/internal/version"
)

func main() {
	// Init the loc.
	loc, _ := time.LoadLocation("Asia/Shanghai")

	// Set timezone.
	version.CreatedAt = time.Now().In(loc)

	adapter.NewRouter()
}
