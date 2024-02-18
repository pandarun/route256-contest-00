package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var _testOk = `10
10 9 2022
21 9 2022
29 2 2022
31 2 2022
29 2 2000
29 2 2100
31 11 1999
31 12 1999
29 2 2024
29 2 2023`

var _testOkResult = `YES
YES
NO
NO
YES
NO
NO
YES
YES
NO`

var _testOk2 = `3
29 2 2000
31 12 1999
29 2 2024`

var _testOkResult2 = `YES
YES
YES`

func TestVerifyDate(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(_testOk))
	out := new(bytes.Buffer)
	err := verifyDate(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	x := out.String()
	if x != _testOkResult2 {
		t.Errorf("test for OK failed - results not match\n %v \n %v", x, _testOkResult)
	}
}

func TestLeapYear(t *testing.T) {
	if !isLeapYear(2000) {
		t.Errorf("test for leap year failed")
	}
	if isLeapYear(2100) {
		t.Errorf("test for leap year failed")
	}
	if !isLeapYear(2020) {
		t.Errorf("test for leap year failed")
	}
	if isLeapYear(2021) {
		t.Errorf("test for leap year failed")
	}
}
