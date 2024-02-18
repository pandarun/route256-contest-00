package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var licensePlatchMatchRegex = `^(?:[A-Z]\d{1,2}[A-Z]{2})+$`
var licensePlateGroupsRegex = `([A-Z]\d{1,2}[A-Z]{2})`

func parseNumber2(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		return 0, fmt.Errorf("unable to parse date part")
	}
	return number, nil
}



func verifyLicensePlate(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	licensePlatesAsString := in.Text()

	licensePlates, err := parseNumber2(licensePlatesAsString)
	if err != nil {
		return fmt.Errorf("unable to parse licensePlates count")
	}
	var i int = 0
	for i < licensePlates {
		in.Scan()
		txt := in.Text()

		var result string
		match, _ := regexp.MatchString(licensePlatchMatchRegex, txt)
		if !match {
			result = "-"
		} else {

			re := regexp.MustCompile(licensePlateGroupsRegex)
			matches := re.FindAllStringSubmatch(txt, -1)

			var licensePlates []string = make([]string, 0, len(matches))
			for _, match := range matches {
				licensePlates = append(licensePlates, match[0])
			}

			result = strings.Join(licensePlates, " ")
		}

		fmt.Fprintln(output, result)

		i++
	}

	return nil
}
