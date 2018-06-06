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
	"testing"
)

func initialTestData() *Shift {
	jack := &Shift{Start: 0,
		End:      6,
		Employee: Employee{"Jack", 17},
	}
	jill := &Shift{
		Start:    6,
		End:      10,
		Employee: Employee{"Jill", 12},
	}
	john := &Shift{
		Start:    10,
		End:      19,
		Employee: Employee{"John", 7},
	}
	john2 := &Shift{
		Start:    19,
		End:      23,
		Employee: Employee{"John", 7},
	}
	jim := &Shift{
		Start:    26,
		End:      48,
		Employee: Employee{"Jim", 4},
	}

	jack.next = jill
	jill.next = john
	john.next = john2
	john2.next = jim

	jill.prev = jack
	john.prev = jill
	john2.prev = john
	jim.prev = john2

	return jack
}

func TestShiftGapNothingOverlapped(t *testing.T) {
	start, end := TimeToNumeric(11, 30), TimeToNumeric(13, 0)
	if initialTestData().Overlaps(Interval(&Shift{Start: start, End: end})) {
		t.Fail()
	}
}

func TestShiftGapOverlappedLeftBound(t *testing.T) {
	start, end := TimeToNumeric(11, 00), TimeToNumeric(13, 0)
	if !initialTestData().Overlaps(Interval(&Shift{Start: start, End: end})) {
		t.Fail()
	}
}

func TestShiftGapOverlappedRightBound(t *testing.T) {
	start, end := TimeToNumeric(11, 30), TimeToNumeric(13, 30)
	if !initialTestData().Overlaps(Interval(&Shift{Start: start, End: end})) {
		t.Fail()
	}
}

func TestShiftGapOverlappedBothBound(t *testing.T) {
	start, end := TimeToNumeric(11, 00), TimeToNumeric(13, 30)
	if !initialTestData().Overlaps(Interval(&Shift{Start: start, End: end})) {
		t.Fail()
	}
}

func TestAddWorkerToShiftSuccessful(t *testing.T) {
	start, end := TimeToNumeric(11, 30), TimeToNumeric(13, 0)
	ritchie := &Shift{Start: start, End: end, Employee: Employee{"Ritchie", 666}}

	shift := initialTestData()

	Add(shift, ritchie)

	var success bool

	for i := shift; i != nil; i = i.next {
		success = success || i == ritchie
	}

	// Test fails when "ritchie" is not in the shift list after adding to the shift
	if !success {
		t.Log("Employee's shift wasn't added")
		t.FailNow()
	}
}

func TestAddWorkerToFirstShift(t *testing.T) {
	shift := initialTestData()

	newFirst := shift.next

	// Delete first shift
	shift.next = nil
	newFirst.prev = nil

	newFirst, _ = Add(newFirst, shift)

	success := false

	for i := newFirst; i != nil; i = i.next {
		success = success || i == shift
	}

	if newFirst.next == nil {
		t.Log("First shift doesn't have next shifts")
		t.Fail()
	}

	if !success {
		// Test fails when "ritchie" is not in the shift list after adding to the shift
		t.Log("Employee's shift wasn't added")
		t.FailNow()
	}

}

func TestAddWorkerToLastShift(t *testing.T) {
	shift := initialTestData()

	var last *Shift

	for i := shift; i != nil; i = i.next {
		last = i
	}

	last.prev.next = nil
	last.prev = nil

	Add(shift, last)

	success := false

	for i := shift; i != nil; i = i.next {
		success = success || i == last
	}

	if !success {
		// Test fails when "ritchie" is not in the shift list after adding to the shift
		t.Log("Employee's shift wasn't added")
		t.FailNow()
	}
}

func TestTimeToNumeric(t *testing.T) {
	hour, min := 11, 30
	numeric := TimeToNumeric(hour, min)

	if numeric != 23 {
		t.Log("11:30 should be equivalent to 23")
		t.Fail()
	}
}

func TestTimeToNumericInvalidHours(t *testing.T) {
	hour, min := 24, 0
	numeric := TimeToNumeric(hour, min)

	if numeric != -1 {
		t.Log("24:00 should be an invalid time for any day")
		t.Fail()
	}
}
