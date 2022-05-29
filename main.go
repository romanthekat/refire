package main

import (
	"fmt"
	"net/http"
)

func main() {
	subreddits := getSubreddits()
	if len(subreddits) == 0 {
		fmt.Println("no subreddits configured, add subreddits and filter keywords in ~/.refire.json")
	}

	client := &http.Client{}

	for _, subreddit := range subreddits {
		printSubredditName(subreddit)
		feed := getFeed(client, "https://www.reddit.com/r/"+subreddit.Name+"/.rss")
		entries := filterFeedByKeywords(feed.Entries, subreddit.FilterKeywords)

		for _, entry := range entries {
			fmt.Printf("%s\n%s\n\n", entry.Title, entry.Link.Link)
		}
	}
}

func printSubredditName(subreddit Subreddit) {
	fmt.Println("\n\033[32m" + subreddit.Name + "\033[0m")
}
