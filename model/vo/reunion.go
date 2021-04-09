package vo

type ListReunionReq struct {
	PageNo   uint `form:"pageNo"`
	PageSize uint `form:"pageSize"`
}

type ListReunionResp struct {
	Reunions []*Reunion `json:"Reunions"`
}

type Reunion struct {
	Title  string   `json:"title"`
	Images []string `json:"images"`
}

func ListReunion(req *ListReunionReq) *ListReunionResp {
	var resp ListReunionResp

	// losts := po.ListLost(req.PageNo, req.PageSize)

	resp.Reunions = append(resp.Reunions, &Reunion{
		Title: "有种幸福叫当中秋遇上国庆，愿你拥有加长版的幸福，愿你拥有加大型的快乐，祝大家双节愉快！！！🎉🎉[跳跳][跳跳][愉快][愉快]",
		Images: []string{
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/1.jpeg",
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/2.jpeg",
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/3.jpeg",
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/4.jpeg",
		},
	})

	resp.Reunions = append(resp.Reunions, &Reunion{
		Title: "2 - 有种幸福叫当中秋遇上国庆，愿你拥有加长版的幸福，愿你拥有加大型的快乐，祝大家双节愉快！！！🎉🎉[跳跳][跳跳][愉快][愉快]",
		Images: []string{
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/1.jpeg",
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/2.jpeg",
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/3.jpeg",
			"https://wechat-1251018873.cos.ap-shanghai.myqcloud.com/mina/reunion_moment/2020/1001/1/4.jpeg",
		},
	})

	return &resp
}
