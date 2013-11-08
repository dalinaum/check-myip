package myip

import (
    "fmt"
    _ "strings"
    "net/http"
)

func init() {
    http.HandleFunc("/raw", raw)
    http.HandleFunc("/", home)
}

func raw(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, r.RemoteAddr)
}

func home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "my ip: %s", r.RemoteAddr)
}
