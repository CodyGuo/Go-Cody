package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const ReposURL = "https://api.github.com/users/CodyGuo/repos"

type ReposResult struct {
	Items []Repos
}

type Repos struct {
	ID       int
	Name     string
	FullName string `json:"full_name"`
}

func SearchRepos(urlStr string) (*ReposResult, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result []Repos
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	var reposResult ReposResult
	reposResult.Items = result
	return &reposResult, nil
}
func main() {
	result, err := SearchRepos(ReposURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Items)
	for _, item := range result.Items {
		fmt.Printf("{\n\t\"id:\" %d,\n"+
			"\t\"name:\" %s,\n"+
			"\t\"full_name:\" %s,\n},\n", item.ID, item.Name, item.FullName)
	}
}
