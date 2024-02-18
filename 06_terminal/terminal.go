package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

const (
	Left  = 'L'
	Right = 'R'
	End   = 'E'
	Home  = 'B'
	New   = 'N'
	Up    = 'U'
	Down  = 'D'
)

func parseNumber5(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}

func terminal(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	dataInputsAsStrings := in.Text()

	dataInputs, err := parseNumber5(dataInputsAsStrings)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs length")
	}

	var cursor = 0
	var currentRow = 0
	var currentTxt = ""
	var i = 0
	for i < dataInputs {

		in.Scan()

		txt := in.Text()

		for _, symbol := range txt {

			// fmt.Fprintln(output)
			// fmt.Fprintf(output, "cursor: %d, currentRow: %d, currentTxt: %s\n", cursor, currentRow, currentTxt)
			// fmt.Fprintln(output)

			switch symbol {
			case Left:
				{
					var rowCursor int
					if currentRow > 0 {
						rowCursorStart := strings.LastIndex(currentTxt[:cursor], "\n")
						rowCursor = cursor - rowCursorStart - 1
					} else {
						rowCursor = cursor
					}

					if rowCursor-1 < 0 {
						continue
					}

					cursor--

				}
			case Right:
				{
					var rowCursor int
					if currentRow > 0 {
						var rowCursorStart int = strings.LastIndex(currentTxt[:cursor], "\n")
						rowCursor = cursor - rowCursorStart - 1
					} else {
						rowCursor = cursor
					}

					rows := strings.Split(currentTxt, "\n")
					currentRowText := rows[currentRow]
					if (rowCursor + 1) > len(currentRowText) { // rowCursor + 1??
						continue
					}

					cursor++
				}
			case End:
				{
					rows := strings.Split(currentTxt, "\n")

					if currentRow == 0 {
						cursor = len(rows[currentRow])
					} else {
						cursor = sum(rows, currentRow, len(rows[currentRow]))
					}

				}
			case Home:
				{
					if currentRow == 0 {
						cursor = 0
					} else {
						rowCursorStart := strings.LastIndex(currentTxt[:cursor], "\n")
						cursor = rowCursorStart + 1
					}
				}
			case New:
				{
					currentRow++
					currentTxt = currentTxt[:cursor] + string('\n') + currentTxt[cursor:]
					cursor++
				}
			case Up:
				{
					if currentRow == 0 {
						continue
					}

					currentRow--

					rows := strings.Split(currentTxt, "\n")
					rowCursorStart := strings.LastIndex(currentTxt[:cursor], "\n")
					rowCursor := cursor - rowCursorStart - 1
					upRowCursor := min(rowCursor, len(rows[currentRow]))

					cursor = sum(rows, currentRow, upRowCursor)

				}
			case Down:
				{
					rows := strings.Split(currentTxt, "\n")

					if currentRow+1 > len(rows)-1 {
						continue
					}

					currentRow++

					rowCursorStart := strings.LastIndex(currentTxt[:cursor], "\n")
					rowCursor := cursor - rowCursorStart - 1
					downRowCursor := min(rowCursor, len(rows[currentRow]))

					cursor = sum(rows, currentRow, downRowCursor)

				}
			default:
				{
					currentTxt = currentTxt[:cursor] + string(symbol) + currentTxt[cursor:]
					cursor++
				}

			}
		}

		fmt.Fprintln(output, currentTxt)
		fmt.Fprintln(output, "-")

		cursor = 0
		currentRow = 0
		currentTxt = ""

		i++
	}

	return nil
}

func sum(rows []string, currentRow int, rowCursor int) int {
	var sum int = 0
	for i := 0; i < currentRow; i++ {
		sum += len(rows[i]) + 1
	}
	return sum + rowCursor
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
