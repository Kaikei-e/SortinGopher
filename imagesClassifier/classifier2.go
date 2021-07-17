package imagesClassifier

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"sync"
	"time"
)


func ImgClassifier(dirName string, wg *sync.WaitGroup) {
	defer wg.Done()

	t := time.Now()
	const layout = "2006-01-02_15-04"
	tFormatted := string(t.Format(layout))

	log.Println("Sorting was started.")


	files, err := ioutil.ReadDir(dirName)
	if err != nil {
		log.Fatalln(err)
	}

	for i, f := range files{
		if f.IsDir(){
			continue
		}
		log.Println(f.Name())
		log.Println(i)

		fileName := filepath.Base(dirName + "/" + f.Name())


		//fName := f.Name()

		replace := regexp.MustCompile(`\-.+?\-`)
		fileName = replace.ReplaceAllString(fileName, "-")

		repID := regexp.MustCompile(`\s*-\s*`)
    result := repID.Split(fileName, -1)

		accountID := result[0]

		extension := filepath.Ext(f.Name())

		rep := regexp.MustCompile(extension)
    fileName = filepath.Base(rep.ReplaceAllString(fileName, ""))

		if _, err := os.Stat(dirName + "/" + accountID); os.IsNotExist(err) {
			err = os.Mkdir(dirName + "/" + accountID, 0755)
				if err != nil{
				log.Fatalln(err)
				}

			if err := os.Rename(dirName + "/" + f.Name(), dirName + "/" + accountID + "-" + tFormatted + "-" + "No" + strconv.Itoa(i) + extension)
				err != nil{
					panic(err)
			}

			if err := os.Rename(dirName + "/" + accountID + "-" + tFormatted + "-" + "No" + strconv.Itoa(i) + extension, dirName + "/" + accountID + "/" + accountID + "-" + tFormatted + "-" + "No" + strconv.Itoa(i) + extension)
				err != nil{
					panic(err)
			}
			// path does not exist
		}

		if _, err := os.Stat(dirName + "/" + accountID); !os.IsNotExist(err) {
			if err := os.Rename(dirName + "/" + f.Name(), dirName + "/" + accountID + "-" + tFormatted + "-" + "No" + strconv.Itoa(i) + extension)
				err != nil{
					panic(err)
			}

			if err := os.Rename(dirName + "/" + accountID + "-" + tFormatted + "-" + "No" + strconv.Itoa(i) + extension, dirName + "/" + accountID + "/" + accountID + "-" + tFormatted + "-" + "No" + strconv.Itoa(i) + extension)
				err != nil{
					panic(err)
			}


			// path exists


		}






	}

	log.Println("Files were sorted.")


}
