package vo

type Tree struct {
	Title    string  `json:"title"`
	Name     string  `json:"name"`
	Spread   bool    `json:"spread"`
	Icon     string  `json:"icon"`
	Children *[]Tree `json:"children"`
}

// Icon: https://baidu.github.io/NoahV/#/doc/component/icon.
func ListTree() *Tree {
	root := Tree{
		Title:  "root",
		Name:   "root",
		Icon:   "DNS",
		Spread: true,
	}

	root.Children = &[]Tree{
		{
			Title:  "airdb",
			Name:   "airdb",
			Icon:   "home",
			Spread: true,
			Children: &[]Tree{
				{
					Title:    "host",
					Name:     "host",
					Spread:   true,
					Icon:     "star",
					Children: nil,
				},
				{
					Title:    "oauth2",
					Name:     "oauth2",
					Icon:     "client",
					Children: nil,
				},
				{
					Title:    "xadmin",
					Name:     "xadmin",
					Icon:     "client",
					Children: nil,
				},
			},
		},
		{
			Title:    "ngo",
			Name:     "ngo",
			Icon:     "home",
			Children: nil,
		},
		{
			Title:    "srehub",
			Name:     "srehub",
			Icon:     "DNS",
			Children: nil,
		},
		{
			Title: "bbhj",
			Name:  "bbhj",
			Icon:  "home",
			Children: &[]Tree{
				{
					Title:    "node1-2-1",
					Name:     "node1-2-1",
					Icon:     "star",
					Children: nil,
				},
				{
					Title:    "node1-2-2",
					Name:     "node1-2-2",
					Icon:     "client",
					Children: nil,
				},
			},
		},

		{
			Title:    "osfun",
			Name:     "osfun",
			Icon:     "console",
			Children: nil,
		},
		{
			Title:    "studygolang",
			Name:     "studygolang",
			Icon:     "th-large",
			Children: nil,
		},
		{
			Title:    "bfe",
			Name:     "bfe",
			Icon:     "network-topology",
			Children: nil,
		},
	}

	return &root
}

type UserResp struct {
	Success bool `json:"success"`
	Data    struct {
		Username string `json:"userName"`
		UserName string `json:"user-name"`
	} `json:"data"`
}

func QueryUser() *UserResp {
	var resp UserResp

	resp.Success = true
	resp.Data.Username = "deancn"
	resp.Data.Username = "deancn"

	return &resp
}
