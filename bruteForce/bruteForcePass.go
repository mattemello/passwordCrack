package bruteforce

import (
	"encoding/csv"
	"time"
	// "fmt"
	"io"
	"os"
)

type testingPass struct {
	TimeTotal time.Duration
	Founded   bool
}

func BruteForcePass(pass string, timePass chan time.Duration) {
	csvResult := searchFromCsv(pass, "./dataset/common_passwords.csv")

	if csvResult.Founded {
		timePass <- csvResult.TimeTotal
		return
	}

	timePass <- csvResult.TimeTotal
	return

	//todo: second bruteForce
}

func searchFromCsv(pass, path string) testingPass {
	var csvSearch testingPass

	timeStart := time.Now()

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		read, err := reader.Read()
		if err == io.EOF {
			break
		}

		if read[0] == pass {
			csvSearch.Founded = true
			csvSearch.TimeTotal = time.Now().Sub(timeStart)

			return csvSearch
		}
	}

	csvSearch.TimeTotal = time.Now().Sub(timeStart)
	csvSearch.Founded = false

	return csvSearch
}
