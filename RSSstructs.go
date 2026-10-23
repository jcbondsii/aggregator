package main

import (
	"fmt"
	"strings"
	"time"
)

type RSSDate time.Time

func (d *RSSDate) UnmarshalText(text []byte) error {
	value := strings.TrimSpace(string(text))
	for _, layout := range []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
	} {
		parsed, err := time.Parse(layout, value)
		if err == nil {
			*d = RSSDate(parsed)
			return nil
		}
	}

	return fmt.Errorf("unsupported RSS date %q", value)
}

func (d RSSDate) Format(layout string) string {
	return time.Time(d).Format(layout)
}

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Description string    `xml:"description"`
		Link        string    `xml:"link"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string  `xml:"title"`
	Description string  `xml:"description"`
	Link        string  `xml:"link"`
	PubDate     RSSDate `xml:"pubDate"`
}
