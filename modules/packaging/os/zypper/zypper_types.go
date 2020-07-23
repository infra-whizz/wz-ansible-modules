package main

import "encoding/xml"

type ZypperSearch struct {
	XMLName xml.Name `xml:"stream"`
	Result  ZypperSearchResult
}

type ZypperSearchResult struct {
	XMLName      xml.Name `xml:"search-result"`
	Version      string   `xml:"version,attr"`
	SolvableList ZypperSolvableList
}

type ZypperSolvableList struct {
	XMLName  xml.Name         `xml:"solvable-list"`
	Packages []ZypperSolvable `xml:"solvable"`
}

type ZypperSolvable struct {
	XMLName xml.Name `xml:"solvable"`
	Name    string   `xml:"name,attr"`
}

type ZypperProgress struct {
	XMLName xml.Name `xml:"progress"`
	Id      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type ZypperDownload struct {
	XMLName xml.Name `xml:"download"`
	Url     string   `xml:"url,attr"`
}

type ZypperMessage struct {
	XMLName xml.Name `xml:"message"`
	Type    string   `xml:"type,attr"`
	Text    string   `xml:",chardata"`
}
