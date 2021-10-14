package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

/*type Client struct {
	client *http.Client
	tick   *time.Ticker
}
type apiTransport struct {
	http.RoundTripper
	key string
}*/
type Result struct {
	AddressComponents []struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	} `json:"address_components"`
	FormattedAddress string `json:"formatted_address"`
	Geometry         struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		// more fields
	} `json:"geometry"`
	PlaceID string `json:"place_id"`
	// more fields
}

func (a apiTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	q := r.URL.Query()
	q.Set("key", a.key)
	r.URL.RawQuery = q.Encode()
	return a.RoundTripper.RoundTrip(r)
}

func NewClient(tick time.Duration, key string) *Client {
	return &Client{
		client: &http.Client{
			Transport: apiTransport{http.DefaultTransport, key},
		},
		tick: time.NewTicker(tick),
	}
}

// https://maps.googleapis.com/maps/api/geocode/json?latlng=40.714224,-73.961452&key=YOUR_API_KEY
// https://maps.googleapis.com/maps/api/geocode/json?latlng=%v,%v
const url = "https://maps.googleapis.com/maps/api/geocode/json?latlng=40.714224,-73.961452&key=AIzaSyDolwQtLcRXoAEbkNeH3ye8iLWOjxhhPqA"
<-c.tick.C
resp, err := c.client.Get(fmt.Sprintf(url, lat, lng))
if err != nil {
	return nil, err
}
defer resp.Body.Close()
var v struct {
	Results []Result `json:"results"`
	Status  string   `json:"status"`
}

func main() {
	// get the result
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	switch v.Status {
	case "OK":
		return v.Results, nil
	case "ZERO_RESULTS":
		return nil, nil
	default:
		return nil, fmt.Errorf("status: %q", v.Status)
	}

	c := NewClient(24*time.Hour/100000, os.Getenv("AIzaSyDolwQtLcRXoAEbkNeH3ye8iLWOjxhhPqA")) // ("your_api")
	start := time.Now()
	for _, l := range [][2]float64{
		{40.4216448, -3.6904040},
		{40.4163111, -3.7047328},
		{40.4123388, -3.7096724},
		{40.4145150, -3.7064412},
	} {
		locs, err := c.ReverseGeocode(l[0], l[1])
		e := time.Since(start)
		if err != nil {
			log.Println(e, l, err)
			continue
		}
		// just print the first location
		if len(locs) != 0 {
			locs = locs[:1]
		}
		log.Println(e, l, locs)
	}
}
