package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	PrintPatternRegex = `(?P<range>\d-\d)|(?P<single>\d)`
)

func parseNumber6(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}

func printer(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	dataInputsAsStrings := in.Text()

	dataInputs, err := parseNumber6(dataInputsAsStrings)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs length")
	}

	//fmt.Fprintln(output, dataInputs)
	var i = 0
	for i < dataInputs {
		in.Scan()
		sequenceLengthAsString := in.Text()

		pagesLength, err := parseNumber6(sequenceLengthAsString)
		if err != nil {
			fmt.Println(sequenceLengthAsString)
			return fmt.Errorf("unable to parse pages length")
		}

		in.Scan()

		printedPagesPattern := in.Text()

		printedPages := strings.Split(printedPagesPattern, ",")

		var document []bool = make([]bool, pagesLength)

		for _, printedPage := range printedPages {
			if strings.Contains(printedPage, "-") {
				pagesRange := strings.Split(printedPage, "-")
				start, err := parseNumber6(pagesRange[0])
				if err != nil {
					return fmt.Errorf("unable to parse pages range start")
				}
				end, err := parseNumber6(pagesRange[1])
				if err != nil {
					return fmt.Errorf("unable to parse pages range end")
				}
				for i := start; i <= end; i++ {
					document[i-1] = true
				}
			} else {
				page, err := parseNumber6(printedPage)
				if err != nil {
					return fmt.Errorf("unable to parse page")
				}
				document[page-1] = true
			}
		}

		toPrint := []string{}

		var start int
		for index, printed := range document {
			if !printed && start == 0 && index != len(document)-1 {
				start = index + 1
			} else if printed && start != 0 {
				if start == index {
					toPrint = append(toPrint, strconv.Itoa(start))
				} else {
					toPrint = append(toPrint, strconv.Itoa(start)+"-"+strconv.Itoa(index))
				}
				start = 0
			} else if index == len(document)-1 && start != 0 {
				toPrint = append(toPrint, strconv.Itoa(start)+"-"+strconv.Itoa(index+1))
			} else if index == len(document)-1 && start == 0 && !printed {
				toPrint = append(toPrint, strconv.Itoa(index+1))
			}
		}

		fmt.Fprintln(output, strings.Join(toPrint, ","))

		i++
	}

	return nil
}
