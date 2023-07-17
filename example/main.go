package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/majid-2002/go-http-sniffer"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, GET!"))
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, POST!"))
}

func PutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, PUT!"))
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, DELETE!"))
}

func HeadHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, HEAD!"))
}

func main() {
	getHandler := http.HandlerFunc(GetHandler)
	postHandler := http.HandlerFunc(PostHandler)
	putHandler := http.HandlerFunc(PutHandler)
	deleteHandler := http.HandlerFunc(DeleteHandler)
	headHandler := http.HandlerFunc(HeadHandler)

	mux := http.NewServeMux()
	mux.HandleFunc("/get", getHandler)
	mux.HandleFunc("/post", postHandler)
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/delete", deleteHandler)
	mux.HandleFunc("/head", headHandler)

	logger := httpsniffer.NewHTTPLogger(mux)

	addr := ":8080"
	fmt.Printf("Starting server and listening on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, logger))
}
