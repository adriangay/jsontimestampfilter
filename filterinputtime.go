package main
import (
	"fmt"
	"time"
	"encoding/json"
	"os"
	"strings"
	"bufio"
	)

const 	ISO8601 = "2006-01-02T15:04Z"

type Msg struct {
    Timestamp int64
}

func main() {

args := os.Args[1]

s := strings.Split(args, "/")
    start, end := s[0], s[1]

    startTime := convertDateToMillis(start)
    endTime := convertDateToMillis(end)

 fi, err := os.Stdin.Stat()
  if err != nil {
    panic(err)
  }

  if fi.Mode() & os.ModeNamedPipe == 0 {
	    fmt.Println("stdin is empty")
  } else {

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
	line := scanner.Text()
	ts := getTimeStampFromLine(line)
	if ( ts > startTime ) && ( ts < endTime ) {
        	fmt.Println(line)
	}
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error")
    }

  } 

}

func convertDateToMillis(date string) int64 {
   var formattedTime  time.Time
   formattedTime, _ = time.Parse(ISO8601, date)

   nanotime := formattedTime.UnixNano()
   return nanotime / 1000000
}

func getTimeStampFromLine(jsonInput string) int64 {
    var msg Msg
    json.Unmarshal([]byte(jsonInput), &msg)
    return msg.Timestamp
}

