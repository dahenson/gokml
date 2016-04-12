package gokml

import "encoding/xml"

const (
	IconPaddleBase = "http://maps.google.com/mapfiles/kml/paddle/"
	IconShapesBase = "http://maps.google.com/mapfiles/kml/shapes/"
)

// Style for features
type Style struct {
	XMLName xml.Name `xml:"Style"`
	Id      string   `xml:"id,attr,omitempty"`
	Icon    IconStyle
	Label   LabelStyle
}

// IconStyle
type IconStyle struct {
	XMLName xml.Name `xml:"IconStyle"`
	Scale   string   `xml:"scale,omitempty"`
	Heading string   `xml:"heading,omitempty"`
	Href    string   `xml:"Icon>href,omitempty"`
}

type LabelStyle struct {
	XMLName xml.Name `xml:"LabelStyle"`
	Scale   string   `xml:"scale,omitempty"`
	Color   string   `xml:"color,omitempty"`
}

func NewStyle(id string) Style {
	s := Style{Id: id}
	return s
}

func (s *Style) AddIconStyle(is IconStyle) {
	s.Icon = is
}

func (s *Style) AddLabelStyle(ls LabelStyle) {
	s.Label = ls
}

func NewIconStyle(scale, heading, href string) IconStyle {
	is := IconStyle{Scale: scale, Heading: heading, Href: href}
	return is
}

func NewLabelStyle(scale string) LabelStyle {
	ls := LabelStyle{Scale: scale}
	return ls
}
