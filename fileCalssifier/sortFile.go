package fileCalssifier

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)


func SortZipFile(thePath string) {

	for i2, v2 := range fileList{
		unzippedFiles, err := Unzip()
	}

}

func zipSearcher(thePath string)([]string, string){
	files, _ := os.ReadDir("./")
	var fileList []string
	dirPath := thePath

  for _, v := range files{
		if strings.Contains(v.Name(), ".zip") {
			fileList = append(fileList, v.Name())
		}
	}

	return fileList, dirPath
}

func zipExtracter(filePaths []string) error {

	for _, fp := range filePaths{
		destPath := fp + "/extracted/"
		images, err := zip.OpenReader(fp)
		if err != nil{
			return err
		}

		defer images.Close()


		for _, f := range images.File{
			image, err := f.Open()
			if err != nil{
				return err
			}
			defer image.Close()

			if f.FileInfo().IsDir(){
				path := filepath.Join(destPath, f.Name)
				os.MkdirAll(path, f.Mode())
			}else{
				buf := make([]byte, f.UncompressedSize)
				_, err = io.ReadFull(image, buf)
				if err != nil{
					return err
				}

				path := filepath.Join(destPath, f.Name)
				err := os.WriteFile(path, buf, f.Mode())
				if err != nil{
					return err
				}
			}
		}
	}

	return nil
}
