package mail

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"gitlab.com/SiberianPanda/selfcard_service/src/pkg/environment"
	gomail "gopkg.in/gomail.v2"
)

type callbackForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func (form *callbackForm) String() string {
	mailBody := "С вашего сайта пришло сообщение:<br/><b>Имя</b>: " + form.Name + "<br/>" + "<b>Email</b>: " + form.Email + "<br/>" + "<b>Message</b>: " + form.Message + "<br/>"
	return mailBody
}

// Send need for send mail to max.bez.l@mail.ru
func Send(formData *http.Request) error {
	var form callbackForm
	newDecoder := json.NewDecoder(formData.Body)

	if newDecoder.Decode(&form) != nil {
		msg := "Error in parse"
		log.Fatal(msg)
		return errors.New(msg)
	}
	// от кого
	log.Print("Start sending")
	defer log.Print("End sending")
	mail := gomail.NewMessage()
	mail.SetAddressHeader("From", environment.MailFrom, environment.MailFromName)
	mail.SetAddressHeader("To", environment.MailTo, environment.MailToName)
	mail.SetHeader("Subject", "Some one want call with you!")
	mail.SetBody("text/html", form.String())

	mailDial := gomail.NewPlainDialer(
		"smtp.yandex.ru",
		465,
		// for yandex smtp login (login + domain) -> test@test.com
		environment.MailLogin,
		environment.MailPass)

	if err := mailDial.DialAndSend(mail); err != nil {
		log.Fatal(err)
		// If you get catch the error, then try to confirm the phone
		return errors.New("Sending failed with: " + string(err.Error()))
	}
	return nil
}
