package userInfo

import (
	f "fmt"
	"io/fs"
	"io/ioutil"
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
			u.Select_FileName = u.FilePATH_FileList[u.Select_FileNumber].Name()
			break
		} else {
			f.Printf("\nInvalid Excel File Number retry")
		}
	}
}

func (u *UserInfo) SetFileName() {

}

func (u *UserInfo) SetExcelSheet() {

}
