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
