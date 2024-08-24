package main

import (
	"encoding/json"
	"log"
	"net/smtp"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

type person struct {
	pName       string
	pDate       time.Time
	pMail       string
	pMailSended bool
}

type settings struct {
	ToleranceDays string
	SenderMail    string
	MailPassword  string
	MailSubject   string
	MailText      string
	SmtpServer    string
	SmtpPort      string
	SmtpUsername  string
}

func main() {

	var (
		newPerson      person
		actualSettings settings
		personList     []person
	)

	//----------------------------------------------------------
	//открываем табличный список Table.xlsx
	file, err := excelize.OpenFile("sourceTable/Table.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	//--getting rows
	rows, err := file.GetRows("Лист1")
	if err != nil {
		log.Fatal(err)
	}
	//---------------------------------------------------------
	//читаем настройки из файла ini
	settings := readFile("settings/settings.ini")

	//преобразование настроек из json формата
	err1 := json.Unmarshal(settings, &actualSettings)
	if err1 != nil {
		log.Fatal(err)
	}

	//--создаем срез для списка людей
	personList = make([]person, len(rows)-1)

	//--счетчик для списка рассылки
	i := 0

	//--преобразуем данные в целочисленный тип для сравнения
	tolerance, _ := strconv.Atoi(actualSettings.ToleranceDays)

	//--------------------------------------------------------
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

	//fmt.Println(personList)
	//fmt.Println()

	preparedList := prepareSendList(personList)
	//fmt.Println(preparedList)
	//--отправка сообщений
	sendEmail(preparedList, personList, &actualSettings)
}

// ----------------
// функция чтения файла
func readFile(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// ---------------------

// функция отправки сообщений
func sendEmail(prepList []string, persList []person, aSettings *settings) {

	// Set up authentication information.
	auth := smtp.PlainAuth("", aSettings.SmtpUsername, aSettings.MailPassword, aSettings.SmtpServer)

	//sendingMessage
	err := smtp.SendMail(aSettings.SmtpServer+":"+aSettings.SmtpPort, auth, aSettings.SenderMail, prepList,
		([]byte("Subject:" + aSettings.MailSubject + "\r\n" + aSettings.MailText + "\r\n")))
	if err != nil {
		panic(err)
	}
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
