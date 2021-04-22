package vo

import (
	"github.com/airdb/scf-airdb/model/po"
	"gorm.io/gorm"
)

type DailyEnglish struct {
	ID uint `json:"id"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`
	Category string
	Tags     string
	En       string
	Cn       string
}

func FromPoDailyEnglish(en *po.TabDailyEnglish) *DailyEnglish {
	return &DailyEnglish{
		ID:       en.ID,
		Category: en.Category,
		Tags:     en.Tags,
		En:       en.En,
		Cn:       en.Cn,
	}
}

func ToPoDailyEnglish(en *DailyEnglish) *po.TabDailyEnglish {
	return &po.TabDailyEnglish{
		Model:    gorm.Model{},
		Category: en.Category,
		Tags:     en.Tags,
		En:       en.En,
		Cn:       en.Cn,
	}
}

type CreateDailyEnglishReq struct {
	Category string `json:"category" binding:"required"`
	Tags     string `json:"tags"`
	En       string `json:"en" binding:"required"`
	Cn       string `json:"cn"`
}

type CreateDailyEnglishResp struct {
}

type ListDailyEnglishReq struct {
	PageNo   int `form:"page_no"`
	PageSize int `form:"page_size"`
}

type ListDailyEnglishResp struct {
	En []*DailyEnglish `json:"en"`
}

func ListDailyEnglish(req *ListDailyEnglishReq) *ListDailyEnglishResp {
	var resp ListDailyEnglishResp

	ens := po.ListDailyEnglish(req.PageNo, req.PageSize)

	for _, en := range ens {
		_en := FromPoDailyEnglish(en)

		resp.En = append(resp.En, _en)
	}

	return &resp
}

func CreateDailyEnglish(req *CreateDailyEnglishReq) *CreateDailyEnglishResp {
	en := &DailyEnglish{
		Tags:     req.Tags,
		En:       req.En,
		Cn:       req.Cn,
		Category: req.Category,
	}

	po.CreateDailyEnglish(ToPoDailyEnglish(en))
	return &CreateDailyEnglishResp{}
}
