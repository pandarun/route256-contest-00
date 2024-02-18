package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var ____testOk1 = `1
9
3 2 1 0 -1 0 6 6 7`

var _____testOkResult1 = `8
3 -4 0 0 6 0 6 1
`

func TestCompression1(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk1))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult1 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult1)
	}
}

var ____testOk2 = `5
9
3 2 1 0 -1 0 6 6 7
1
1000
7
1 2 3 4 5 6 7
7
1 3 5 7 9 11 13
11
100 101 102 103 19 20 21 22 42 41 40
`

var _____testOkResult2 = `8
3 -4 0 0 6 0 6 1
2
1000 0
2
1 6
14
1 0 3 0 5 0 7 0 9 0 11 0 13 0
6
100 3 19 3 42 -2
`

func TestCompression(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk2))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult2)
	}
}

var ____testOk5 = `1
1
1000
`

var _____testOkResult5 = `2
1000 0
`

func TestCompression3(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk5))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult5 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult5)
	}
}

var ____testOk6 = `1
7
1 2 3 4 5 6 7
`

var _____testOkResult6 = `2
1 6
`

func TestCompression6(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk6))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult6 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult6)
	}
}

var ____testOk7 = `1
7
1 3 5 7 9 11 13
`

var _____testOkResult7 = `14
1 0 3 0 5 0 7 0 9 0 11 0 13 0
`

func TestCompression7(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk7))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult7 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult7)
	}
}

var ____testOk8 = `1
11
100 101 102 103 19 20 21 22 42 41 40
`

var _____testOkResult8 = `6
100 3 19 3 42 -2
`

func TestCompression8(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk8))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult8 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult8)
	}
}


var ____testOk11 = `1
6
-1 -2 -3 -3 -2 -1
`

var _____testOkResult11 = `4
-1 -2 -3 2
`

func TestCompression11(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk11))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult11 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult11)
	}
}


var ____testOk9 = `1
4
1 2 2 1
`

var _____testOkResult9 = `4
1 1 2 -1
`

func TestCompression9(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(____testOk9))
	out := new(bytes.Buffer)
	err := compress(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}

	x := out.String()
	if x != _____testOkResult9 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _____testOkResult9)
	}
}