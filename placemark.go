package gokml

import (
	"encoding/xml"
	"strings"
)

/*
 * Placemark
 */
type Placemark struct {
	XMLName     xml.Name `xml:"Placemark"`
	Id          string   `xml:"id,attr,omitempty"`
	Name        string   `xml:"name,omitempty"`
	Description string   `xml:"description,omitempty"`
	StyleUrl    string   `xml:"styleUrl,omitempty"`
	Coordinates string   `xml:"Point>coordinates"`
	Extended    ExtendedData
}

// NewPlacemark() creates a new placemark.  All parameters are strings.
func NewPlacemark(id, name, desc, lat, lon, alt, styleurl string) Placemark {
	p := Placemark{Id: id, Name: name, Description: desc}
	if styleurl != "" {
		p.StyleUrl = styleurl
	}
	p.SetCoordinates(lat, lon, alt)
	return p
}

// SetCoordinates takes latitude, longitude, and altitude
// and sets the Coordinates value of the Placemark.
func (p *Placemark) SetCoordinates(lat, lon, alt string) {
	s := []string{lon, lat, alt}
	p.Coordinates = strings.Join(s, ",")
}

// AddExtendedData() adds an ExtendedData structure to the Placemark
func (p *Placemark) AddExtendedData(e ExtendedData) {
	p.Extended = e
}

// Placemark.Marshal()
func (p *Placemark) Marshal() ([]byte, error) {
	return xml.MarshalIndent(p, "", "	")
}
