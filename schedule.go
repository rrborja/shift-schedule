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
	"math"
)

// Shift is the structure for the Shift of the employee. Start is the
// starting time of the employee inclusively while the End is the
// ending shift of the employee exclusively. This struct is the node
// of a doubly-linked list. The next node must be the next shift
// accordingly by time.
type Shift struct {
	Start int
	End   int

	Employee

	prev *Shift
	next *Shift
}

// Employee is the struct for Employee having a name and an ID
type Employee struct {
	Name string
	Id   int
}

// Add adds the new employee to the day's shift's linked list and
// automatically validates any overlapping schedules and inserts
// according to the order of time
func Add(first *Shift, employee *Shift) (*Shift, bool) {
	if first == nil {
		return employee, true
	}

	for i := first; i != nil; i = i.next {
		if i.Add(employee) {
			return employee, true
		}
	}

	return employee, false
}

// Add validates the correctness of employee's schedule to the
// current shift's list of all employees' schedule and adds to
// the shift if the new employee doesn't overlap any schedule
// of all employees
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

// Overlaps checks the interval whether it will not overlap
// to the current shifts schedule. Returns true if it overlaps.
// Returns otherwise.
func (shift *Shift) Overlaps(interval func() (int, int)) bool {
	start, end := interval()

	for i := shift; i != nil; i = i.next {
		if math.Min(float64(i.End), float64(end)) > math.Max(float64(i.Start), float64(start)) {
			return true
		}
	}
	return false
}

// Interval is the wrapper function that extract's the employee's shift
// time information into interval
func Interval(shift *Shift) func() (start, end int) {
	return func() (start, end int) {
		return shift.Start, shift.End
	}
}

// TimeToNumeric converts hour and minute into the equivalent
// integer representation that serves as the order or index
// of a 24-hour pattern
func TimeToNumeric(hour int, minute int) int {
	if hour < 0 || hour >= 24 {
		return -1
	}

	return hour*2 - ^(^minute&(minute+^0))>>31
}
