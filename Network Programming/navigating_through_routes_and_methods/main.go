package main

import "fmt"

/*switch r.Method {
case http.MethodGet:
// GET implementation
case http.MethodPost:
// POST implementation
default:
http.NotFound(w, r)
}
*/

type methodHandler map[string]http.Handler
func (m methodHandler) ServeHTTP(w http.ResponseWriter, r
*http.Request) {
h, ok := m[strings.ToUpper(r.Method)]
if !ok {
http.NotFound(w, r)
return
}
h.ServeHTTP(w, r)
}
func main() {
	http.HandleFunc("/path1", methodHandler{
	http.MethodGet: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Showing record")
	}),
	http.MethodPost: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Updated record")
	}),
	})
	if err := http.ListenAndServe(":3000", nil); err != nil {
	log.Fatal(err)
	}
}
	