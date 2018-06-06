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
	"errors"
	"os"
	"testing"
)

func TestSavedShift(t *testing.T) {
	testRecord := DayRecord{6, 6, 2006}
	testRecord.ClockIn(Employee{"Tester", 666}, 0, 24)
	testRecord.ClockIn(Employee{"Finalizer", 777}, 24, 48)

	if _, err := os.Stat(testRecord.String()); os.IsNotExist(err) {
		t.Error("Test record was not checked in, thus not saved")
	}

	defer os.RemoveAll(testRecord.String())

	if Construct(testRecord) == nil {
		t.Error("Test record was not checked in, thus not saved")
	}
}

func TestNoIOPermissionsPanic(t *testing.T) {
	defer func() {
		recover()
	}()
	NoIOPermissionsPanic(errors.New("dummy"))
	t.Fail()
}

func TestCheckNumber(t *testing.T) {
	defer func() {
		recover()
	}()
	CheckNumber(-1, errors.New("dummy"))
	t.Fail()
}
