package vo

import (
	"fmt"
	"time"

	"github.com/airdb/scf-mina/model/po"
)

type Lost struct {
	ID uint `json:"id"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`

	UUID      string `json:"uuid"`
	AvatarURL string `json:"avatar_url"`
	Nickname  string `json:"nickname"`
	// 0: unknown,  1: male   2: female
	Gender          uint      `json:"gender"`
	Title           string    `json:"title"`
	Subject         string    `json:"subject"`
	Characters      string    `json:"characters"`
	Details         string    `json:"details"`
	DataFrom        string    `json:"data_from"`
	BirthedProvince string    `json:"birthed_province"`
	BirthedCity     string    `json:"birthed_city"`
	BirthedCountry  string    `json:"birthed_country"`
	BirthedAddress  string    `json:"birthed_address"`
	BirthedAt       time.Time `json:"birthed_at"`

	MissedCountry  string    `json:"missed_country"`
	MissedProvince string    `json:"missed_province"`
	MissedCity     string    `json:"missed_city"`
	MissedAddress  string    `json:"missed_address"`
	MissedAt       time.Time `json:"missed_at"`
	Handler        string    `json:"handler"`
	Babyid         string    `json:"babyid"`
	Category       string    `json:"category"`
	Height         string    `json:"Height"`
	SyncStatus     int       `json:"sync_status"`
	Images         []string  `json:"images"`
	ViewNum        uint64    `json:"view_num"`
	CommentNum     uint64    `json:"comment_num"`
	ShareNum       uint64    `json:"share_num"`
}

type Attachment struct {
	ImageURL string `json:"image_url"`
}

func FromPoLost(lost *po.Lost) *Lost {
	return &Lost{
		ID:              lost.ID,
		UUID:            lost.UUID,
		AvatarURL:       lost.AvatarURL,
		Nickname:        lost.Nickname,
		Gender:          lost.Gender,
		Title:           lost.Title,
		Subject:         lost.Subject,
		Characters:      lost.Characters,
		Details:         lost.Details,
		DataFrom:        lost.DataFrom,
		BirthedProvince: lost.BirthedProvince,
		BirthedCity:     lost.BirthedCity,
		BirthedCountry:  lost.BirthedCountry,
		BirthedAddress:  lost.BirthedAddress,
		BirthedAt:       lost.BirthedAt,
		MissedCountry:   lost.MissedCountry,
		MissedProvince:  lost.MissedProvince,
		MissedCity:      lost.MissedCity,
		MissedAddress:   lost.MissedAddress,
		MissedAt:        lost.MissedAt,
		Handler:         lost.Handler,
		Babyid:          lost.Babyid,
		Category:        lost.Category,
		Height:          lost.Height,
		SyncStatus:      lost.SyncStatus,
		ViewNum:         1000, //nolint
		CommentNum:      101,  //nolint
		ShareNum:        100,  //nolint
	}
}

type ListLostReq struct {
	PageNo   int `form:"pageNo"`
	PageSize int `form:"pageSize"`
}

type ListLostResp struct {
	Articles []*Lost `json:"articles"`
}

type QueryArticleReq struct {
	ID uint `form:"id"`
}

type QueryArticleResp struct {
	Article *Lost `json:"article"`
}

func FromPoLosts(losts []*po.Lost) []*Lost {
	_losts := make([]*Lost, 0, len(losts))

	for _, lost := range losts {
		_losts = append(_losts, FromPoLost(lost))
	}

	return _losts
}

func SearchLost(keywords string) []*Lost {
	var losts []*Lost

	aa := po.SearchLost(keywords)
	for _, a := range aa {
		fmt.Println("po data", a)
	}

	losts = append(losts, FromPoLosts(po.SearchLost(keywords))...)

	return losts
}

func ListLost(req *ListLostReq) *ListLostResp {
	var resp ListLostResp

	losts := po.ListLost(req.PageNo, req.PageSize)

	for _, lost := range losts {
		article := FromPoLost(lost)

		article.Images = append(article.Images,
			[]string{
				"https://attachment.baobeihuijia.com/forum/201305/20/1257299vn00r3yhanr4s3g.jpg",
				"https://bbs.baobeihuijia.com/data/attachment/forum/202006/02/164751myyf1ss1jxzfd9is.jpg",
			}...,
		)

		resp.Articles = append(resp.Articles, article)
	}

	return &resp
}

func QueryArticle(req *QueryArticleReq) *QueryArticleResp {
	var resp QueryArticleResp

	lost := po.QueryLostByID(req.ID)

	resp.Article = FromPoLost(lost)

	resp.Article.Images = append(resp.Article.Images,
		[]string{
			"https://bbs.baobeihuijia.com/data/attachment/forum/202006/02/164751myyf1ss1jxzfd9is.jpg",
			"https://bbs.baobeihuijia.com/data/attachment/forum/202006/02/164751myyf1ss1jxzfd9is.jpg",
			"https://bbs.baobeihuijia.com/data/attachment/forum/202006/02/164751myyf1ss1jxzfd9is.jpg",
			"https://bbs.baobeihuijia.com/data/attachment/forum/202006/02/164751myyf1ss1jxzfd9is.jpg",
		}...,
	)

	return &resp
}
