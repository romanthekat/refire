package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	subreddits := getSubreddits()
	if len(subreddits) == 0 {
		fmt.Println("no subreddits configured, add subreddits and filter keywords in ~/.refire.json")
	}

	client := &http.Client{}

	for _, subreddit := range subreddits {
		//fmt.Println("\n---- " + subreddit.Name)
		feed := getFeed(client, "https://www.reddit.com/r/"+subreddit.Name+"/.rss")
		entries := filterFeedByKeywords(feed.Entries, subreddit.FilterKeywords)

		for _, entry := range entries {
			fmt.Println(entry.Title, "\n", entry.Link.Link, "\n")
		}
	}
}

func filterFeedByKeywords(entries []Entry, keywords []string) []Entry {
	var result []Entry

	for _, entry := range entries {
		for _, keyword := range keywords {
			title := strings.ToLower(entry.Title)
			if strings.Contains(title, strings.ToLower(keyword)) {
				result = append(result, entry)
				break
			}
		}
	}

	return result
}

func getFeed(client *http.Client, url string) Feed {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "refire/1.0")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("response status for %d is not OK, response body: %s\n", resp.StatusCode, string(bodyBytes))
	}

	var feed Feed
	err = xml.Unmarshal(bodyBytes, &feed)
	if err != nil {
		panic(err)
	}

	return feed
}
