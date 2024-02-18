package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
)

func parseNumber3(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}

type AirConditioner struct {
	TemperatureMin int
	TemperatureMax int
}

func checkAirconditioner(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	dataInputsAsStrings := in.Text()

	dataInputs, err := parseNumber3(dataInputsAsStrings)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs count")
	}
	var i, j int = 0, 0

	for i < dataInputs {
		in.Scan()
		employeesAsString := in.Text()

		employees, err := parseNumber3(employeesAsString)
		if err != nil {
			fmt.Println(employeesAsString)
			return fmt.Errorf("unable to parse employees count")
		}

		var ac = &AirConditioner{TemperatureMin: 15, TemperatureMax: 30}
		j = 0
		prev := 0
		for j < employees {

			in.Scan()

			if prev == -1 {
				j++
				fmt.Fprintln(output, prev)
				continue
			}

			txt := in.Text()

			temperatureParts := strings.Split(txt, " ")

			limit, err := parseNumber3(temperatureParts[1])
			if err != nil {
				return fmt.Errorf("unable to parse limit")
			}

			var result int

			switch temperatureParts[0] {
			case ">=":
				{
					if limit < ac.TemperatureMin {
						break
					}

					if limit >= ac.TemperatureMin && limit <= ac.TemperatureMax {
						ac.TemperatureMin = limit
					} else {
						result = -1
						goto END
					}
				}
			case "<=":
				{
					if limit > ac.TemperatureMax {
						break
					}

					if limit >= ac.TemperatureMin && limit <= ac.TemperatureMax {
						ac.TemperatureMax = limit
					} else {
						result = -1
						goto END
					}
				}
			}

			if (ac.TemperatureMax - ac.TemperatureMin) > 0 {
				result = rand.Intn(ac.TemperatureMax-ac.TemperatureMin) + ac.TemperatureMin
			} else {
				result = ac.TemperatureMin
			}

		END:
			//fmt.Fprint(output, "[", ac.TemperatureMin, " - ", ac.TemperatureMax, "]", " ")
			fmt.Fprintln(output, result)

			prev = result

			j++
		}

		var result string = ""

		fmt.Fprintln(output, result)

		i++
	}

	return nil
}
