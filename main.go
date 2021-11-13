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
	var start, output string

	flag.StringVar(&start, "start", "", "Specify the start date(YY-MM-DD). Default is the first day of the week of the current day.")
	flag.StringVar(&output, "output", "stdout", "Specify the output location. Default is `stdout`.")
	flag.Parse()

	if start == "" {
		start = carbon.Now().StartOfWeek().DateString()
	}

	startDate, err := carbon.Parse(carbon.DateFormat, start, "Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("./default.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	dayOfWeek := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}

	var days []string

	for i := 0; i < 7; i++ {
		d := startDate.AddDays(i)
		days = append(days, d.DateString()+"("+dayOfWeek[d.Weekday()]+")")
	}

	var tpl bytes.Buffer

	if err = t.Execute(&tpl, days); err != nil {
		log.Fatal(err)
	}

	switch output {
	case "stdout":
		fmt.Println(tpl.String())
	case "clipboard":
		clipboard.WriteAll(tpl.String())
		fmt.Println("Sent to Clipboard!")
	}
}
