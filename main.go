package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"text/template"
	"time"

	"github.com/atotto/clipboard"
	"github.com/uniplaces/carbon"
)

/**
 * Create note from template as String.
 */
func create(start string) string {
	startDate, err := carbon.Parse(carbon.DateFormat, start, time.Now().Location().String())
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

	return tpl.String()
}

/**
 * Output note to the location
 */
func output(content string, location string) {
	switch location {
	case "stdout":
		fmt.Println(content)
	case "clipboard":
		clipboard.WriteAll(content)
		fmt.Println("Sent to Clipboard!")
	}
}

func main() {
	var start, location, result string

	flag.StringVar(&start, "start", carbon.Now().StartOfWeek().DateString(), "Specify the start date(YY-MM-DD). Default is the first day of the week of the current day.")
	flag.StringVar(&location, "location", "stdout", "Specify the output location. `stdout` or `clipboard`.")
	flag.Parse()

	result = create(start)
	output(result, location)
}
