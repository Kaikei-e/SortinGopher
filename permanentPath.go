package main

import (
	"log"
	"os"
	"time"
)

func permanentPath(){
	t := time.Now()

	const layout = "2006-01-02"
	fileName := t.Format(layout) + t.Weekday().String()

	f, err := os.Create("permanent" + fileName + ".txt")
	if err != nil{
		log.Fatal(err)
	}

	defer f.Close()

	
}


