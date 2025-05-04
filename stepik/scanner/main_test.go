package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"
)

var testOkData = `1
2
3
3
4
5`

var testOkResult = `1
2
3
4
5
`

func TestOk(t *testing.T) {
	input := bufio.NewReader(strings.NewReader(testOkData))
	output := new(bytes.Buffer)

	if err := ScanSortedExample(input, output); err != nil {
		t.Errorf("TestOk failed: error")
	}

	result := output.String()
	if result != testOkResult {
		t.Errorf("TestOk failed: result not matching\nResult: %v\nExpected: %v", testOkResult, result)
	}
}

var testErrorData = `1
2
1`

func TestError(t *testing.T) {
	input := bufio.NewReader(strings.NewReader(testErrorData))
	output := new(bytes.Buffer)

	if err := ScanSortedExample(input, output); err == nil {
		t.Errorf("TestError failed: error not raised")
	}
}
