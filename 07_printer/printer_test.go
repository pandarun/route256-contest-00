package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var ________testOk1 = `7
8
7
8
1,7,1
8
1-5,1,7-7
10
1-5
10
1,2,3,4,5,6,8,9,10
3
1-2
100
1-2,3-7,10-20,100`

var ________testOkResult1 = `1-6,8
2-6,8
6,8
6-10
7
3
8-9,21-99
`

func TestPrint1(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(________testOk1))
	out := new(bytes.Buffer)
	err := printer(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ________testOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ________testOkResult1)
	}
}

var ________testOk2 = `1
8
7`

var ________testOkResult2 = `1-6,8
`

func TestPrint2(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(________testOk2))
	out := new(bytes.Buffer)
	err := printer(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ________testOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ________testOkResult2)
	}
}

var ________testOk3 = `1
8
1,7,1`

var ________testOkResult3 = `2-6,8
`

func TestPrint3(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(________testOk3))
	out := new(bytes.Buffer)
	err := printer(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ________testOkResult3 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ________testOkResult3)
	}
}

var ________testOk4 = `1
8
1-5,1,7-7`

var ________testOkResult4 = `6,8
`

func TestPrint4(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(________testOk4))
	out := new(bytes.Buffer)
	err := printer(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ________testOkResult4 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ________testOkResult4)
	}
}

var ________testOk5 = `1
10
1-5`

var ________testOkResult5 = `6-10
`

func TestPrint5(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(________testOk5))
	out := new(bytes.Buffer)
	err := printer(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ________testOkResult5 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ________testOkResult5)
	}
}

var ________testOk6 = `1
10
1,2,3,4,5,6,8,9,10`

var ________testOkResult6 = `7
`

func TestPrint6(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(________testOk6))
	out := new(bytes.Buffer)
	err := printer(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != ________testOkResult6 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, ________testOkResult6)
	}
}
