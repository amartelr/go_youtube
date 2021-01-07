package model

type Thumbnails struct {
	Default Default `json:"default"`
	Medium  Default `json:"medium"`
	High    Default `json:"high"`
}
