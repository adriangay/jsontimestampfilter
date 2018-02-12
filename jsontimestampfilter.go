package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

const ISO8601 = "2006-01-02T15:04Z"

type Msg struct {
	Timestamp int64
}

func main() {

	args := os.Args[1]

	s := strings.Split(args, "/")
	start, end := s[0], s[1]

	startTime := convertDateToMillis(start)
	endTime := convertDateToMillis(end)

    lineNumber := 0
    startLine := ""
    endLine := ""

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("stdin is empty")
	} else {

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			line := scanner.Text()

			if (lineNumber == 0 ) {
				startLine = line
			}

			ts := getTimeStampFromLine(line)
			if (ts > startTime) && (ts < endTime) {
				fmt.Println(line)
			}

			endLine = line
			lineNumber++
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error")
		}

		fmt.Printf("Startline: %v\n", getTimeStampFromLine(startLine))
		fmt.Printf("Endline: %v\n", getTimeStampFromLine(endLine))
		fmt.Printf("Total Number of lines: %v\n", lineNumber)


	}

}

func convertDateToMillis(date string) int64 {
	var formattedTime time.Time
	formattedTime, _ = time.Parse(ISO8601, date)

	nanotime := formattedTime.UnixNano()
	return nanotime / 1000000
}

func getTimeStampFromLine(jsonInput string) int64 {
	var msg Msg
	json.Unmarshal([]byte(jsonInput), &msg)
	return msg.Timestamp
}
