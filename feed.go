package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Id      string   `xml:"id"`
	Link    Link     `xml:"link"`
	Title   string   `xml:"title"`

	Author    string `xml:"author"`
	Category  string `xml:"category"`
	Content   string
	Updated   string
	Published string
}

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

type Link struct {
	XMLName xml.Name `xml:"link"`
	Link    string   `xml:"href,attr"`
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
