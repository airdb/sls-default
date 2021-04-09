package vo

type Info struct {
	Cmd    string `json:"cmd"`
	Date   string `json:"date"`
	Detail string `json:"detail"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Person string `json:"person"`
}

type InfoListResp struct {
	Data    []Info `json:"data"`
	Success bool   `json:"success"`
}

var Tdata = `{
    "data": [
        {
            "id": "01",
            "name": "通用上线模板",
            "detail": "非云上产品请使用该模板",
            "cmd": "empty",
            "date": "2018-06-17 00:00:01",
            "person": "李小明"
        }
    ],
    "success": true
}`

func InfoList() string {
	return Tdata
}

func InfoListV2() *InfoListResp {
	resp := &InfoListResp{
		Success: true,
	}

	resp.Data = append(resp.Data, Info{
		Cmd:    "adb login",
		Date:   "2018-06-17 00:00:01",
		Detail: "非云上产品请使用该模板",
		ID:     "01",
		Name:   "hello",
		Person: "deancn",
	})

	return resp
}
