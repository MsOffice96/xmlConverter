package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("./energyData.xlsx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sheets := f.GetSheetList()
	// log.Println("Sheet List: ", sheets)

	rows, err := f.GetRows(sheets[5])
	if err != nil {
		fmt.Println(err)
	}

	var SortExcel [][12]string

	for _, row := range rows {
		// fmt.Println(i, row)
		var colDatas [12]string
		for colDataNum, colData := range row {
			var renewal_colData string

			renewal_colData = strings.Replace(colData, " ", "_", -1)
			renewal_colData = strings.Replace(renewal_colData, "/", "_", -1)
			// colDatas[colDataNum] = colData
			colDatas[colDataNum] = renewal_colData
		}
		SortExcel = append(SortExcel, colDatas)
	}

	var MaxLevel int = 7

	var TotalXML []XmlLevel0

	// var TestXml []string

	for rowNum := 1; rowNum < len(SortExcel); rowNum++ {
		for colNum := 0; colNum < MaxLevel; colNum++ {
			// 마지막 col인지 체크
			IsFinalCol := false
			if SortExcel[rowNum][colNum+1] == "" {
				IsFinalCol = true
			}
			log.Println("Data : ", SortExcel[rowNum][colNum])
			// if SortExcel[rowNum][colNum+1] != " " { // 마지막 데이터가 아닌경우

			switch colNum {
			case 0:
				totalxmlExist := false
				for _, totalxml := range TotalXML {
					if totalxml.GetTag() == SortExcel[rowNum][colNum] {
						totalxmlExist = true
						break
					}
				}
				if !totalxmlExist {
					xml0 := &Xml0{
						Tag: SortExcel[rowNum][colNum],
					}
					TotalXML = append(TotalXML, xml0)
				}

			case 1:
				for _, totalxml := range TotalXML {
					if totalxml.GetTag() == SortExcel[rowNum][colNum-1] {
						level1 := totalxml.Findlevel1String(SortExcel[rowNum][colNum])
						if level1 == nil {
							xml1 := &Xml1{
								Tag: SortExcel[rowNum][colNum],
							}
							totalxml.SetLevel1(xml1)
						}
					}
				}

			case 2:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-2] {
						level1 := level0.Findlevel1String(SortExcel[rowNum][colNum-1])
						if !IsFinalCol {
							xmllevel2 := &Xml2{
								Tag: SortExcel[rowNum][colNum],
							}
							level2 := level1.FindLevel2(xmllevel2)
							if level2 == nil {
								level1.SetLevel2(xmllevel2)
							}
						} else {
							log.Println("Input End Data")
							xml2end := &Xml2End{
								Tag:     SortExcel[rowNum][colNum],
								endData: SortExcel[rowNum][8],
							}
							level1.SetLevel2(xml2end)
						}
					}
				}

			case 3:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-3] {
						level1 := level0.Findlevel1String(SortExcel[rowNum][colNum-2])
						if level1 == nil {
							log.Fatalln("e1")
						}
						level2 := level1.FindLevel2String(SortExcel[rowNum][colNum-1])

						if !IsFinalCol {
							xmllevel3 := &Xml3{
								Tag: SortExcel[rowNum][colNum],
							}
							level3 := level2.FindLevel3(xmllevel3)
							if level3 == nil {
								level2.SetLevel3(xmllevel3)
							}
						} else {
							xml3end := &Xml3End{
								Tag:     SortExcel[rowNum][colNum],
								endData: SortExcel[rowNum][8],
							}
							level2.SetLevel3(xml3end)
						}
					}
				}

			case 4:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-4] {

					}
				}

			}
			// } else { // 마지막 데이터인 경우

			// }

			if IsFinalCol {
				log.Println("=======> break")
				break
			}
		}

	}

	log.Println("TestXml: ", TotalXML[1].XmlOutPut())

}
