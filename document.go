package gokml

import (
	"encoding/xml"
)

/*
 * Container type Document
 */
type Document struct {
	XMLName     xml.Name `xml:"Document"`
	Id          string   `xml:"id,attr,omitempty"`
	Name        string   `xml:"name,omitempty"`
	Visibility  int      `xml:"visibility,omitempty"`
	Open        int      `xml:"open,omitempty"`
	Address     string   `xml:"address,omitempty"`
	PhoneNumber string   `xml:"phoneNumber,omitempty"`
	Description string   `xml:"description,omitempty"`
	Placemarks  []Placemark
	Folders     []Folder
	DocStyle    []Style
}

// NewDocument() creates a new document
func NewDocument(id, name string) Document {
	d := Document{Id: id, Name: name}
	return d
}

// AddPlacemark() adds a placemark to the document
func (d *Document) AddPlacemark(p Placemark) {
	d.Placemarks = append(d.Placemarks, p)
}

// AddStyle() adds a style to the document
func (d *Document) AddStyle(s Style) {
	d.DocStyle = append(d.DocStyle, s)
}

// AddFolder() adds a folder to the document
func (d *Document) AddFolder(f Folder) {
	d.Folders = append(d.Folders, f)
}

// Document.Marshal() returns the marshalled xml structure
func (d *Document) Marshal() ([]byte, error) {
	return xml.MarshalIndent(d, "", "	")
}
