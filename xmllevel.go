package main

import (
	"fmt"
)

// xml Level0
type XmlLevel0 interface {
	GetTag() string
	SetLevel1(XmlLevel1)
	FindLevel1(XmlLevel1) (XmlLevel1, bool)
	Findlevel1String(string) XmlLevel1
	GetAllLevel1() []XmlLevel1
	XmlOutPut() string
}

type Xml0 struct {
	Tag       string
	xmlLevel1 []XmlLevel1
}

func (x *Xml0) GetTag() string {
	return x.Tag
}

func (x *Xml0) SetLevel1(x1 XmlLevel1) {
	x.xmlLevel1 = append(x.xmlLevel1, x1)
}

func (x *Xml0) FindLevel1(x1 XmlLevel1) (XmlLevel1, bool) {
	for _, xmllevel1 := range x.xmlLevel1 {
		if xmllevel1.GetTag() == x1.GetTag() {
			return xmllevel1, true
		}
	}
	return nil, false
}

func (x *Xml0) Findlevel1String(tag string) XmlLevel1 {
	for _, level1 := range x.xmlLevel1 {
		if level1.GetTag() == tag {
			return level1
		}
	}
	return nil
}

func (x *Xml0) GetAllLevel1() []XmlLevel1 {
	return x.xmlLevel1
}

func (x *Xml0) XmlOutPut() string {
	var TotalTag string
	TotalTag += fmt.Sprintf("<%s>", x.Tag)
	for _, level1tag := range x.xmlLevel1 {
		TotalTag += level1tag.XmlOutPut()
	}
	TotalTag += fmt.Sprintf("</%s>", x.Tag)
	return TotalTag
}

type Xml0End struct {
	Tag     string
	endData string
}

func (x *Xml0End) GetTag() string {
	return x.endData
}

// xml Level1
type XmlLevel1 interface {
	GetTag() string
	SetLevel2(XmlLevel2)
	FindLevel2(XmlLevel2) XmlLevel2
	FindLevel2String(string) XmlLevel2
	XmlOutPut() string
}

type Xml1 struct {
	Tag       string
	xmlLevel2 []XmlLevel2
}

func (x *Xml1) GetTag() string {
	return x.Tag
}

func (x *Xml1) SetLevel2(x1 XmlLevel2) {
	x.xmlLevel2 = append(x.xmlLevel2, x1)
}

func (x *Xml1) FindLevel2(x2 XmlLevel2) XmlLevel2 {
	// exist := false
	for _, xx2 := range x.xmlLevel2 {
		if xx2.GetTag() == x2.GetTag() {
			// exist = true
			return xx2
		}
	}
	return nil
}

func (x *Xml1) FindLevel2String(tag string) XmlLevel2 {
	for _, xx2 := range x.xmlLevel2 {
		if xx2.GetTag() == tag {
			return xx2
		}
	}
	return nil
}

func (x *Xml1) XmlOutPut() string {
	var TotalTag string
	TotalTag += fmt.Sprintf("<%s>", x.Tag)
	for _, level2 := range x.xmlLevel2 {
		TotalTag += level2.XmlOutPut()
	}
	TotalTag += fmt.Sprintf("</%s>", x.Tag)
	return TotalTag
}

type Xml1End struct {
	Tag     string
	endData string
}

func (x *Xml1End) GetTag() string {
	return x.endData
}

// xml Level2
type XmlLevel2 interface {
	GetTag() string
	SetLevel3(XmlLevel3)
	FindLevel3(XmlLevel3) XmlLevel3
	FindLevel3String(string) XmlLevel3
	XmlOutPut() string
}

type Xml2 struct {
	Tag       string
	xmlLevel3 []XmlLevel3
}

func (x *Xml2) GetTag() string {
	return x.Tag
}

func (x *Xml2) SetLevel3(x3 XmlLevel3) {
	x.xmlLevel3 = append(x.xmlLevel3, x3)
}

func (x *Xml2) FindLevel3(x3 XmlLevel3) XmlLevel3 {
	for _, xx3 := range x.xmlLevel3 {
		if xx3.GetTag() == x3.GetTag() {
			return xx3
		}
	}
	return nil
}

func (x *Xml2) FindLevel3String(tag string) XmlLevel3 {
	for _, xx3 := range x.xmlLevel3 {
		if xx3.GetTag() == tag {
			return xx3
		}
	}
	return nil
}

func (x *Xml2) XmlOutPut() string {
	var TotalTag string
	TotalTag += fmt.Sprintf("<%s>", x.Tag)
	for _, level3 := range x.xmlLevel3 {
		TotalTag += level3.XmlOutPut()
	}
	TotalTag += fmt.Sprintf("</%s>", x.Tag)
	return TotalTag
}

type Xml2End struct {
	Tag     string
	endData string
}

func (x *Xml2End) GetTag() string {
	return x.endData
}

func (x *Xml2End) SetLevel3(xml3 XmlLevel3) {

}

func (x *Xml2End) FindLevel3(x3 XmlLevel3) XmlLevel3 {
	return nil
}

func (x *Xml2End) FindLevel3String(tag string) XmlLevel3 {
	return nil
}

func (x *Xml2End) XmlOutPut() string {
	return fmt.Sprintf("<%s>%s</%s>", x.Tag, x.endData, x.Tag)

}

// xml Level3
type XmlLevel3 interface {
	GetTag() string
	SetLevel4(XmlLevel4)
	FindLevel4(XmlLevel4) XmlLevel4
	FindLevel4String(string) XmlLevel4
	XmlOutPut() string
}

type Xml3 struct {
	Tag       string
	xmlLevel4 []XmlLevel4
}

func (x *Xml3) GetTag() string {
	return x.Tag
}

func (x *Xml3) SetLevel4(x4 XmlLevel4) {
	x.xmlLevel4 = append(x.xmlLevel4, x4)
}

func (x *Xml3) FindLevel4(x4 XmlLevel4) XmlLevel4 {
	for _, xx4 := range x.xmlLevel4 {
		if xx4.GetTag() == x4.GetTag() {
			return xx4
		}
	}
	return nil
}

func (x *Xml3) FindLevel4String(tag string) XmlLevel4 {
	for _, xx4 := range x.xmlLevel4 {
		if xx4.GetTag() == tag {
			return xx4
		}
	}
	return nil
}

func (x *Xml3) XmlOutPut() string {
	var TotalTag string
	TotalTag += fmt.Sprintf("<%s>", x.Tag)
	for _, level4 := range x.xmlLevel4 {
		TotalTag += level4.XmlOutPut()
	}
	TotalTag += fmt.Sprintf("</%s>", x.Tag)
	return TotalTag
}

type Xml3End struct {
	Tag     string
	endData string
}

func (x *Xml3End) GetTag() string {
	return x.endData
}

func (x *Xml3End) SetLevel4(xml4 XmlLevel4) {

}

func (x *Xml3End) FindLevel4(x4 XmlLevel4) XmlLevel4 {
	return nil
}

func (x *Xml3End) FindLevel4String(tag string) XmlLevel4 {
	return nil
}

func (x *Xml3End) XmlOutPut() string {
	return fmt.Sprintf("<%s>%s</%s>", x.Tag, x.endData, x.Tag)
}

// Level 4
type XmlLevel4 interface {
	GetTag() string
	SetLevel5(XmlLevel5)
	FindLevel5(XmlLevel5) XmlLevel5
	FindLevel5String(string) XmlLevel5
	XmlOutPut() string
}

type Xml4 struct {
	Tag       string
	xmlLevel5 []XmlLevel5
}

func (x *Xml4) GetTag() string {
	return x.Tag
}

func (x *Xml4) SetLevel5(x5 XmlLevel5) {
	x.xmlLevel5 = append(x.xmlLevel5, x5)
}

func (x *Xml4) FindLevel5(x5 XmlLevel5) XmlLevel5 {
	for _, xx5 := range x.xmlLevel5 {
		if xx5.GetTag() == x5.GetTag() {
			return xx5
		}
	}
	return nil
}

func (x *Xml4) FindLevel5String(tag string) XmlLevel5 {
	for _, xx5 := range x.xmlLevel5 {
		if xx5.GetTag() == tag {
			return xx5
		}
	}
	return nil
}

func (x *Xml4) XmlOutPut() string {
	var TotalTag string
	TotalTag += fmt.Sprintf("<%s>", x.Tag)
	for _, level5 := range x.xmlLevel5 {
		TotalTag += level5.XmlOutPut()
	}
	TotalTag += fmt.Sprintf("</%s>", x.Tag)
	return TotalTag
}

type Xml4End struct {
	Tag     string
	endData string
}

func (x *Xml4End) GetTag() string {
	return x.endData
}

func (x *Xml4End) SetLevel5(xml5 XmlLevel5) {

}

func (x *Xml4End) FindLevel5(x5 XmlLevel5) XmlLevel5 {
	return nil
}

func (x *Xml4End) FindLevel5String(tag string) XmlLevel5 {
	return nil
}

func (x *Xml4End) XmlOutPut() string {
	return fmt.Sprintf("<%s>%s</%s>", x.Tag, x.endData, x.Tag)
}

// level 5
type XmlLevel5 interface {
	GetTag() string
	// SetLevel6(XmlLevel5)
	// FindLevel5(XmlLevel5) XmlLevel5
	// FindLevel5String(string) XmlLevel5
	XmlOutPut() string
}

type Xml5 struct {
	Tag string
	// xmlLevel5 []XmlLevel6
}

func (x *Xml5) GetTag() string {
	return x.Tag
}

func (x *Xml5) XmlOutPut() string {
	return fmt.Sprintf("<%s>%s</%s>", x.Tag, "???", x.Tag)
}

type Xml5End struct {
	Tag     string
	endData string
}

func (x *Xml5End) GetTag() string {
	return x.Tag
}

func (x *Xml5End) XmlOutPut() string {
	return fmt.Sprintf("<%s>%s</%s>", x.Tag, x.endData, x.Tag)
}

// Test All Include Data
type XmlLevelX interface {
	GetTag() string
	SetNextLevel()
}

type XmlX struct {
	Tag       string
	NextLevel []XmlX
}

type XmlEnd struct {
	Tag     string
	endData string
}
