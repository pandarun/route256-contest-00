package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var __testOk = `6
R48FAO00OOO0OOA99OKA99OK
R48FAO00OOO0OOA99OKA99O
A9PQ
A9PQA
A99AAA99AAA99AAA99AA
AP9QA`

var __testOkResult = `R48FA O00OO O0OO A99OK A99OK
-
A9PQ
-
A99AA A99AA A99AA A99AA
-
`

func TestVerifyLicensePlate(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(__testOk))
	out := new(bytes.Buffer)
	err := verifyLicensePlate(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != __testOkResult {
		t.Errorf("test for OK failed - results not match\n %v \n----\n %v", x, __testOkResult)
	}
}
