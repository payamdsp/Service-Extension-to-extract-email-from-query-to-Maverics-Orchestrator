package main

import (
    "fmt"
    "net/http"

    "example.com/maverics/service"
    "github.com/strata-io/service-extension/orchestrator"
)

// mockOrchestratorAPI simulates the Orchestrator API interface for testing purposes.
type mockOrchestratorAPI struct{}

func (m *mockOrchestratorAPI) Logger() orchestrator.Logger {
    return &mockLogger{}
}

func (m *mockOrchestratorAPI) Session() (*orchestrator.Session, error) {
    return &orchestrator.Session{}, nil
}

type mockLogger struct{}

func (l *mockLogger) Info(component, message string, args ...interface{}) {
    fmt.Printf("[INFO] %s: %s
", component, message)
}

func (l *mockLogger) Debug(component, message string, args ...interface{}) {
    fmt.Printf("[DEBUG] %s: %s
", component, message)
}

func main() {
    http.HandleFunc("/create-header", func(w http.ResponseWriter, r *http.Request) {
        header, err := service.CreateEmailHeader(&mockOrchestratorAPI{}, w, r)
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

    fmt.Println("Orchestrator Service running at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}