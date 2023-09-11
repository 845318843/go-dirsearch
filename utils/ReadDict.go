package utils

import (
	"bufio"
	"godir/common"
	"os"
	"strings"
)

var DictValue chan string = make(chan string, 100)

func ReadFromFile() {
	file, err := os.Open(common.WordList)
	if err != nil {
		common.Colerr.Printf("Read dict Fail!")
		os.Exit(1)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if strings.Contains(scan.Text(), "%EXT%") {
			for _, data := range common.ExtGroup {
				data = strings.ReplaceAll(scan.Text(), "%EXT%", data)
				DictValue <- data
			}
			continue
		}
		DictValue <- scan.Text()
	}
}

func ReadFromChan() string {
	return <-DictValue
}
