// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    subscription, err := UnmarshalSubscription(bytes)
//    bytes, err = subscription.Marshal()

package model

type Subscription struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	NextPageToken string   `json:"nextPageToken"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Item   `json:"items"`
}

type Item struct {
	Kind           string         `json:"kind"`
	Etag           string         `json:"etag"`
	ID             string         `json:"id"`
	Snippet        Snippet        `json:"snippet"`
	ContentDetails ContentDetails `json:"contentDetails"`
}

type ContentDetails struct {
	TotalItemCount int64  `json:"totalItemCount"`
	NewItemCount   int64  `json:"newItemCount"`
	ActivityType   string `json:"activityType"`
}

type Snippet struct {
	PublishedAt string     `json:"publishedAt"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ResourceID  ResourceID `json:"resourceId"`
	ChannelID   string     `json:"channelId"`
	Thumbnails  Thumbnails `json:"thumbnails"`
}

type ResourceID struct {
	Kind      string `json:"kind"`
	ChannelID string `json:"channelId"`
}

type Thumbnails struct {
	Default Default `json:"default"`
	Medium  Default `json:"medium"`
	High    Default `json:"high"`
}

type Default struct {
	URL string `json:"url"`
}

type PageInfo struct {
	TotalResults   int64 `json:"totalResults"`
	ResultsPerPage int64 `json:"resultsPerPage"`
}
