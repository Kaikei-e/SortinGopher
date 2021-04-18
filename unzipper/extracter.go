package unzipper

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
	"log"
)


func SortZipFile(thePath string) {
	zipExtracter(zipSearcher(thePath))

}

func zipSearcher(thePath string)([]string, string){
	files, _ := os.ReadDir(thePath)
	var fileList []string
	//dirPath := thePath

  for _, v := range files{
		if strings.Contains(v.Name(), ".zip") {
			fileList = append(fileList, thePath + v.Name())
		}
	}

	return fileList, thePath //, dirPath
}

func zipExtracter(filePaths []string, folder string) error {

	for _, fp := range filePaths{
		destPath := folder //+ "/extracted/"
		filepath.Clean(destPath)
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
				_, err := os.Create(path)
				if  err != nil {
					log.Fatal(err)
				}




				errWrite := os.WriteFile(path, buf, f.Mode())
				if errWrite != nil{
					return errWrite
				}
			}
		}
	}

	return nil
}
