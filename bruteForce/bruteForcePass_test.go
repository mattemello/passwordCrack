package bruteforce

import (
	"fmt"
	"testing"
)

func TestCsvAttack(t *testing.T) {
	type testCase struct {
		password string
		path     string
		expected bool
	}

	pass := []testCase{
		{password: "test", path: "../dataset/test.csv", expected: true},
		{password: "test2", path: "../dataset/test.csv", expected: false},
		{password: "123qwe", path: "../dataset/test.csv", expected: true},
	}

	fmt.Println("Start of the search")

	for _, test := range pass {

		ris := searchFromCsv(test.password, test.path)

		if ris.Founded == test.expected {
			fmt.Println("[] -> Pass ", ris.TimeTotal)
		} else {
			t.Errorf("Fail %s expected: %t; result: %t \n", test.password, test.expected, ris.Founded)
		}

	}

}
