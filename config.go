package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const ConfigFilename = ".refire.json"

type Subreddit struct {
	Name           string   `json:"name,omitempty"`
	FilterKeywords []string `json:"filter_keywords"`
}

func getSubreddits() []Subreddit {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	path := home + "/" + ConfigFilename

	jsonFile, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		writeDefaultConfig(path)
		return nil
	} else if err != nil {
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

func writeDefaultConfig(path string) {
	defaultSubreddits := []Subreddit{{
		Name:           "subreddit_name",
		FilterKeywords: []string{"a keyword to show entry"},
	}}

	result, err := json.MarshalIndent(defaultSubreddits, "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(path, result, 0644)
	if err != nil {
		panic(err)
	}
}
