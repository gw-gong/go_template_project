package app

// ==============================================================
// POST /app/router_group02/testx
// ==============================================================

type TestxRequest struct {
	Field01 string `json:"field01" binding:"required"`
	Field02 string `json:"field02" binding:"required"`
}

type TestxResponse struct {
	Field01 string `json:"field01"`
}
