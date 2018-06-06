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
	"net/http"

	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"strconv"
)

// WorkerShift is the json structure for the PUT REST route used to store
// the employee's shift in the storage, assuming no schedule overlapping
// occurs.
type WorkerShift struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

// HttpRoutes is the combination of all routes to expose the scheduling
// API to the web.
// The sole route /{month}/{day}/{year} has two methods:
// 		1. GET - To retrieve the list of all employees' shift given
//				 a certain shift date
//		2. PUT - Creates a new employee shift to the list of all
// 				 employee shifts assuming it doesn't overlap with the
// 				 schedule.
func HttpRoutes() http.Handler {
	routes := mux.NewRouter()

	routes.HandleFunc("/{month}/{day}/{year}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		month := CheckNumber(strconv.Atoi(vars["month"]))
		day := CheckNumber(strconv.Atoi(vars["day"]))
		year := CheckNumber(strconv.Atoi(vars["year"]))

		fmt.Fprintf(w, print(Construct(DayRecord{month, day, year})))
	}).Methods("GET", "OPTIONS")

	routes.HandleFunc("/{month}/{day}/{year}", func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		vars := mux.Vars(r)

		month := CheckNumber(strconv.Atoi(vars["month"]))
		day := CheckNumber(strconv.Atoi(vars["day"]))
		year := CheckNumber(strconv.Atoi(vars["year"]))

		record := Construct(DayRecord{month, day, year})

		var shift *WorkerShift

		err := json.NewDecoder(r.Body).Decode(&shift)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		employee := Employee{shift.Name, shift.Id}

		if _, ok := Add(record, &Shift{Start: shift.Start, End: shift.End, Employee: employee}); ok {
			if r.Header.Get("X-Debug") != "true" {
				DayRecord{month, day, year}.ClockIn(employee, shift.Start, shift.End)
			}
		} else {
			http.Error(w, "Employee's shift was not added due to backend errors", 400)
		}
	}).Methods("PUT", "OPTIONS")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(headersOk, originsOk, methodsOk)(routes)
}

func print(shift *Shift) string {
	if shift == nil {
		return "[]"
	}

	var arr []WorkerShift

	for i := shift; i != nil; i = i.next {
		arr = append(arr, WorkerShift{i.Name, i.Id, i.Start, i.End})
	}

	res, _ := json.Marshal(arr)

	return string(res)
}
