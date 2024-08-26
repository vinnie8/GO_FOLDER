package main

import (
	"time"
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
	SmtpServer    string
	SmtpPort      string
	SmtpUsername  string
	MailSubject   string
	MailText      string
}

func main() {

	var (
		actualSettings settings
		personList     []person
	)

	//открываем табличный список Table.xlsx
	rows := openExcelFile()

	//читаем настройки из файла ini и записываем в actualSettings
	readFileAndUnmarshal("settings/settings.ini", &actualSettings)

	//составляем список людей
	personList = createPersonList(rows, &actualSettings)

	//отправка сообщений
	sendEmail(personList, &actualSettings)

}
