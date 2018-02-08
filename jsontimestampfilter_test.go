package main

import "testing"

const (
	TEST_DATE  = "2006-01-02T15:04Z"
	TEST_EPOCH = 1136214240000
	JSON_LINE  = `{"timestamp":1511290852857,"animal":"rabbit","car":"mustang","fruit":"apple" }`
	JSON_EPOCH = 1511290852857
)

func TestConvertDateToMillis(t *testing.T) {
	time := convertDateToMillis(TEST_DATE)

	if time != TEST_EPOCH {
		t.Errorf("Epoch Time was wrong, got: %d, want: %d.", time, TEST_EPOCH)
	}
}

func TestGetTimeStampFromLine(t *testing.T) {
	timestamp := getTimeStampFromLine(JSON_LINE)

	if timestamp != JSON_EPOCH {
		t.Errorf("Epoch Time was wrong, got: %d, want: %d.", timestamp, JSON_EPOCH)
	}
}
