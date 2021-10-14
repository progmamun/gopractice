package main

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r
	*http.Request) {
	fmt.Fprint(w, "Hello!")
	})
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
	log.Fatal("ListenAndServe: ", err)
	}
	}
	