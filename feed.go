package main

import "encoding/xml"

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
