package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "World")
}

func main() {
	// Greet(os.Stdout, "Chipz")
	http.ListenAndServe(":8000", http.HandlerFunc(MyGreeterHandler))
}
