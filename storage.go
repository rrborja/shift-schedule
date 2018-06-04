/*
 * State Server API
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
	"os"
	"io/ioutil"
	"sort"
	"strconv"
	"fmt"
	"path"
	"strings"
)

type DayRecord struct {
	Month int
	Day int
	Year int
}

func NoIOPermissionsPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckNumber(num int, err error) int {
	if err != nil {
		panic(err)
	}

	return num
}

func (dayRecord DayRecord) String() string {
	return fmt.Sprintf("%d.%d.%d", dayRecord.Month, dayRecord.Day, dayRecord.Year)
}

func (store DayRecord) ClockIn(worker Employee, start, end int) {
	if _, err := os.Stat(store.String()); os.IsNotExist(err) {
		os.Mkdir(store.String(), 0700)
	}

	files, err := ioutil.ReadDir(store.String())

	NoIOPermissionsPanic(err)

	counter := 1

	if len(files) > 0 {
		sort.Slice(files, func(i,j int) bool{
			return files[i].Name() < files[j].Name()
		})

		filename := files[len(files)-1].Name()
		filename = filename[:len(filename)-4]

		counter = CheckNumber(strconv.Atoi(filename))

		counter++
	}

	workerFile, err3 := os.Create(path.Join(store.String(), fmt.Sprintf("%d.txt", counter)))
	defer workerFile.Close()

	NoIOPermissionsPanic(err3)

	fmt.Fprintf(workerFile, "%s:%d:%d:%d", worker.Name, worker.Id, start, end)
}

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