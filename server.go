package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	var buf bytes.Buffer
	file, _, err := r.FormFile("filetoupload")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	// name := strings.Split(header.Filename, ".")
	io.Copy(&buf, file)

	contents := buf.String()
	b, _ := io.ReadAll(r.Response.Body)
	fmt.Println(string(b))

	f, err := os.Create("uploaded_file")
	if err != nil {
		panic(err)
	}

	n, err := f.WriteString(contents)
	fmt.Print(n)

}
