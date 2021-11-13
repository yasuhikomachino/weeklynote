package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"text/template"

	"github.com/atotto/clipboard"
	"github.com/uniplaces/carbon"
)

func main() {
	flagStart := flag.String("start", "", "Specify the start date(YY-MM-DD). Default is the first day of the week of the current day.")
	flag.Parse()

	var start *carbon.Carbon

	if *flagStart == "" {
		start = carbon.Now().StartOfWeek()
	} else {
		_start, err := carbon.Parse(carbon.DateFormat, *flagStart, "Asia/Tokyo")
		if err != nil {
			log.Fatal(err)
		}

		start = _start
	}

	t, err := template.ParseFiles("./default.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	dayOfWeek := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	var days []string

	for i := 0; i < 7; i++ {
		d := start.AddDays(i)
		days = append(days, d.DateString()+"("+dayOfWeek[d.Weekday()]+")")
	}

	var tpl bytes.Buffer

	if err = t.Execute(&tpl, days); err != nil {
		log.Fatal(err)
	}

	// stdout
	fmt.Println(tpl.String())

	// clipboard
	clipboard.WriteAll(tpl.String())
}
