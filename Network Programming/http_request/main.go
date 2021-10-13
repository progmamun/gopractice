package main

import (
	"log"
	"net/http"
	"time"
)

/* resp, err := http.Get("http://example.com/")
resp, err := client.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
resp, err := client.Post("http://example.com/upload", "image/jpeg", &buf)
...
values := url.Values{"key": {"Value"}, "id": {"123"}}
resp, err := http.PostForm("http://example.com/form", values)
resp, err := client.PostForm("http://example.com/form", values)

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("Content-Type", "text/html")
resp, err := client.Do(req)
// ...
*/
type logTripper struct {
	http.RoundTripper
}

func (l logTripper) RoundTrip(r *http.Request) (*http.Response,
	error) {
	log.Println(r.URL)
	r.Header.Set("X-Log-Time", time.Now().String())
	return l.RoundTripper.RoundTrip(r)
}

func main() {
	client := http.Client{Transport: logTripper{http.DefaultTransport}}
	req, err := http.NewRequest("GET", "https://www.google.com/search?q=golang+net+http", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("Status code:", resp.StatusCode)
}
