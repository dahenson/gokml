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
	Placemarks []Placemark `xml:"Placemark"`
	Folders    []Folder    `xml:"Folder"`
}

// ExtendedData for placemarks
type ExtendedData struct {
	XMLName xml.Name `xml:"ExtendedData"`
	Datas   []Data   `xml:"Data"`
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
func (k *Kml) MarshalIndent() ([]byte, error) {
	return xml.MarshalIndent(k, "", "	")
}

func (k *Kml) Marshal() ([]byte, error) {
	return xml.Marshal(k)
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
