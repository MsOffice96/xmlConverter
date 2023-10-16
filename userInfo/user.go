package userInfo

import (
	f "fmt"
	"io/fs"
	"io/ioutil"

	"github.com/xuri/excelize/v2"
)

type UserInfo struct {
	FilePATH          string
	FilePATH_FileList []fs.FileInfo
	Select_FileNumber int
	Select_FileName   string
	ExcelSheet        int
}

func NewUserInfo() *UserInfo {
	userInfo := &UserInfo{}
	userInfo.SetFilePATH()
	userInfo.SetFileNumber()
	userInfo.SetExcelSheet()

	return userInfo
}

func (u *UserInfo) SetFilePATH() {
	for {
		f.Println("Enter Your Excel File Path :")
		f.Scan(&u.FilePATH)

		fileList, err := ioutil.ReadDir(u.FilePATH)
		if err == nil {
			u.FilePATH_FileList = fileList
			return
		} else {
			f.Printf("\nInvalid Excel File path retry")
		}
	}
}

func (u *UserInfo) SetFileNumber() {
	for {
		for fileIndex, file := range u.FilePATH_FileList {
			f.Printf("%d. %s\n", fileIndex, file.Name())
		}
		f.Println("\nSelect Your Excel File Number :")
		f.Scan(&u.Select_FileNumber)
		if u.Select_FileNumber < len(u.FilePATH_FileList) {
			// u.Select_FileName = u.FilePATH_FileList[u.Select_FileNumber].Name()
			u.SetFileName(u.FilePATH_FileList[u.Select_FileNumber].Name())
			break
		} else {
			f.Printf("\nInvalid Excel File Number retry")
		}
	}
}

func (u *UserInfo) SetFileName(FileName string) {
	u.Select_FileName = FileName
}

func (u *UserInfo) SetExcelSheet() {
	for {
		//
		full_Path := f.Sprintf("%s/%s", u.FilePATH, u.Select_FileName)
		f.Printf("===========\t %s ", full_Path)
		//

		excelfile, err := excelize.OpenFile(u.FilePATH + "/" + u.Select_FileName)
		if err != nil {
			f.Printf("excelfile error")
		}
		excel_sheets := excelfile.GetSheetList()
		f.Printf("\nSelect Your Excel Sheet: \n")
		for i, sheet := range excel_sheets {
			f.Printf("%d. %s \n", i, sheet)
		}
		f.Scan(&u.ExcelSheet)
		if u.ExcelSheet < len(excel_sheets) {
			break
		} else {
			f.Printf("Invalid Excel Sheet Number")
		}
	}
}

func (u *UserInfo) GetUserInfo() {
	f.Printf("File PATH: %s\n", u.FilePATH)
	f.Printf("File PATH File list: %+v \n", u.FilePATH_FileList)
	f.Printf("Select File Number: %d\n", u.Select_FileNumber)
	f.Printf("Select File Name: %s\n", u.Select_FileName)
	f.Printf("Select Excel Sheet : %d\n", u.ExcelSheet)
}
