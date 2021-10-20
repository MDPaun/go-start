package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type product struct {
	ean      int
	price    string
	quantity int
}

func main() {

	csvfile, err := os.Open("ean.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		ean, _ := strconv.Atoi(record[0])
		crawl(ean)
		// fmt.Printf(record[0])
	}

	// for p := range products {
	// 	writer.Write([]string{
	// 		strconv.Itoa(products.product.ean), strconv.Itoa(int(p.price)), strconv.Itoa(p.quantity),
	// 	})
	// }

}

var products []product

func crawl(ean int) {

	fName := "stock.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"EAN", "Price", "Quantity"})

	c := colly.NewCollector(
		colly.AllowedDomains("www.stoklasa.ro"),
	)

	c.OnHTML("#dkz", func(e *colly.HTMLElement) {
		temp := product{}

		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		uean := fmt.Sprintf("input[data-ean='%d']", ean)
		code, _ := strconv.Atoi(re.FindString(e.ChildAttr(uean, "name")))
		// fmt.Println(code)

		e.ForEach(fmt.Sprintf("#dkz_volba_baleni > div.dkz_volba_baleni_spec.dkz_volba_baleni_spec_%d", code), func(_ int, e1 *colly.HTMLElement) {
			quantity, _ := strconv.Atoi(re.FindString(e1.ChildText("div:first-child > div:nth-child(2) > div:nth-child(5)")))
			temp.quantity = quantity

			price := strings.ReplaceAll(re.FindString(e1.ChildText("div:first-child > div:nth-child(2) > div:nth-child(3)")), ",", ".")
			// fmt.Println(price)
			// temp.price, _ = strconv.ParseFloat(price, 64)
			temp.price = price
			fmt.Println(temp.price)

		})
		temp.ean = ean
		fmt.Println(ean)
		products = append(products, temp)
	})

	url := fmt.Sprintf("https://www.stoklasa.ro/index.php?text=%d&skupina=h01", ean)
	// fmt.Println(url)
	c.Visit(url)

	// fmt.Println(products)

	for _, p := range products {
		writer.Write([]string{
			strconv.Itoa(p.ean), p.price, strconv.Itoa(p.quantity),
		})
		// writer.Flush()
	}

}
