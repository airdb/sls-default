package vo

import (
	"fmt"
	"time"

	"github.com/airdb/scf-airdb/model/po"
	"gorm.io/gorm"
)

type HostedZone struct {
	ID uint `json:"id"`
	// CreatedAt time.Time  `json:"created_at"`
	// UpdatedAt time.Time  `json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty"`
	HostedZone string    `json:"hosted_zone"`
	Registry   string    `json:"registry"`
	ExpiredAt  time.Time `json:"expired_at"`
}

func FromPoHostedZone(zone *po.TabHostedZone) *HostedZone {
	return &HostedZone{
		ID:         zone.ID,
		HostedZone: zone.HostedZone,
		Registry:   zone.Registry,
		ExpiredAt:  zone.ExpiredAt,
	}
}

func ToPoHostedZone(zone *HostedZone) *po.TabHostedZone {
	return &po.TabHostedZone{
		Model:      gorm.Model{},
		HostedZone: zone.HostedZone,
		Registry:   zone.Registry,
		ExpiredAt:  zone.ExpiredAt,
	}
}

type CreateHostedZoneReq struct {
	HostedZone string    `json:"hosted_zone"`
	Registry   string    `json:"registry"`
	ExpiredAt  time.Time `json:"expired_at"`
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

	for _, z := range zones {
		resp.HostedZones = append(resp.HostedZones, FromPoHostedZone(z))
	}

	return &resp
}

func CreateHostedZone(req *CreateHostedZoneReq) *CreateHostedZoneResp {
	zone := &HostedZone{
		HostedZone: req.HostedZone,
		Registry:   req.Registry,
		ExpiredAt:  req.ExpiredAt,
	}
	fmt.Println("zone is ", zone)

	po.CreateHostedZone(ToPoHostedZone(zone))
	return &CreateHostedZoneResp{}
}
