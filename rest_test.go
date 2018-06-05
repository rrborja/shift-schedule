package schedule

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestWorkingGetShiftRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/06/06/2006", nil)
	if err != nil {
		t.Fatal(err)
	}

	routes := HttpRoutes()

	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestWorkingAddShiftRoute(t *testing.T) {
	obj, _ := json.Marshal(WorkerShift{"Test", 1, 0, 1})

	req, err := http.NewRequest("PUT", "/06/06/2006", bytes.NewReader(obj))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("X-Debug", "true")

	routes := HttpRoutes()

	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestPrint(t *testing.T) {
	if print(&Shift{Start: 0, End: 48, Employee: Employee{"Tester", 666}}) == "" {
		t.Fail()
	}
}

func TestWrongCorruptedRequestBody(t *testing.T) {
	req, err := http.NewRequest("PUT", "/06/06/2006", bytes.NewReader([]byte("{")))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("X-Debug", "true")

	routes := HttpRoutes()

	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Should fail the Http request")
	}
}

func TestNoRequestBody(t *testing.T) {
	req, err := http.NewRequest("PUT", "/06/06/2006", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("X-Debug", "true")

	routes := HttpRoutes()

	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Should fail the Http request")
	}
}

func TestConflictSchedulesHttpRequest(t *testing.T) {
	obj, _ := json.Marshal(WorkerShift{"Test", 1, 0, 1})

	req, err := http.NewRequest("PUT", "/06/06/2006", bytes.NewReader(obj))
	if err != nil {
		t.Fatal(err)
	}

	routes := HttpRoutes()

	recorder := httptest.NewRecorder()

	routes.ServeHTTP(recorder, req)

	defer func() {
		testRecord := DayRecord{6, 6, 2006}
		defer os.RemoveAll(testRecord.String())
	}()

	req, err = http.NewRequest("PUT", "/06/06/2006", bytes.NewReader(obj))
	if err != nil {
		t.Fatal(err)
	}

	routes.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("Should fail the Http request")
	}

	if strings.EqualFold(recorder.Body.String(), "Employee's shift was not added due to backend errors") {
		t.Error("Not the expected error message")
	}
}
