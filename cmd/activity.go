package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var client = &http.Client{}

type Event struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Actor struct {
		Login string `json:"login"`
	} `json:"actor"`
}

func FetchUserEvents(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	var events []Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	for _, event := range events {
		fmt.Printf("Event ID: %s, Type: %s, Repo: %s\n", event.ID, event.Type, event.Repo.Name)
	}
}
