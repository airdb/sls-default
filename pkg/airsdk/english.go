package airsdk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"airdb.io/airdb/sailor"
	"github.com/airdb/scf-airdb/model/vo"
	"github.com/pkg/errors"
)

func ListDailyEnglish(input *vo.ListDailyEnglishReq,
) (*vo.ListDailyEnglishResp, error) {
	output := Response{}

	client := sailor.NewHTTPClient()

	client.SetDomain("scf.baobeihuijia.com")
	client.SetEndpoint("/test/airdb/v1/english/list")
	client.SetBody(&input)
	client.SetUserAgent("scf-airdb/v0.0.1")
	// client.SetDebug()

	if err := client.HTTPRequest(client, &output); err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "request failed")
	}

	res := vo.ListDailyEnglishResp{}
	err := json.Unmarshal(output.Data, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal fail")
	}

	return &res, nil
}

func CreateDailyEnglish(input *vo.CreateDailyEnglishReq,
) (*vo.CreateDailyEnglishResp, error) {
	output := Response{}

	client := sailor.NewHTTPClient()

	client.SetDomain("scf.baobeihuijia.com")
	client.SetEndpoint("/test/airdb/v1/english/create")
	client.SetMethod(http.MethodPost)
	client.SetBody(&input)
	client.SetUserAgent("scf-airdb/v0.0.1")

	if err := client.HTTPRequest(client, &output); err != nil {
		fmt.Println(err)
		return nil, errors.Wrap(err, "request failed")
	}

	res := vo.CreateDailyEnglishResp{}
	err := json.Unmarshal(output.Data, &res)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal fail")
	}

	return &res, nil
}
