// A small library aimed at writing simple kml files
package gokml

import (
	"encoding/xml"
)

const (
	DefaultNamespace = "http://www.opengis.net/kml/2.2"
)

// Kml is the base Kml document
type Kml struct {
	XMLName    xml.Name `xml:"kml"`
	Namespace  string   `xml:"xmlns,attr"`
	Document   Document
	Placemarks []Placemark
	Folders    []Folder
}

// Style for features
type Style struct {
	XMLName xml.Name `xml:"Style"`
	Id      string   `xml:"id,attr,omitempty"`
	Icon    IconStyle
}

// IconStyle
type IconStyle struct {
	XMLName xml.Name `xml:"IconStyle"`
	Scale   string   `xml:"scale,omitempty"`
	Heading string   `xml:"heading,omitempty"`
	Href    string   `xml:"Icon>href,omitempty"`
}

// ExtendedData for placemarks
type ExtendedData struct {
	XMLName xml.Name `xml:"ExtendedData"`
	Datas   []Data
}

// Data that goes in ExtendedData
type Data struct {
	XMLName xml.Name `xml:"Data"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value"`
}

// NewKml() creates a new Kml structure with the specified namespace
func NewKml(namespace string) Kml {
	k := Kml{}
	if namespace == "" {
		k.Namespace = DefaultNamespace
	} else {
		k.Namespace = namespace
	}
	return k
}

// AddDocument() adds a document to the kml structure
func (k *Kml) AddDocument(d Document) {
	k.Document = d
}

// AddPlacemark() adds a placemark to the document
func (k *Kml) AddPlacemark(p Placemark) {
	k.Placemarks = append(k.Placemarks, p)
}

// Marshal() returns a properly indented XML structure
func (k *Kml) Marshal() ([]byte, error) {
	return xml.MarshalIndent(k, "", "	")
}

func NewStyle(id string) Style {
	s := Style{Id: id}
	return s
}

func (s *Style) AddIconStyle(is IconStyle) {
	s.Icon = is
}

func NewIconStyle(scale, heading, href string) IconStyle {
	is := IconStyle{Scale: scale, Heading: heading, Href: href}
	return is
}

// NewExtendedData() creates new extended data
func NewExtendedData() ExtendedData {
	return ExtendedData{}
}

// AddData() adds data of {name, value} to the ExtendedData structure
func (e *ExtendedData) AddData(name, value string) {
	d := Data{Name: name, Value: value}
	e.Datas = append(e.Datas, d)
}
