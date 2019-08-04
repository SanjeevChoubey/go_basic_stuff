package main

import (
	"net/url"
)

type Client struct {
	BaseURL *url.URL
	Icao    []string
	Result  map[string]interface{}
}

type Report struct {
	StationID string
	wind      Wind
	weather   Weather
}

type Weather struct{}

type Wind struct {
	IsCalm bool
	Speed  string
	Gust   string
}
