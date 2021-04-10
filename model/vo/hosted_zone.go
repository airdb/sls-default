package vo

import (
	"fmt"
	"time"

	"airdb.io/airdb/sailor/timeutil"
	"github.com/airdb/scf-airdb/model/po"
	"gorm.io/gorm"
)

type HostedZone struct {
	ID uint `json:"id"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`
	HostedZone   string    `json:"hosted_zone"`
	Registry     string    `json:"registry"`
	Status       string    `json:"status"`
	ExpiredAt    time.Time `json:"expired_at"`
	RegisteredAt time.Time `json:"registered_at"`
}

func FromPoHostedZone(zone *po.TabHostedZone) *HostedZone {
	return &HostedZone{
		ID:           zone.ID,
		HostedZone:   zone.HostedZone,
		Registry:     zone.Registry,
		ExpiredAt:    zone.ExpiredAt,
		RegisteredAt: zone.RegisteredAt,
	}
}

func ToPoHostedZone(zone *HostedZone) *po.TabHostedZone {
	return &po.TabHostedZone{
		Model:        gorm.Model{},
		HostedZone:   zone.HostedZone,
		Registry:     zone.Registry,
		ExpiredAt:    zone.ExpiredAt,
		RegisteredAt: zone.RegisteredAt,
	}
}

type CreateHostedZoneReq struct {
	HostedZone    string  `json:"hosted_zone"`
	Registry      string  `json:"registry"`
	ExpiredDay    *string `json:"expired_day"`
	RegisteredDay *string `json:"registered_day"`
	// RegisteredDay string  `json:"registered_day,omitempty"`
	// ExpiredAt  time.Time `json:"expired_at"`
}

type CreateHostedZoneResp struct {
}

type ListHostedZoneReq struct {
	PageNo   int `form:"page_no"`
	PageSize int `form:"page_size"`
}

type ListHostedZoneResp struct {
	HostedZones []*HostedZone `json:"hosted_zones"`
}

func ListHostedZone(req *ListHostedZoneReq) *ListHostedZoneResp {
	var resp ListHostedZoneResp

	zones := po.ListHostedZone(req.PageNo, req.PageSize)

	expiringHours := 1440 // 60 * 24
	for _, z := range zones {
		_zone := FromPoHostedZone(z)
		switch {
		case time.Now().After(_zone.ExpiredAt):
			_zone.Status = "expired"
		case time.Now().Add(time.Duration(expiringHours) * time.Hour).After(_zone.ExpiredAt):
			_zone.Status = "expiring"
		default:
			_zone.Status = "active"
		}

		resp.HostedZones = append(resp.HostedZones, _zone)
	}

	return &resp
}

func CreateHostedZone(req *CreateHostedZoneReq) *CreateHostedZoneResp {
	zone := &HostedZone{
		HostedZone: req.HostedZone,
		Registry:   req.Registry,
		// ExpiredAt:  req.ExpiredAt,
		ExpiredAt: time.Now(),
	}

	fmt.Println("zone is ", zone)
	if req.ExpiredDay != nil {
		t, err := time.Parse(timeutil.TimeFormatDay, *req.ExpiredDay)
		if err != nil {
			return nil
		}

		zone.ExpiredAt = t
	}

	if req.RegisteredDay != nil {
		t, err := time.Parse(timeutil.TimeFormatDay, *req.RegisteredDay)
		if err != nil {
			return nil
		}

		zone.RegisteredAt = t
	}

	po.CreateHostedZone(ToPoHostedZone(zone))
	return &CreateHostedZoneResp{}
}
