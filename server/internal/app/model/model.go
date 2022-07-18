package model

type Request struct {
	Str  string `json:"str"`
	Char string `json:"char"`
}

type Response struct {
	Count int `json:"count"`
}
