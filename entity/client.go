package entity

import "net/http"

// Client .
type Client struct {
	EndPoint   string
	HttpClient *http.Client
}
