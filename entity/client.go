package entity

import "net/http"

// Client .
type Client struct {
	EndPoint   string
	ApiKey     string
	HttpClient *http.Client
}
