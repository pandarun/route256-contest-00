package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// func main() {
// 	err := compress(os.Stdin, os.Stdout)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

func parseNumber4(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}

func compress(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	dataInputsAsStrings := in.Text()

	dataInputs, err := parseNumber4(dataInputsAsStrings)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs length")
	}

	//fmt.Fprintln(output, dataInputs)
	var i = 0
	for i < dataInputs {
		in.Scan()
		sequenceLengthAsString := in.Text()

		_, err := parseNumber4(sequenceLengthAsString)
		if err != nil {
			fmt.Println(sequenceLengthAsString)
			return fmt.Errorf("unable to parse sequence length")
		}

		in.Scan()

		txt := in.Text()

		sequenceParts := strings.Split(txt, " ")

		var result []string

		var start int
		var growingIterator int
		var decreasingIterator int

		var currentNumber int
		var prevNumber int

		if (len(sequenceParts)) == 1 {
			result = append(result, sequenceParts[0])
			result = append(result, strconv.Itoa(0))
		} else {
			for index, itemAsString := range sequenceParts {

				if index == 0 {
					prevNumber, err = parseNumber4(itemAsString)
					if err != nil {
						return fmt.Errorf("unable to parse current number")
					}
					continue
				}

				currentNumber, err = parseNumber4(itemAsString)
				if err != nil {
					return fmt.Errorf("unable to parse current number")
				}

				if prevNumber < currentNumber && Abs(currentNumber-prevNumber) == 1 {

					if decreasingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(-decreasingIterator))
						decreasingIterator = 0
						start = index

					} else {
						growingIterator++
					}

				} else if prevNumber > currentNumber && Abs(currentNumber-prevNumber) == 1 {

					if growingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(growingIterator))
						growingIterator = 0
						start = index
					} else {
						decreasingIterator++
					}

				} else if growingIterator == 0 && decreasingIterator == 0 && Abs(currentNumber-prevNumber) > 1 {

					result = append(result, sequenceParts[start])
					result = append(result, strconv.Itoa(0))
					start = index
				} else if growingIterator == 0 && decreasingIterator == 0 && currentNumber == prevNumber {

					result = append(result, sequenceParts[start])
					result = append(result, strconv.Itoa(0))
					start = index
				} else if currentNumber == prevNumber {
					if growingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(growingIterator))
						growingIterator = 0
					} else if decreasingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(-decreasingIterator))
						decreasingIterator = 0
					}

					// result = append(result, sequenceParts[start])
					// result = append(result, strconv.Itoa(0))
					start = index
				} else {

					if growingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(growingIterator))
						growingIterator = 0
					} else if decreasingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(-decreasingIterator))
						decreasingIterator = 0
					}

					start = index

				}

				if index == len(sequenceParts)-1 {
					if growingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(growingIterator))
						growingIterator = 0
					} else if decreasingIterator > 0 {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(-decreasingIterator))
						decreasingIterator = 0
					} else {
						result = append(result, sequenceParts[start])
						result = append(result, strconv.Itoa(0))
					}
				}

				prevNumber = currentNumber
			}
		}

		fmt.Fprintln(output, len(result))
		fmt.Fprintln(output, strings.Join(result, " "))

		i++
	}

	return nil
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
