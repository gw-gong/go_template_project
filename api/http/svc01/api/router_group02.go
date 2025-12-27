package api

type TestRequest02 struct {
	Field01 string `json:"field01" binding:"required"`
	Field02 string `json:"field02" binding:"required"`
}

type TestResponse02 struct {
	Field01 string `json:"field01"`
	Field02 string `json:"field02"`
}
