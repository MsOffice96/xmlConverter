package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/xuri/excelize/v2"
)

// 1. Select filePATH
// 2. Select fileName
// 2. Select fileSheet
// 3. Show default Invalid Tag to convert
// 4. Input Change Tag for Invalid Tag from User
// 5. Input OutPut FileName ((*).txt) from User

// 6. Description : Invalid Tag, Unit, type, format
// if (Unit == empty) {not use Unit}
// if (Type == "") {not use Type}
// if (Format == "")

// 7. Process

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

		// for ii, data := range row {
		// 	log.Printf("row: %d, col: %d, data: %s", i, ii, data)
		// }

		var colDatas [12]string
		for colDataNum, colData := range row {

			if colDataNum <= 7 {
				var renewal_colData string

				renewal_colData = strings.Replace(colData, " ", "_", -1)
				renewal_colData = strings.Replace(renewal_colData, "/", "_", -1)

				// 임시 string 대체
				renewal_colData = strings.Replace(renewal_colData, "(", "_", -1)
				renewal_colData = strings.Replace(renewal_colData, ")", "_", -1)
				renewal_colData = strings.Replace(renewal_colData, "&", "_", -1)
				renewal_colData = strings.Replace(renewal_colData, "+", "_", -1)
				renewal_colData = strings.Replace(renewal_colData, "%", "_", -1)
				renewal_colData = strings.Replace(renewal_colData, "°", "_", -1)
				renewal_colData = strings.Replace(renewal_colData, "–", "_", -1)
				// Level Tag 에는 45가 들어가면 안됨. 따라서 어떻게 할지 상의 필요
				renewal_colData = strings.Replace(renewal_colData, "45_", "parameter", -1)

				// colDatas[colDataNum] = colData
				colDatas[colDataNum] = renewal_colData
			} else {

				colDatas[colDataNum] = colData
			}

		}
		SortExcel = append(SortExcel, colDatas)

	}

	var MaxLevel int = 7

	var TotalXML []XmlLevel0

	// var TestXml []string

	for rowNum := 1; rowNum < len(SortExcel); rowNum++ {
		for colNum := 0; colNum <= MaxLevel; colNum++ {

			// 마지막 col인지 체크
			IsFinalCol := false

			if SortExcel[rowNum][colNum+1] == "" || colNum == MaxLevel {
				IsFinalCol = true
			}

			// log.Println("Data : ", SortExcel[rowNum][colNum])

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
							xml2end := &Xml2End{
								Tag:     SortExcel[rowNum][colNum],
								endData: SortExcel[rowNum][8],
								unit:    SortExcel[rowNum][9],
								typee:   SortExcel[rowNum][10],
								format:  SortExcel[rowNum][11],
							}
							level1.SetLevel2(xml2end)
						}
					}
				}

			case 3:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-3] {
						level1 := level0.Findlevel1String(SortExcel[rowNum][colNum-2])
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
								unit:    SortExcel[rowNum][9],
								typee:   SortExcel[rowNum][10],
								format:  SortExcel[rowNum][11],
							}
							level2.SetLevel3(xml3end)
						}
					}
				}

			case 4:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-4] {
						level1 := level0.Findlevel1String(SortExcel[rowNum][colNum-3])
						level2 := level1.FindLevel2String(SortExcel[rowNum][colNum-2])
						level3 := level2.FindLevel3String(SortExcel[rowNum][colNum-1])
						if !IsFinalCol {
							xmllevel4 := &Xml4{
								Tag: SortExcel[rowNum][colNum],
							}
							level4 := level3.FindLevel4(xmllevel4)
							if level4 == nil {
								level3.SetLevel4(xmllevel4)
							}
						} else {
							xml4end := &Xml4End{
								Tag:     SortExcel[rowNum][colNum],
								endData: SortExcel[rowNum][8],
								unit:    SortExcel[rowNum][9],
								typee:   SortExcel[rowNum][10],
								format:  SortExcel[rowNum][11],
							}
							level3.SetLevel4(xml4end)
						}
					}
				}

			case 5:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-5] {
						level1 := level0.Findlevel1String(SortExcel[rowNum][colNum-4])
						level2 := level1.FindLevel2String(SortExcel[rowNum][colNum-3])
						level3 := level2.FindLevel3String(SortExcel[rowNum][colNum-2])
						level4 := level3.FindLevel4String(SortExcel[rowNum][colNum-1])
						if !IsFinalCol {
							xmllevel5 := &Xml5{
								Tag: SortExcel[rowNum][colNum],
							}
							level5 := level4.FindLevel5(xmllevel5)
							if level5 == nil {
								level4.SetLevel5(xmllevel5)
							}
						} else {
							xml5end := &Xml5End{
								Tag:     SortExcel[rowNum][colNum],
								endData: SortExcel[rowNum][8],
								unit:    SortExcel[rowNum][9],
								typee:   SortExcel[rowNum][10],
								format:  SortExcel[rowNum][11],
							}
							level4.SetLevel5(xml5end)
						}
					}
				}
			case 6:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-6] {
						level1 := level0.Findlevel1String(SortExcel[rowNum][colNum-5])
						level2 := level1.FindLevel2String(SortExcel[rowNum][colNum-4])
						level3 := level2.FindLevel3String(SortExcel[rowNum][colNum-3])
						level4 := level3.FindLevel4String(SortExcel[rowNum][colNum-2])
						level5 := level4.FindLevel5String(SortExcel[rowNum][colNum-1])
						if !IsFinalCol {
							xmllevel6 := &Xml6{
								Tag: SortExcel[rowNum][colNum],
							}
							level6 := level5.FindLevel6(xmllevel6)
							if level6 == nil {
								level5.SetLevel6(xmllevel6)
							}
						} else {
							xml6end := &Xml6End{
								Tag:     SortExcel[rowNum][colNum],
								endData: SortExcel[rowNum][8],
								unit:    SortExcel[rowNum][9],
								typee:   SortExcel[rowNum][10],
								format:  SortExcel[rowNum][11],
							}
							level5.SetLevel6(xml6end)
						}
					}
				}

			case 7:
				for _, level0 := range TotalXML {
					if level0.GetTag() == SortExcel[rowNum][colNum-7] {
						level1 := level0.Findlevel1String(SortExcel[rowNum][colNum-6])
						level2 := level1.FindLevel2String(SortExcel[rowNum][colNum-5])
						level3 := level2.FindLevel3String(SortExcel[rowNum][colNum-4])
						level4 := level3.FindLevel4String(SortExcel[rowNum][colNum-3])
						level5 := level4.FindLevel5String(SortExcel[rowNum][colNum-2])
						level6 := level5.FindLevel6String(SortExcel[rowNum][colNum-1])
						// if !IsFinalCol {
						//// Level7은 현재 무조건 마지막 Tag임
						// }
						xml7end := &Xml7End{
							Tag:     SortExcel[rowNum][colNum],
							endData: SortExcel[rowNum][8],
							unit:    SortExcel[rowNum][9],
							typee:   SortExcel[rowNum][10],
							format:  SortExcel[rowNum][11],
						}
						level6.SetLevel7(xml7end)

					}
				}

			}

			if IsFinalCol {
				// log.Println("=======> break")
				break
			}

		}

	}

	log.Println("TestXml: ", TotalXML[2].XmlOutPut())

	var total_xml string
	total_xml += fmt.Sprintf("<GRschema>")
	for _, level0 := range TotalXML {
		total_xml += level0.XmlOutPut()
	}
	total_xml += fmt.Sprintf("</GRschema>")

	log.Printf("\n\n%s\n\n", total_xml)

	// Write file
	ResultFile, err := os.Create("xmlResult.txt")
	if err != nil {
		log.Fatalln("Create Result File Error")
	}

	resultWriter := bufio.NewWriter(ResultFile)

	resultwriter_bytes, err := resultWriter.WriteString(total_xml)
	if err != nil {
		log.Fatalln("resultWriter.WriteString error")
	} else {
		log.Printf("resultWriter Write %d bytes", resultwriter_bytes)
		resultWriter.Flush()
	}

}