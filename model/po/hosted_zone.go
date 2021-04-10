package po

import (
	"time"

	"airdb.io/airdb/sailor/dbutil"
	"gorm.io/gorm"
)

type TabHostedZone struct {
	gorm.Model
	HostedZone   string
	Registry     string
	ExpiredAt    time.Time `gorm:"type:datetime"`
	RegisteredAt time.Time `gorm:"type:datetime"`
}

func ListHostedZone(pageNo, pageSize int) []*TabHostedZone {
	var HostedZones []*TabHostedZone

	var offset int = 0
	if pageSize > 0 {
		offset = (pageNo - 1) * pageSize
	}

	dbutil.ReadDB(dbMasterRead).Debug().Offset(offset).Limit(pageSize).Find(&HostedZones)

	return HostedZones
}

func CreateHostedZone(zone *TabHostedZone) {
	dbutil.WriteDB(dbMasterWrite).Create(&zone)
}
