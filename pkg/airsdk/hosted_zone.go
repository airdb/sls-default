package airsdk

import (
	"airdb.io/airdb/sailor"
	"github.com/airdb/scf-airdb/model/vo"
	"github.com/pkg/errors"
)

func ListHostedZone(input *vo.ListHostedZoneReq,
) (*vo.ListHostedZoneResp, error) {
	var output vo.ListHostedZoneResp

	endpoint := "https://scf.baobeihuijia.com" + "/test/airdb/v1/hosted_zone/list"

	var req sailor.HTTPClient

	req.SetURL(endpoint)
	req.SetBody(&input)

	if err := sailor.HTTPRequest(&req, &output); err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	return &output, nil
}
