package handler

import (
  "fmt"
  "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
  var name string = "My New Friend"
  queries := r.URL.Query()
  if val, ok := queries["name"]; ok {
    name = val[0]
  }
  fmt.Fprintf(w, fmt.Sprintf("Holla! %s", name))
}
