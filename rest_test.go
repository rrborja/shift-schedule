package schedule

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
