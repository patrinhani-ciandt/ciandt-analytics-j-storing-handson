package csvprocessor

import (
	"os"
	"fmt"
	"encoding/csv"
	"io"
)

type ProcessCSVLine func([]string)

func ProcessCSVFile(filePath string, delimiter rune, processHeader ProcessCSVLine, processData ProcessCSVLine) {
	
	file, err := os.Open(filePath)
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return
	}
	// automatically call Close() at the end of current method
	defer file.Close()

	reader := csv.NewReader(file)
	// options are available at:
	// http://golang.org/src/pkg/encoding/csv/reader.go?s=3213:3671#L94
	reader.Comma = delimiter
	lineCount := 0
		
	for {
		// read just one record, but we could ReadAll() as well
		record, err := reader.Read()
		// end-of-file is fitted into err
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		
		if (lineCount == 0) {
			
			processHeader(record)
		} else {
			
			processData(record)
		}
		
		lineCount += 1
	}
}
