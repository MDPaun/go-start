package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/mail"
	"os"
	"text/template"

	"gopkg.in/gomail.v2"
)

// TODO
//! https://github.com/emersion/go-msgauth - setat DKIM pentru email
//* alceva

//? https://gist.github.com/ivanmrchk/e30eb45808536159bbec9aac20058b78 - cod
func sendMail(i string) {
	// lista := []string{i.Email}

	from := "comenzi@decocraft.ro"
	sub := "Discount - decocraft.ro"

	t := template.New("discount-20.html")

	var err error
	t, err = t.ParseFiles("discount-20.html")
	if err != nil {
		log.Println(err)
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, "i.Name"); err != nil { // nu folosim i.name https://gist.github.com/ivanmrchk/e30eb45808536159bbec9aac20058b78
		log.Println(err)
	}

	result := tpl.String()

	d := gomail.NewDialer("mail.decocraft.ro", 587, "comenzi@decocraft.ro", "mammut12")
	s, err := d.Dial()

	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetAddressHeader("To", i, "DecoCraft")
	m.SetHeader("Subject", sub)
	m.SetBody("text/html", result)
	// m.Attach("template.html") // attach factura/logo etc

	if err := gomail.Send(s, m); err != nil {
		log.Printf("Could not send email to %q: %v", i, err)
	} else {
		fmt.Println("Trimis: ", i)
	}

	m.Reset()
}

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func main() {

	emails, err := os.Open("lista-emails-20.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer emails.Close()

	r := csv.NewReader(emails)

	for {
		email, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// x := email[0]

		if valid(email[0]) {
			sendMail(email[0])
		} else {
			fmt.Printf("this %v is false\n", email[0])
		}

	}

}
