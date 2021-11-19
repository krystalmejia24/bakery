package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, HealthcheckPath, nil)
	rr := httptest.NewRecorder()

	HealthcheckHandler{}.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
