package main

import (
    "io"
    "os"
    "time"
    "net/http"
)

func hello(w http.ResponseWriter, r* http.Request) {
    io.WriteString(w,"Hello World from Pod " + os.Getenv("POD_NAME") + " - " + time.Now().Format("2006-01-02 15:04:05"))
}

func main() {
    http.HandleFunc("/",hello)
    http.ListenAndServe(":8080",nil)
}