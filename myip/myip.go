package myip

import (
    "fmt"
    _ "strings"
    "net/http"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "my ip: %s", r.RemoteAddr)
}

