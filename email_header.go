package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User struct from JSONPlaceholder
type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

// GetUserByID fetches a user from the placeholder API
func GetUserByID(id int) (string, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch user: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return user.Email, nil
}

// Simulate CreateEmailHeader from Maverics
func CreateEmailHeader() (http.Header, error) {
	log.Println("Building email custom header...")

	// Simulate Orchestrator session success
	_, err := simulateOrchestratorSession()
	if err != nil {
		return nil, fmt.Errorf("unable to simulate session: %w", err)
	}

	// Fetch user email
	email, err := GetUserByID(2)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user email: %w", err)
	}

	log.Printf("Retrieved email: %s\n", email)

	// Set header
	header := make(http.Header)
	header.Set("CUSTOM-EMAIL", email)

	return header, nil
}

// SimulateOrchestratorSession mocks a successful session
func simulateOrchestratorSession() (bool, error) {
	// Assume always successful for testing
	return true, nil
}

func main() {
	header, err := CreateEmailHeader()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Custom Header Set:")
	for key, value := range header {
		fmt.Printf("%s: %v\n", key, value)
	}
}
