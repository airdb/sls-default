package airsdk

import (
	"encoding/json"
	"fmt"

	"airdb.io/airdb/sailor"
	"github.com/airdb/scf-airdb/model/vo"
	"github.com/pkg/errors"
)

type Response struct {
	Success bool            `json:"success"`
	Error   interface{}     `json:"error"`
	Data    json.RawMessage `json:"data"`
}

func ListHostedZone(input *vo.ListHostedZoneReq,
) (*vo.ListHostedZoneResp, error) {
	output := Response{}

	client := sailor.NewHTTPClient()

	client.SetDomain("scf.baobeihuijia.com")
	client.SetEndpoint("/test/airdb/v1/hosted_zone/list")
	client.SetBody(&input)
	client.SetUserAgent("scf-airdb/v0.0.1")

	if err := client.HTTPRequest(client, &output); err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "request failed")
	}

	res := vo.ListHostedZoneResp{}
	err := json.Unmarshal(output.Data, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal fail")
	}

	return &res, nil
}
