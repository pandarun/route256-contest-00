package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var frameBorderRegex = `\*{3,}`

type Coords struct {
	x int
	y int
}

type Range struct {
	start int
	end   int
}

type Frame struct {
	leftTop     Coords
	rightBottom Coords
	level       int
	id          int
}

func parseNumber8(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}

type FramePrototype struct {
	y     int
	level int
	id    int
}

func frame(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	dataInputsAsStrings := in.Text()

	dataInputs, err := parseNumber8(dataInputsAsStrings)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs length")
	}

	frameRegex := regexp.MustCompile(frameBorderRegex)

	var dataInput = 0
	for dataInput < dataInputs {

		in.Scan()
		fieldDimensions := in.Text()

		dimensions := strings.Split(fieldDimensions, " ")

		n, err := parseNumber8(dimensions[0])
		if err != nil {
			return fmt.Errorf("unable to parse n length")
		}

		m, err := parseNumber8(dimensions[1])
		if err != nil {
			return fmt.Errorf("unable to parse m length")
		}

		field := make([]string, n)

		for i := 0; i < n; i++ {

			in.Scan()
			field[i] = in.Text()
		}

		framePrototypes := map[Range]FramePrototype{}

		frames := []Frame{}

		id := 1
		ranges := make([]int, m+1, m+1)
		rangeByFrameIds := map[int]Range{}
		for y := 0; y < n; y++ {

			matches := frameRegex.FindAllStringSubmatchIndex(field[y], -1)
			if len(matches) == 0 {
				continue
			}
			for _, match := range matches {
				frameX1X2 := Range{start: match[0], end: match[1]}
				if framePrototype, ok := framePrototypes[frameX1X2]; !ok {

					if parentRange, ok := ExistsParentFor(frameX1X2, framePrototypes, ranges, rangeByFrameIds); ok {
						framePrototypes[frameX1X2] = FramePrototype{y: y, level: framePrototypes[parentRange].level + 1, id: id}

					} else {
						framePrototypes[frameX1X2] = FramePrototype{y: y, level: 0, id: id}
					}

					ranges[frameX1X2.start] = id
					ranges[frameX1X2.end] = id

					rangeByFrameIds[id] = frameX1X2

					id++

				} else {
					f := Frame{
						leftTop:     Coords{x: frameX1X2.start, y: framePrototype.y},
						rightBottom: Coords{x: frameX1X2.end, y: y},
						level:       framePrototype.level,
						id:          framePrototype.id,
					}
					frames = append(frames, f)
					delete(framePrototypes, frameX1X2)

					ranges[frameX1X2.start] = 0
					ranges[frameX1X2.end] = 0

					delete(rangeByFrameIds, framePrototype.id)

				}
			}
		}

		result := make([]int, 0, len(frames))

		for _, frame := range frames {
			result = append(result, frame.level)
		}

		sort.Slice(result, func(first, second int) bool {
			return result[first] < result[second]
		})

		var buffer bytes.Buffer

		for _, item := range result {

			buffer.WriteString(strconv.Itoa(item))
			buffer.WriteString(" ")
		}
		fmt.Fprintln(output, buffer.String())

		dataInput++
	}

	return nil
}

func ExistsParentFor(frameX1X2 Range, framePrototypes map[Range]FramePrototype, ranges []int, rangeByFrameId map[int]Range) (Range, bool) {

	coveringRanges := make([]Range, 0, len(rangeByFrameId))
	for _, currentRange := range rangeByFrameId {

		left := currentRange.start
		right := currentRange.end

		if frameX1X2.start > left && frameX1X2.end < right {
			coveringRanges = append(coveringRanges, Range{start: left, end: right})
		}

	}

	if len(coveringRanges) > 0 {

		sort.Slice(coveringRanges, func(first, second int) bool {
			return Abs1(coveringRanges[first].start-coveringRanges[first].end) < Abs1(coveringRanges[second].start-coveringRanges[second].end)
		})

		return coveringRanges[0], true
	}

	return Range{}, false
}

func Abs1(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
