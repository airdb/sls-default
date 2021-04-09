package vo

type (
	QueryRobotReq struct {
		Message string `form:"message"`
	}

	QueryRobotResp struct {
		Reunions []*Reunion `json:"Reunions"`
	}
)
