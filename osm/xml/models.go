package xml

import "encoding/xml"

type Node struct {
	XMLName xml.Name `xml:"node"`
	Id int64 `xml:"id,attr"`
	Latitude float64 `xml:"lat,attr"`
	Longitude float64 `xml:"lon,attr"`
	Visible bool `xml:"visible,attr"`
	Version int16 `xml:"version,attr"`
	Changeset int64 `xml:"changeset,attr"`
	Tags []Tag `xml:"tag"`
}

type Tag struct {
	XMLName xml.Name `xml:"tag"`
	Key string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}