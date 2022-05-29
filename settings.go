package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Subreddit struct {
	Name           string   `json:"name,omitempty"`
	FilterKeywords []string `json:"filter_keywords"`
}

func getSubreddits() []Subreddit {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.Open(home + "/.refire.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	var subreddits []Subreddit
	err = json.Unmarshal(bytes, &subreddits)
	if err != nil {
		panic(err)
	}

	return subreddits
}
