package main

import (
    "fmt"
    "net/http"

    "example.com/maverics/service"
)

func main() {
    http.HandleFunc("/create-header", func(w http.ResponseWriter, r *http.Request) {
        header, err := service.CreateEmailHeader(nil, w, r)
        if err != nil {
            http.Error(w, fmt.Sprintf("Failed to create email header: %v", err), http.StatusInternalServerError)
            return
        }

        for key, values := range header {
            for _, value := range values {
                w.Header().Add(key, value)
            }
        }

        fmt.Fprintf(w, "Email header created and added to response.")
    })

    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}