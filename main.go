package main

import (
	"bytes"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/uniplaces/carbon"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"text/template"
	"time"
)

/**
 * Get day of week string slice
 */
func dayOfWeek(language string) []string {
	var s []string

	switch language {
	case "ja":
		s = []string{"日", "月", "火", "水", "木", "金", "土"}
	default:
		s = []string{"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"}
	}

	return s
}

/**
 * Create note from template as String.
 */
func create(start string, location string, language string) string {
	startDate, err := carbon.Parse(carbon.DateFormat, start, location)
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("./default.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	dayOfWeek := dayOfWeek(language)

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
	var language string
	var start string
	var location string
	var result string

	startOfWeek := carbon.Now().StartOfWeek().DateString()
	timeLocation := time.Now().Location().String()

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "print only the version",
	}

	app := &cli.App{
		Name:    "weeklynote",
		Version: "v1.0.1",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "start",
				Aliases:     []string{"s"},
				Value:       startOfWeek,
				DefaultText: startOfWeek,
				Usage:       "Specify the start date(YY-MM-DD). Default is the first day of the week of the current day.",
				Destination: &start,
			},

			&cli.StringFlag{
				Name:        "language",
				Aliases:     []string{"lang"},
				Value:       "en",
				DefaultText: "en",
				Usage:       "Specify the display language. \"en\" or \"ja\".",
				Destination: &language,
			},

			&cli.StringFlag{
				Name:        "location",
				Aliases:     []string{"loc"},
				Value:       "stdout",
				DefaultText: "stdout",
				Usage:       "Specify the output location. \"stdout\" or \"clipboard\"",
				Destination: &location,
			},
		},

		Action: func(c *cli.Context) error {
			result = create(start, timeLocation, language)
			output(result, location)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
