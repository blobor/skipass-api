package bapi

import (
	"net/http"
)

type Client struct {
	base string
	http *http.Client
}

func NewClient(c *http.Client) *Client {
	return &Client{
		base: "http://tickets.bukovel.com",
		http: c,
	}
}
