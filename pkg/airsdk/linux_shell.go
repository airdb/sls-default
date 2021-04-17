package airsdk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"airdb.io/airdb/sailor"
	"github.com/airdb/scf-airdb/model/vo"
	"github.com/pkg/errors"
)

func ListLinuxShell(input *vo.ListLinuxShellReq,
) (*vo.ListLinuxShellResp, error) {
	output := Response{}

	client := sailor.NewHTTPClient()

	client.SetDomain("scf.baobeihuijia.com")
	client.SetEndpoint("/test/airdb/v1/shell/list")
	client.SetBody(&input)
	client.SetUserAgent("scf-airdb/v0.0.1")
	client.SetDebug()

	if err := client.HTTPRequest(client, &output); err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "request failed")
	}

	res := vo.ListLinuxShellResp{}
	err := json.Unmarshal(output.Data, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal fail")
	}

	return &res, nil
}

func CreateLinuxShell(input *vo.CreateLinuxShellReq,
) (*vo.CreateLinuxShellResp, error) {
	output := Response{}

	client := sailor.NewHTTPClient()

	client.SetDomain("scf.baobeihuijia.com")
	client.SetEndpoint("/test/airdb/v1/shell/create")
	client.SetMethod(http.MethodPost)
	client.SetBody(&input)
	client.SetUserAgent("scf-airdb/v0.0.1")

	if err := client.HTTPRequest(client, &output); err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "request failed")
	}

	res := vo.CreateLinuxShellResp{}
	err := json.Unmarshal(output.Data, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal fail")
	}

	return &res, nil
}
