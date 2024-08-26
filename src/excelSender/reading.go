package main

import(
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
)

// чтение из файла эксель
func openExcelFile() (rows [][]string) {
	file, err := excelize.OpenFile("sourceTable/Table.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	//загрузка ячеек из файла
	rows, err = file.GetRows("Лист1")
	if err != nil {
		log.Fatal(err)
	}
	return
}


// функция чтения из файла настроек
func readFileAndUnmarshal(path string, actualSettings *settings) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	//преобразование настроек из json формата
	err = json.Unmarshal(data, &actualSettings)
	if err != nil {
		log.Fatal(err)
	}
}