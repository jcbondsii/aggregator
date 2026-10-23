package main

import (
	"encoding/xml"
	"testing"
	"time"
)

func TestRSSDateUnmarshalText(t *testing.T) {
	var feed RSSFeed
	if err := xml.Unmarshal([]byte(`<rss><channel><item><pubDate>Sun, 08 Jan 2023 00:00:00 +0000</pubDate></item></channel></rss>`), &feed); err != nil {
		t.Fatalf("unmarshal RSS date: %v", err)
	}

	want := "Sun, 08 Jan 2023 00:00:00 +0000"
	if got := feed.Channel.Item[0].PubDate.Format(time.RFC1123Z); got != want {
		t.Errorf("PubDate.Format() = %q, want %q", got, want)
	}
}
