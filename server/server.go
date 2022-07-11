package main

import (
	"fmt"
	"io"
	"net/http"
)

const keyServerAddr = "localhost"

func GetRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got \"/\" request\n", ctx.Value(keyServerAddr))

	// fmt.Printf("got \"/\" request\n")
	io.WriteString(w, "This is my website!\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Printf("%s: got \"/hello\" request\n", ctx.Value(keyServerAddr))

	// fmt.Printf("got \"/hello\" request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
