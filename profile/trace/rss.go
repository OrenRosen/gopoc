package main

import "encoding/xml"

type ItemResponse struct {
	Title           string                  `xml:"title"`
	Link            string                  `xml:"link"`
	Author          string                  `xml:"author"`
	Guid            GUIDResponse            `xml:"guid"`
	Description     string                  `xml:"description"`
	PublicationDate string                  `xml:"pubDate"`
	Media           MediaResponse           `xml:"media:thumbnail"`
	Content         *EncodedContentResponse `xml:"content:encoded,omitempty"`
}

type MediaResponse struct {
	Thumbnail string `xml:"url,attr"`
}

type GUIDResponse struct {
	IsPermaLink bool   `xml:"isPermaLink,attr"`
	GUID        string `xml:",chardata"`
}

type EncodedContentResponse struct {
	Content string `xml:",cdata"`
}

type AtomLinkResponse struct {
	Rel  string `xml:"rel,attr"`
	Type string `xml:"type,attr"`
	Href string `xml:"href,attr"`
}

type rssResponseXML struct {
	XMLName       xml.Name         `xml:"rss"`
	Version       string           `xml:"version,attr"`
	Atom          string           `xml:"xmlns:atom,attr"`
	MediaSchema   string           `xml:"xmlns:media,attr"`
	ContentSchema string           `xml:"xmlns:content,attr"`
	AtomLink      AtomLinkResponse `xml:"channel>atom:link"`
	Description   string           `xml:"channel>description"`
	Title         string           `xml:"channel>title"`
	Link          string           `xml:"channel>link"`
	Items         []ItemResponse   `xml:"channel>item"`
}

type RSSResponseXML struct {
	RSSLink     string
	Title       string
	Link        string
	Description string
	Items       []ItemResponse
}

