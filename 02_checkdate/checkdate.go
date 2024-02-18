package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var maxDaysInMonth = map[int][]int{
	1:  []int{31},
	2:  []int{28, 29},
	3:  []int{31},
	4:  []int{30},
	5:  []int{31},
	6:  []int{30},
	7:  []int{31},
	8:  []int{31},
	9:  []int{30},
	10: []int{31},
	11: []int{30},
	12: []int{31},
}

func isLeapYear(year int) bool {
	return year%4 == 0 && !(year%100 == 0 && year%400 != 0)
}

func checkDay(day int, month int, isLeapYear bool) bool {
	if month == 2 {
		if !isLeapYear {
			return day > 0 && day <= maxDaysInMonth[month][0]
		} else {
			return day > 0 && day <= maxDaysInMonth[month][1]
		}
	} else {
		if day > 0 && day <= maxDaysInMonth[month][0] {
			return true
		} else {
			return false
		}
	}
}

func checkMonth(month int) bool {
	return month > 0 && month < 13
}

func checkYear(year int) bool {
	return year > 1949 && year < 2301
}

func parseNumber(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		return 0, fmt.Errorf("unable to parse date part")
	}
	return number, nil
}

func verifyDate(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	fieldsAsString := in.Text()

	datesCount, err := strconv.Atoi(fieldsAsString)
	if err != nil {
		return fmt.Errorf("unable to parse fields")
	}

	var i int = 0
	for i < datesCount {
		in.Scan()
		txt := in.Text()

		dateParts := strings.Split(txt, " ")

		var day, month, year int

		day, err = parseNumber(dateParts[0])
		if err != nil {
			return fmt.Errorf("unable to parse day")
		}

		month, err = parseNumber(dateParts[1])
		if err != nil {
			return fmt.Errorf("unable to parse month")
		}

		year, err = parseNumber(dateParts[2])
		if err != nil {
			return fmt.Errorf("unable to parse year")
		}

		var result string
		if checkDay(day, month, isLeapYear(year)) && checkMonth(month) && checkYear(year) {
			result = "YES"
		} else {
			result = "NO"
		}

		fmt.Fprintln(output, result)

		i++
	}

	return nil
}
