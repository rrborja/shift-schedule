# Shift Scheduler

[![Build Status](https://travis-ci.org/rrborja/shift-schedule.svg?branch=master)](https://travis-ci.org/rrborja/shift-schedule)
[![codecov](https://codecov.io/gh/rrborja/shift-schedule/branch/master/graph/badge.svg)](https://codecov.io/gh/rrborja/shift-schedule)
[![Go Report Card](https://goreportcard.com/badge/github.com/rrborja/shift-schedule)](https://goreportcard.com/report/github.com/rrborja/shift-schedule)
[![GoDoc](https://godoc.org/github.com/rrborja/shift-schedule?status.svg)](https://godoc.org/github.com/rrborja/shift-schedule) 
[![License: GPL v2+](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl.txt)

Another way to simply schedule employee's shift using Go for the backend and React for the user interface

# Usage

1. Clone this repository `git clone https://github.com/rrborja/shift-schedule`
2. Go to the `cmd` directory and run the shell script `install.sh`
3. When you run the shell script, make sure you are in the current working directory of `cmd`
4. Make sure the Node.js and Go runtime environments are installed in your system
5. Once NPM installation is complete, you may access the page `http://localhost:3000`

# API

1. To retrieve the current shift
    -   GET `http://localhost:8080/MM/dd/yyyy`
2. To add the employee to the current shift
    -   PUT `http://localhost:8080/MM/dd/yyyy`
    -   The content type for the response body is JSON
            - `name`: name of the employee
            - `id`: id of the employee
            - `start`: 0-indexed exclusive start time of the employee
            - `end`: 0-indexed inclusive end time of the employee