package model

type ResourceID struct {
	Kind      string `json:"kind"`
	ChannelID string `json:"channelId"`
}