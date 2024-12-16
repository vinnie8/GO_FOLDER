package main

import(
	"net/smtp"
)


// функция отправки сообщений
func sendEmail(persList []person, aSettings *settings) {

	preparedList := prepareSendList(persList)

	// Set up authentication information.
	auth := smtp.PlainAuth("", aSettings.SmtpUsername, aSettings.MailPassword, aSettings.SmtpServer)

	//sendingMessage
	err := smtp.SendMail(aSettings.SmtpServer+":"+aSettings.SmtpPort, auth, aSettings.SenderMail, preparedList,
		([]byte("Subject:" + aSettings.MailSubject + "\r\n" + aSettings.MailText + "\r\n")))

	if err != nil {
		panic(err)
	}
}