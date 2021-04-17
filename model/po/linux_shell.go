package po

import (
	"airdb.io/airdb/sailor/dbutil"
	"gorm.io/gorm"
)

type TabLinuxShell struct {
	gorm.Model
	Command string
	Shell   string
	Tags    string
	Comment string
	Ref     string
}

func ListLinuxShell(pageNo, pageSize int) []*TabLinuxShell {
	var shells []*TabLinuxShell

	var offset int = 0
	if pageSize > 0 {
		offset = (pageNo - 1) * pageSize
	}

	dbutil.ReadDB(dbMasterRead).Debug().Offset(offset).Limit(pageSize).Find(&shells)

	return shells
}

func CreateLinuxShell(shell *TabLinuxShell) {
	dbutil.WriteDB(dbMasterWrite).Create(&shell)
}
