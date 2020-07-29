package main

import (
	"./operations"
	"fmt"
	"strings"

	//"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
)

func TestGetTodos(t *testing.T) {
	req, err := http.NewRequest("GET", "/casas/consulta", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(operations.GetTodos)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"status":200,"message":"","content":[{"_id":"5f205488296efcaa09939951","casa":"casa_1","nombre":"Miguel Aviña","debe":true,"cobros":[{"monto":800,"causa":"mantenimiento","fecha":"2020-07-27T01:45:51.441Z"},{"monto":400,"causa":"Por no recoger heces de perros","fecha":"2020-07-27T01:47:51.441Z"}]},{"_id":"5f205583296efcaa09939952","casa":"casa_2","nombre":"Luis Orozco","debe":true,"cobros":[{"monto":800,"causa":"mantenimiento","fecha":"2020-07-27T01:45:51.441Z"}]}]}`
	if normalizeString(rr.Body.String()) != normalizeString(expected) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body, expected)
	} else {
		fmt.Println("Got: ",
			rr.Body, " Expected: ", expected)
	}
}

func normalizeString (str string) string {
	var input string
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(input, "\r\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}
	return input
}