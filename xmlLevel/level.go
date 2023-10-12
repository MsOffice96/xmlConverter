package xmlLevel

import (
	"fmt"
)

type XmlLevel interface {
	GetTag() string
	SetNextLevel(XmlLevel)
	GetNextLevel(XmlLevel) XmlLevel
	GetNextLevelByTag(string) XmlLevel
	XmlOutPut() string
}

type Xml struct {
	Tag        string
	NextLevels []XmlLevel
}

func (x *Xml) GetTag() string {
	return x.Tag
}

func (x *Xml) SetNextLevel(nextXml XmlLevel) {
	x.NextLevels = append(x.NextLevels, nextXml)
}

func (x *Xml) GetNextLevel(nextXml XmlLevel) XmlLevel {
	for _, myNextlevel := range x.NextLevels {
		if myNextlevel.GetTag() == nextXml.GetTag() {
			return myNextlevel
		}
	}
	return nil
}

func (x *Xml) GetNextLevelByTag(tag string) XmlLevel {
	for _, myNextlevel := range x.NextLevels {
		if myNextlevel.GetTag() == tag {
			return myNextlevel
		}
	}
	return nil
}

func (x *Xml) XmlOutPut() string {
	var TotalTag string
	TotalTag += fmt.Sprintf("<%s>", x.Tag)
	for _, nextlevel := range x.NextLevels {
		TotalTag += nextlevel.XmlOutPut()
	}
	TotalTag += fmt.Sprintf("</%s>", x.Tag)
	return TotalTag
}

type XmlEnd struct {
	Tag     string
	endData string
	unit    string
	typee   string
	format  string
}

func (x *XmlEnd) GetTag() string {
	return x.endData
}

func (x *XmlEnd) SetNextLevel(nextlevel XmlLevel) {

}

func (x *XmlEnd) GetNextLevel(nextXml XmlLevel) XmlLevel {
	return nil
}

func (x *XmlEnd) GetNextLevelByTag(tag string) XmlLevel {
	return nil
}

func (x *XmlEnd) XmlOutPut() string {
	var xmlFrontTag string
	xmlFrontTag += fmt.Sprintf("<%s", x.Tag)
	if x.typee != "" {
		xmlFrontTag += fmt.Sprintf(" type=\"%s\"", x.typee)
	}
	if x.unit != "" {
		xmlFrontTag += fmt.Sprintf(" unit=\"%s\"", x.unit)
	}
	if x.format != "" {
		xmlFrontTag += fmt.Sprintf(" format=\"%s\"", x.format)
	}

	return xmlFrontTag + fmt.Sprintf(">%s</%s>", x.endData, x.Tag)
}
