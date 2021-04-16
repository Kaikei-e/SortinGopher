package permanentPath

import (
	"fmt"
	"log"
	"os"
	"time"
)

func MkPathFile(thePath string) {
	t := time.Now()

	const layout = "2006-01-02"
	fileName := t.Format(layout) + t.Weekday().String()

	f, err1 := os.Create("permanent" + fileName + ".txt")
	if err1 != nil{
		log.Fatal(err1)
	}

	defer f.Close()

	_, err2 := f.WriteString(thePath)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("done")

}


