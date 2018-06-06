/*
 * Shift Scheduler
 * Copyright (C) 2018  Ritchie Borja
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License along
 * with this program; if not, write to the Free Software Foundation, Inc.,
 * 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
 */

package schedule

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

// DayRecord is a struct that holds the Month, Day, and Year
// to form the basis of the name of the directory for later
// storage of all shifts. This struct contains methods that
// let's the user save the new employee's shift to the day
// shift in the DayRecord directory
type DayRecord struct {
	Month int
	Day   int
	Year  int
}

// NoIOPermissionsPanic is a checker if an error is returned
// by the File Handling library
func NoIOPermissionsPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckNumber is a checker if an error is returned by
// the string conversion library
func CheckNumber(num int, err error) int {
	if err != nil {
		panic(err)
	}

	return num
}

// String is the string representation of the DayRecord struct
func (dayRecord DayRecord) String() string {
	return fmt.Sprintf("%d.%d.%d", dayRecord.Month, dayRecord.Day, dayRecord.Year)
}

// ClockIn clocks in the employee with the start and end time interval
// to the current shift day. This makes use of the scheduler's overlapping
// logic and the operations of the linked list used by the scheduler.
// When nothing overlaps, the employee shift information will be stored
// in a text file inside the directory that is named by the date of the
// shift in question.
func (dayRecord DayRecord) ClockIn(worker Employee, start, end int) {
	if _, err := os.Stat(dayRecord.String()); os.IsNotExist(err) {
		os.Mkdir(dayRecord.String(), 0700)
	}

	files, err := ioutil.ReadDir(dayRecord.String())

	NoIOPermissionsPanic(err)

	counter := len(files) + 1

	workerFile, err3 := os.Create(path.Join(dayRecord.String(), fmt.Sprintf("%d.txt", counter)))
	defer workerFile.Close()

	NoIOPermissionsPanic(err3)

	fmt.Fprintf(workerFile, "%s:%d:%d:%d", worker.Name, worker.Id, start, end)
}

// Construct is the construction of the linked list of all employees' shift
// from a certain shift day. This retrieves all text files in a certain
// directory named after the shift day in m.d.YYYY format
func Construct(record DayRecord) *Shift {
	if _, err := os.Stat(record.String()); os.IsNotExist(err) {
		return nil
	}

	files, err := ioutil.ReadDir(record.String())

	NoIOPermissionsPanic(err)

	var shift *Shift

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		f, err2 := ioutil.ReadFile(path.Join(record.String(), file.Name()))

		NoIOPermissionsPanic(err2)

		workerRecord := strings.Split(string(f), ":")

		if len(workerRecord) != 4 {
			continue
		}

		employeeName := workerRecord[0]
		employeeId := CheckNumber(strconv.Atoi(workerRecord[1]))
		start := CheckNumber(strconv.Atoi(workerRecord[2]))
		end := CheckNumber(strconv.Atoi(workerRecord[3]))

		newShift := &Shift{Start: start, End: end, Employee: Employee{employeeName, employeeId}}

		if shift == nil {
			shift, _ = Add(shift, newShift)
		}

		Add(shift, newShift)
	}

	return shift
}
