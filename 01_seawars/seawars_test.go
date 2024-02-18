package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOk = `5
2 1 3 1 2 3 1 1 4 2
1 1 1 2 2 2 3 3 3 4
1 1 1 1 2 2 2 3 3 4
4 3 3 2 2 2 1 1 1 1
4 4 4 4 4 4 4 4 4 4
`

var testOkResult = `YES
NO
YES
YES
NO
`

func TestVerify(t *testing.T) {
	in := bufio.NewReader(strings.NewReader(testOk))
	out := new(bytes.Buffer)
	err := verifyField(in, out)
	if err != nil {
		t.Errorf("test for Ok Failes")
	}
	if out.String() != testOkResult {
		t.Errorf("test for OK failed - results not match\n %v %v", out.String(), testOkResult)
	}
}
