package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("valid_email_deco.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	oklist := []string{}

	for _, record := range records {
		email := record[0]
		err = ValidateFormat(email)
		if err != nil {
			continue
		}
		var (
			serverHostName    = "smtp.gmail.com"      // set your SMTP server here
			serverMailAddress = "pauncraft@gmail.com" // set your valid mail address here
		)
		err = ValidateHostAndUser(serverHostName, serverMailAddress, email)

		if smtpErr, ok := err.(SmtpError); ok && err != nil {
			fmt.Printf("Code: %s, Msg: %s", smtpErr.Code(), smtpErr)
			continue
		}
		oklist = append(oklist, email)

	}

	listaok, err := os.Create("oklist.csv")
	checkError("Cannot create file", err)
	defer listaok.Close()

	writer := csv.NewWriter(listaok)

	defer writer.Flush()

	for _, value := range oklist {
		record := []string{value}
		err := writer.Write(record)
		checkError("Cannot write to file", err)
	}

}
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
