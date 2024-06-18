# HTTP Request CLI tool

A Golang CLI tool made using [spf13/cobra](https://github.com/spf13/cobra) to make Get and Post requests.

### Basic server code to get started

```
package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/get", getHandler).Methods("GET")
	r.HandleFunc("/post", postHandler).Methods("POST")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}

// GET
func getHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Response to a GET request"
	w.Write([]byte(msg))
}

// POST
func postHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(body)
}
```
A basic Golang server using [Gorilla Mux](https://github.com/gorilla/mux) that handles two routes - `/get` and `/post`. The response to a Get request is a string 'Response to a GET request'. In the Post request, you can send a JSON object and the server will just return the same object back.

## How to run
Assuming the server is running on http://localhost:8080

Sort dependencies
```
go mod tidy
```

Build the binary or run directly
- To make a get request
```
go run main.go get http://localhost:8080/get
```
- To make a post request
```
go run main.go post https://localhost:8080/post -m=abcd
```
