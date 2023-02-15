package main

import(
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "./static/index.html")
}

func main() {
  fmt.Println("running...")
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
