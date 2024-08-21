package main

import (
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
)

type person struct {
	pName string
	pDate time.Time
	pMail string
}

func main() {

	var (
		newPerson person
	)

	today := time.Now().Format("01-02-06") // formatted date

	//opening file
	file, err := excelize.OpenFile("sourceTable/Table.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	//getting rows
	rows, err := file.GetRows("Лист1")
	if err != nil {
		log.Fatal(err)
	}

	//showing results
	for _, row := range rows[1:] {

		newPerson.pName = row[0]
		newPerson.pDate, err = time.Parse("01-02-06", row[1])
		if err != nil {
			log.Fatal(err)
		}
		newPerson.pMail = row[2]
		fmt.Println(newPerson)
	}
	fmt.Println(today)
	fmt.Println()
}
