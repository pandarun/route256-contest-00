package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func verifyField(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	fieldsAsString := in.Text()

	fieldsCount, err := strconv.Atoi(fieldsAsString)
	if err != nil {
		return fmt.Errorf("unable to parse fields")
	}

	var i int = 0
	for i < fieldsCount {
		in.Scan()
		txt := in.Text()

		ships := strings.Split(txt, " ")

		field := map[int]int{}
		for _, item := range ships {
			ship, err := strconv.Atoi(item)
			if err != nil {
				return fmt.Errorf("unable to parse ship")
			}

			field[ship]++
		}

		var result string
		if field[1] == 4 && field[2] == 3 && field[3] == 2 && field[4] == 1 {
			result = "YES"
		} else {
			result = "NO"
		}

		fmt.Fprintln(output, result)

		i++
	}

	return nil
}
