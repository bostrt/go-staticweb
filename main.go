package main

import (
  "net/http"
  "fmt"
)

// Thanks https://github.com/golang/go/wiki/HttpStaticFiles
func main() {
  fmt.Println("Listening on port 8080...")
  panic(http.ListenAndServe(":8080", http.FileServer(http.Dir("./files"))))
}
