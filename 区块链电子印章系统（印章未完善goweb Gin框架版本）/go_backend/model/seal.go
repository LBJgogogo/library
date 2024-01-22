package model

type SealFrom struct {
	Filename  string
	ImgBase64 string
}
type SealRecord struct {
	Id        int    `json:"id"`
	Filename  string `json:"filename"`
	Type      string `json:"type"`
	Code      string `json:"code"`
	Datetime  string `json:"datetime"`
	ImgBase64 string `json:"imgBase64"`
}
type SealVerifyMessage struct {
	Name         string `json:"name"`
	CardId       string `json:"cardId"`
	ChainAccount string `json:"chainAccount"`
	Filename     string `json:"filename"`
	Code         string `json:"code"`
	Datetime     string `json:"datetime"`
}
type SealBase64 struct {
	ImgBase64 string `json:"imgBase64"`
}
