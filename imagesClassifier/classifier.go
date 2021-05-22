package imagesClassifier

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
	//"regexp"
)

func FilesClassifier(dirName string, wg *sync.WaitGroup){
	defer wg.Done()

	imgClassifier(fileSearcher(dirName), dirName)
}

func fileSearcher(dirName string) []string {
	files, err := ioutil.ReadDir(dirName)
	if err != nil{
		panic(err)
	}


	var paths []string
	for _, file := range files{
		if file.IsDir(){
			paths = append(paths, fileSearcher(filepath.Join(dirName, file.Name()))...)

			continue
		}

		paths = append(paths, filepath.Join(dirName, file.Name()))
	}



	return paths
}

func imgClassifier(paths []string, folderPath string){
	filenameExtension := ".jpg"

	fmt.Println("Sorting was started.")

	sort.Slice(paths, func(i, j int) bool {return paths[i] < paths[j]})
	t := time.Now()
	const layout = "2006-01-02_15-04"
	tFormatted := string(t.Format(layout))
	//tTrimed := strings.Trim(tFormatted, '"')

	for i, p := range paths{
		imagePath := p//写真のパス



		if i == 0{
			continue
		}
		// get extend filename extension
		if strings.Contains(imagePath, ".jpg"){
			filenameExtension = ".jpg"
		}else if strings.Contains(imagePath, ".png"){
			filenameExtension = ".png"
		}else if strings.Contains(imagePath, ".mp4"){
			filenameExtension = ".mp4"
		}else if strings.Contains(imagePath, ".zip"){
			filenameExtension = ".zip"
		}else if strings.Contains(imagePath, ".csv"){
			filenameExtension = ".csv"
		}else if strings.Contains(imagePath, ".log"){
			filenameExtension = ".log"
		}

		var slice []string
		folderName := ""

		slice = strings.Split(imagePath, "")
		length := len(slice)

		for i := 0; i < length; i++{
			if strings.Contains(slice[i], "-"){
				break
			}
			folderName += slice[i]

		}

		fileName := folderName[len(folderPath) + 1:] //写真のパスからフォルダまでの文字列をトリム

		if strings.Contains(fileName, `\`){//フォルダの下の階層がある場合にスキップする処理
			continue
		}

		if strings.Contains(fileName, `/`){//フォルダの下の階層がある場合にスキップする処理
			continue
		}

		//fileName = fileName[0:8]//写真ファイルの名前、先頭9文字を取得

		/*
		 //2週目以降に通過
		imageFile, err := os.Open(imagePath)
		if err != nil{
			panic(err)
		}
		defer imageFile.Close()
		*/

		if strings.Contains(paths[i], fileName){
			if f, err := os.Stat(folderPath + "/" + fileName); os.IsNotExist(err) || !f.IsDir(){
				err = os.Mkdir(folderPath + "/" + fileName, 0755)
				if err != nil{
				panic(err)
				}

				if err := os.Rename(imagePath, folderPath + "/" + fileName + "/" + fileName + "-"+ tFormatted + "No" + strconv.Itoa(i) + filenameExtension)
				err != nil{
					panic(err)
				}
			}else{
				if err := os.Rename(imagePath, folderPath + "/" + fileName + "/" + fileName + "-" + tFormatted + "No" + strconv.Itoa(i) + filenameExtension)
				err != nil{
					panic(err)
				}
			}

		}


	}


	fmt.Println("Files were sorted.")
}

