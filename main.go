package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	"github.com/atotto/clipboard"
	"github.com/uniplaces/carbon"
)

func main() {
	t, err := template.ParseFiles("./default.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	dayOfWeek := []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	start := carbon.Now().StartOfWeek()

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
