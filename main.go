package main

import (
	"fmt"
	"net/http"
)

func main() {
	subreddits := getSubreddits()
	if len(subreddits) == 0 {
		fmt.Println("no subreddits configured, add subreddits and filter keywords in ~/", ConfigFilename)
	}

	client := &http.Client{}

	for _, subreddit := range subreddits {
		printSubredditName(subreddit)
		feed := getFeed(client, "https://www.reddit.com/r/"+subreddit.Name+"/.rss")
		entries := filterFeedByKeywords(feed.Entries, subreddit.FilterKeywords)

		for _, entry := range entries {
			fmt.Printf("%s\n\t%s\n\n", entry.Title, entry.Link.Link)
		}
	}
}

func printSubredditName(subreddit Subreddit) {
	colorPrefix := "\u001B[32m" //green
	colorSuffix := "\033[0m"
	//if runtime.GOOS == "windows" {
	//	colorPrefix = ""
	//	colorSuffix = ""
	//}

	fmt.Printf("\n%s- %s%s\n", colorPrefix, subreddit.Name, colorSuffix)
}
