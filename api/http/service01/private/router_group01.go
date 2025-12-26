package private

// ==============================================================
// POST /internal/router_group01/test
// ==============================================================

type TestRequest struct {
	Field01 string `json:"field01" binding:"required"`
	Field02 string `json:"field02" binding:"required"`
}

type TestResponse struct {
	Field01 string `json:"field01"`
}
