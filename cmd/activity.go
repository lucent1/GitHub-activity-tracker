package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var client = &http.Client{}

type Event struct {
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
	Repo      struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload struct {
		Size    int `json:"size"`
		Commits []struct {
			Sha     string `json:"sha"`
			Message string `json:"message"`
			Author  struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"commits"`
	} `json:"payload"`
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

	// Check the status code (200)
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
		if event.Type == "PushEvent" {
			fmt.Printf("Push event in repo: %s on %s\n", event.Repo.Name, event.CreatedAt)
			fmt.Printf("Number of commits: %d\n", event.Payload.Size)

			for _, commit := range event.Payload.Commits {
				fmt.Printf("Commit: %s - %s by %s\n", commit.Sha, commit.Message, commit.Author.Name)
			}
		}
	}
}
