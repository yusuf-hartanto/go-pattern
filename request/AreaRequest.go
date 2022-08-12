package request

type (
	AreaReq struct {
		Width  int64  `json:"width" binding:"required"`
		Height int64  `json:"height" binding:"required"`
		Type   string `json:"type" binding:"required"`
	}
)
