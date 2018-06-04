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
	"math"
)

type Shift struct {
	Start int
	End int

	Employee

	prev *Shift
	next *Shift
}

type Shifter interface {
	Add(employee Employee)
	Overlaps(start int, end int) bool
}

type Employee struct {
	Name string
	Id int
}

func ShiftDay(date string) Shifter {
	return nil
}

func Add(first *Shift, employee *Shift) (nouse *Shift, success bool) {
	if first != nil {
		for i := first; i != nil; i = i.next {
			if success = i.Add(employee); success {
				break
			}
		}
	} else {
		success = true
	}

	return employee, success
}

// Current shift is the previous of the employee shift
func (shift *Shift) Add(employee *Shift) bool {
	if l, r := Interval(employee)(); l > r {
		return false
	}

	if shift.prev == nil && shift.Start > employee.Start && !shift.Overlaps(Interval(employee)) {
		employee.next = shift
		shift.prev = employee
		return true
	}

	if shift.next == nil && shift.End <= employee.Start && !shift.Overlaps(Interval(employee)) {
		employee.prev = shift
		shift.next = employee
		return true
	}

	if shift.Start < employee.Start && employee.Start < shift.next.Start &&
		!shift.Overlaps(Interval(employee)) {
			nextShift := shift.next

			employee.prev = shift
			shift.next = employee

			nextShift.prev = employee
			employee.next = nextShift

			return true
	}

	return false
}

func (shift *Shift) Overlaps(interval func() (int, int)) bool {
	start, end := interval()

	for i:=shift; i!=nil; i=i.next {
		if math.Min(float64(i.End), float64(end)) > math.Max(float64(i.Start), float64(start)) {
			return true
		}
	}
	return false
}

func Interval(shift *Shift) func() (start, end int) {
	return func() (start, end int) {
		return shift.Start, shift.End
	}
}

func TimeToNumeric(hour int, minute int) int {
	if hour < 0 || hour >= 24 {
		return -1
	}

	return hour * 2 - ^( ^minute & ( minute + ^0 ) ) >> 31
}