package fileCalssifier

import (
	"os"
	"strings"
	"archive/zip"
)


func SortZipFile(thePath string) {
	files, _ := os.ReadDir("./")
	var fileList []string
	dirPath := thePath

  for _, v := range files{
		if strings.Contains(v.Name(), ".zip") {
			fileList = append(fileList, v.Name())
		}
	}

	for i2, v2 := range fileList{
		unzippedFiles, err := Unzip()
	}



}
