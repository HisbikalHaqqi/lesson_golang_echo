/*
	DENGAN MENGGUNAAKAN GZIP DATA DI COMPRESS
 */
package main

import (
	"io"
	"net/http"
	"os"
)

func main(){
	mux := new(http.ServeMux)
	mux.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		f, err := os.Open("node.jpg")
		if err != nil {
			defer f.Close()
		}

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		_, err = io.Copy(writer, f)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
	})

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux
	server.ListenAndServe()
}