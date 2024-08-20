package main

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func main() {

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
	for _, row := range rows {
		for _, col := range row {
			fmt.Print(col, "\t")
		}
		fmt.Println()
	}

}
