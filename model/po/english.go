package po

import (
	"airdb.io/airdb/sailor/dbutil"
	"gorm.io/gorm"
)

type TabDailyEnglish struct {
	gorm.Model
	Category string
	Tags     string
	En       string
	Cn       string
}

func ListDailyEnglish(pageNo, pageSize int) []*TabDailyEnglish {
	var shells []*TabDailyEnglish

	var offset int = 0
	if pageSize > 0 {
		offset = (pageNo - 1) * pageSize
	}

	dbutil.ReadDB(dbMasterRead).Debug().Offset(offset).Limit(pageSize).Find(&shells)

	return shells
}

func CreateDailyEnglish(en *TabDailyEnglish) {
	dbutil.WriteDB(dbMasterWrite).Create(&en)
}
