package model

type Snippet struct {
	PublishedAt string     `json:"publishedAt"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ResourceID  ResourceID `json:"resourceId"`
	ChannelID   string     `json:"channelId"`
	Thumbnails  Thumbnails `json:"thumbnails"`
}
