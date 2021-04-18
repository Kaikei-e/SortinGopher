package unzipper

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	//"strings"
)


func SortZipFile(thePath string, wg *sync.WaitGroup) {
	defer wg.Done()

	zipExtracter(zipSearcher(thePath))

}

func zipSearcher(thePath string)([]string, string){
	var zipFiles []string
  err := filepath.Walk(thePath, func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    if info.IsDir() {
        return nil
    }
    if matched, err := filepath.Match("*.zip", filepath.Base(path)); err != nil {
        return err
    } else if matched {
        zipFiles = append(zipFiles, path)
    }
    return nil
  })
  if err != nil {
      return nil, err.Error()
  }

	fmt.Println("Path is " + thePath)

	/*
	files, err := os.ReadDir(thePath)
	if err != nil {
		log.Fatal(err)
	}

	*/
	fmt.Println("Got the file list. Length is ")
	fmt.Sprintln(len(zipFiles))

	//var fileList []string
	//dirPath := thePath

	/*
  for _, v := range zipFiles{
		if strings.Contains(v.Name(), ".zip") {
			fileList = append(fileList, thePath + v.Name())
		}
	}
	*/
	fmt.Println("Got the zip file list.")

	return zipFiles, thePath //, dirPath
}

func zipExtracter(filePaths []string, folder string) error {

	fmt.Println("Extracting is started")
	for _, fp := range filePaths{


		destPath := folder
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
				out, err := os.Create(path)
				if  err != nil {
					log.Fatal(err)
				}
				defer out.Close()




				errWrite := os.WriteFile(path, buf, f.Mode())
				if errWrite != nil{
					return errWrite
				}
			}
		}
	}

	fmt.Println("Extracted the zip.")

	return nil
}
