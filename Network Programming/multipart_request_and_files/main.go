package main

/*
MIME-Version: 1.0
Content-Type: multipart/mixed; boundary=xxxx
This part before boundary is ignored
--xxxx
Content-Type: text/plain
First part of the message. The next part is binary data encoded in base64
--xxxx
Content-Type: application/octet-stream
Content-Transfer-Encoding: base64
PGh0bWw+CiAgPGhlYWQ+CiAgPC9oZWFkPgogIDxib2R5PgogICAgPHA+VGhpcyBpcyB0aGUg
Ym9keSBvZiB0aGUgbWVzc2FnZS48L3A+CiAgPC9ib2R5Pgo8L2h0bWw+Cg==
--xxxx--
*/
const (
	param = "file"
	endpoint = "/upload"
	content = `<html><body>` +
	`<form enctype="multipart/form-data" action="%s" method="POST">` +
	`<input type="file" name="%s"/><input type="submit"
	value="Upload"/>` +
	`</form></html></body>`
	)

	mux.HandleFunc(endpoint, func(w http.ResponseWriter, r
		*http.Request) {
		if r.Method == "GET" {
		fmt.Fprintf(w, content, endpoint, param)
		return
		} else if r.Method != "POST" {
		http.NotFound(w, r)
		return
		}
		path, err := upload(r)
		if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		}
		fmt.Fprintf(w, "Uploaded to %s", path)
		})
		func upload(r *http.Request) (string, error) {
			f, h, err := r.FormFile(param)
			if err != nil {
			return "", err
			}
			defer f.Close()
			p := filepath.Join(os.TempDir(), h.Filename)
			fw, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
			return "", err
			}
			defer fw.Close()
			if _, err = io.Copy(fw, f); err != nil {
			return "", err
			}
			return p, nil
		}					