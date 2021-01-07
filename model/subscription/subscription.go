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
