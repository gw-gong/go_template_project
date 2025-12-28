package api

type TestRequest01 struct {
	Field01 string `json:"field01" binding:"required"`
	Field02 string `json:"field02" binding:"required"`
}

type TestResponse01 struct {
	Field01 string `json:"field01"`
	Field02 string `json:"field02"`
}
