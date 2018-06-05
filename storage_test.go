package schedule

import (
	"os"
	"testing"
)

func TestSavedShift(t *testing.T) {
	testRecord := DayRecord{6, 6, 2006}
	testRecord.ClockIn(Employee{"Tester", 666}, 0, 48)

	if _, err := os.Stat(testRecord.String()); os.IsNotExist(err) {
		t.Error("Test record was not checked in, thus not saved")
	}

	defer os.RemoveAll(testRecord.String())

	if Construct(testRecord) == nil {
		t.Error("Test record was not checked in, thus not saved")
	}
}
