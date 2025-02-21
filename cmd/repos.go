/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Repository struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	} `json:"owner"`
	Private         bool   `json:"private"`
	HTMLURL         string `json:"html_url"`
	Description     string `json:"description"`
	Fork            bool   `json:"fork"`
	Language        string `json:"language"`
	StargazersCount int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
}

// reposCmd represents the repos command
var reposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Fetch user repositories",
	Run: func(cmd *cobra.Command, args []string) {
		if username == "" {
			fmt.Println("Please provide a GitHub username using --user")
			return
		}
		fetchUserRepos(username)
	},
}

func init() {
	reposCmd.Flags().StringVarP(&username, "user", "u", "", "GitHub username")
	rootCmd.AddCommand(reposCmd)
}

func fetchUserRepos(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %s\n", resp.Status)
		return
	}

	var repos []Repository
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	for _, repo := range repos {
		fmt.Printf("Repo: %s, Stars: %d, Forks: %d, Language: %s\n",
			repo.FullName, repo.StargazersCount, repo.ForksCount, repo.Language)
	}

	fmt.Println()
	fmt.Println("Total amount of repos:", len(repos))
}
