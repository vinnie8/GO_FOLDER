package main

import(
	
	"log"
	"strconv"
	"sort"
	"time"
)
//создаем список лиц для рассылки
func createPersonList(rows [][]string, actualSettings *settings) (personList []person) {
	var (
		newPerson person
		i         int
		err 	  error
	)
	personList = make([]person, len(rows)-1)

	//преобразуем допуск в целочисленный тип для сравнения с датой
	tolerance, _ := strconv.Atoi(actualSettings.ToleranceDays)

	//читаем данные из массива ячеек
	for _, row := range rows[1:] { //--начинаем с 1 чтобы пропустить заголовок таблицы

		//--преобразуем и вносим данные из массива в структуру (дата)
		date, _ := time.Parse("01-02-06", row[1])

		//--расчет количества дней до наступления даты
		daysUntil := int(time.Until(date).Hours()) / 24

		//--считываем имя пользователя
		newPerson.pName = row[0]

		//--считываем требуемую дату, указанную в файле
		newPerson.pDate, err = time.Parse("01-02-06", row[1])
		if err != nil {
			log.Fatal(err)
		}

		//--считываем, отправлено ли было письмо ранее
		newPerson.pMail = row[2]
		if row[3] == "Нет" {
			newPerson.pMailSended = false
		} else {
			newPerson.pMailSended = true
		}

		//создание списка рассылки
		if daysUntil <= tolerance && !newPerson.pMailSended {
			personList[i] = newPerson //--вносим подходящих людей в список рассылки
			i++
		}
	}
	return
}

// функция сортировки и подготовки списка
func prepareSendList(list []person) []string {

	sortedList := make([]string, len(list))
	//формирование массива
	for key, val := range list {
		if val.pMail != "" {
			sortedList[key] = val.pMail
		}

	}

	//сортировка массива
	sort.Strings(sortedList)
	//очистка от дубликатов
	for i := 0; i < len(sortedList)-1; i++ {
		if sortedList[i] == sortedList[i+1] {
			sortedList = append(sortedList[0:i], sortedList[i+1:]...)
			i-- //откатываемся на шаг назад чтобы начать со свежей записи
		}

	}
	return sortedList
}